// Copyright 2023 sigma
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package upload

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/distribution/reference"
	"github.com/labstack/echo/v4"
	"github.com/opencontainers/go-digest"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/go-sigma/sigma/pkg/consts"
	"github.com/go-sigma/sigma/pkg/dal/models"
	"github.com/go-sigma/sigma/pkg/storage"
	"github.com/go-sigma/sigma/pkg/utils"
	"github.com/go-sigma/sigma/pkg/utils/counter"
	"github.com/go-sigma/sigma/pkg/utils/hash"
	"github.com/go-sigma/sigma/pkg/xerrors"
)

// PutUpload handles the put upload request
func (h *handler) PutUpload(c echo.Context) error {
	dgest, err := digest.Parse(c.QueryParam("digest"))
	if err != nil {
		log.Error().Err(err).Str("digest", c.QueryParam("digest")).Msg("Parse digest failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeDigestInvalid)
	}
	c.Response().Header().Set(consts.ContentDigest, dgest.String())

	uri := c.Request().URL.Path
	uploadID := strings.TrimPrefix(uri[strings.LastIndex(uri, "/"):], "/")
	c.Response().Header().Set("Location", fmt.Sprintf("%s://%s%s", c.Scheme(), c.Request().Host, uri))

	ctx := log.Logger.WithContext(c.Request().Context())

	blobUploadService := h.blobUploadServiceFactory.New()
	uploadObj, err := blobUploadService.GetLastPart(ctx, uploadID)
	if err != nil {
		log.Error().Err(err).Msg("Get blob upload record failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}
	srcPath := fmt.Sprintf("%s/%s", consts.BlobUploads, uploadObj.FileID)

	blobService := h.blobServiceFactory.New()
	exist, err := blobService.Exists(ctx, dgest.String())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Str("digest", dgest.String()).Msg("Check blob exist failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}
	if exist {
		err = storage.Driver.AbortUpload(ctx, srcPath, uploadObj.UploadID)
		if err != nil {
			log.Error().Err(err).Msg("Abort upload failed")
			return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
		}
		return c.NoContent(http.StatusCreated)
	}

	repository := h.getRepository(c)
	if !reference.NameRegexp.MatchString(repository) {
		log.Error().Str("repository", repository).Msg("Invalid repository name")
		return xerrors.NewDSError(c, xerrors.DSErrCodeNameInvalid)
	}

	etags, err := blobUploadService.TotalEtagsByUploadID(ctx, uploadID)
	if err != nil {
		log.Error().Err(err).Msg("Get blob upload etags failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}

	sizeBefore, err := blobUploadService.TotalSizeByUploadID(ctx, uploadID)
	if err != nil {
		log.Error().Err(err).Msg("Get blob upload size failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}

	length, err := utils.GetContentLength(c.Request())
	if err != nil {
		log.Error().Err(err).Msg("Get content length failed")
		return err
	}
	if length != 0 {
		counterReader := counter.NewCounter(c.Request().Body)
		etag, err := storage.Driver.UploadPart(ctx, srcPath, uploadObj.UploadID, int64(uploadObj.PartNumber+1), counterReader)
		if err != nil {
			log.Error().Err(err).Msg("Upload part failed")
			return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
		}
		size := counterReader.Count()
		etags = append(etags, etag)
		err = blobUploadService.Create(ctx, &models.BlobUpload{
			PartNumber: uploadObj.PartNumber + 1,
			UploadID:   uploadID,
			Etag:       strings.Trim(etag, "\""),
			Repository: repository,
			FileID:     uploadObj.FileID,
			Size:       size,
		})
		if err != nil {
			log.Error().Err(err).Msg("Create blob upload record failed")
			return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
		}
		c.Response().Header().Set("Content-Range", fmt.Sprintf("%d-%d", sizeBefore, sizeBefore+size))
	}

	err = storage.Driver.CommitUpload(ctx, srcPath, uploadID, etags)
	if err != nil {
		log.Error().Err(err).Str("id", uploadID).Strs("etags", etags).Msg("Commit upload failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}

	srcPathReader, err := storage.Driver.Reader(ctx, srcPath, 0)
	if err != nil {
		log.Error().Err(err).Str("srcPath", srcPath).Msg("Get blob upload failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}
	srcPathHash, err := hash.Reader(srcPathReader, dgest.Algorithm().String())
	if err != nil {
		log.Error().Err(err).Str("srcPath", srcPath).Msg("Hash blob upload failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}
	if fmt.Sprintf("%s:%s", dgest.Algorithm().String(), srcPathHash) != dgest.String() {
		log.Error().Str("srcPath", srcPath).Str("srcPathHash", fmt.Sprintf("%s:%s", dgest.Algorithm().String(), srcPathHash)).Str("targetHash", dgest.String()).Msg("Hash blob upload mismatch")
		return xerrors.NewDSError(c, xerrors.DSErrCodeBlobUploadDigestMismatch)
	}

	destPath := path.Join(consts.Blobs, utils.GenPathByDigest(dgest))
	err = storage.Driver.Move(ctx, srcPath, destPath)
	if err != nil {
		log.Error().Err(err).Str("path", srcPath).Str("digest", dgest.String()).Str("dest", destPath).Msg("Move blob failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}

	err = storage.Driver.Delete(ctx, srcPath)
	if err != nil {
		log.Error().Err(err).Msg("Delete blob upload failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}

	err = blobUploadService.DeleteByUploadID(ctx, uploadID)
	if err != nil {
		log.Error().Err(err).Msg("Delete blob upload record failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}

	contentType := c.Request().Header.Get("Content-Type")
	err = blobService.Create(ctx, &models.Blob{
		Digest:      dgest.String(),
		Size:        sizeBefore + length,
		ContentType: contentType,
		PushedAt:    time.Now(),
	})
	if err != nil {
		log.Error().Err(err).Msg("Create blob record failed")
		return xerrors.NewDSError(c, xerrors.DSErrCodeUnknown)
	}

	return c.NoContent(http.StatusCreated)
}
