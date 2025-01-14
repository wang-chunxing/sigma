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

package tag

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/go-sigma/sigma/pkg/consts"
	"github.com/go-sigma/sigma/pkg/types"
	"github.com/go-sigma/sigma/pkg/utils"
	"github.com/go-sigma/sigma/pkg/utils/ptr"
	"github.com/go-sigma/sigma/pkg/xerrors"
)

// ListTag handles the list tag request
//
//	@Summary	List tag
//	@Tags		Tag
//	@security	BasicAuth
//	@Accept		json
//	@Produce	json
//	@Router		/namespaces/{namespace}/tags/ [get]
//	@Param		limit		query		int64		false	"limit"	minimum(10)	maximum(100)	default(10)
//	@Param		page		query		int64		false	"page"	minimum(1)	default(1)
//	@Param		sort		query		string		false	"sort field"
//	@Param		method		query		string		false	"sort method"	Enums(asc, desc)
//	@Param		namespace	path		string		true	"namespace"
//	@Param		repository	query		string		false	"repository"
//	@Param		name		query		string		false	"search tag with name"
//	@Param		type		query		[]string	false	"search tag with type"	Enums(image, imageIndex, chart, cnab, cosign, wasm, provenance, unknown)	collectionFormat(multi)
//	@Success	200			{object}	types.CommonList{items=[]types.TagItem}
//	@Failure	404			{object}	xerrors.ErrCode
//	@Failure	500			{object}	xerrors.ErrCode
func (h *handlers) ListTag(c echo.Context) error {
	ctx := log.Logger.WithContext(c.Request().Context())

	var req types.ListTagRequest
	err := utils.BindValidate(c, &req)
	if err != nil {
		log.Error().Err(err).Msg("Bind and validate request body failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeBadRequest, err.Error())
	}

	namespaceService := h.namespaceServiceFactory.New()
	namespaceObj, err := namespaceService.GetByName(ctx, req.Namespace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Str("namespace", req.Namespace).Msg("Namespace not found")
			return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeNotFound, fmt.Sprintf("Namespace(%s) not found: %v", req.Namespace, err))
		}
		log.Error().Err(err).Str("namespace", req.Namespace).Msg("Namespace find failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeInternalError, fmt.Sprintf("Namespace(%s) find failed: %v", req.Namespace, err))
	}

	repositoryService := h.repositoryServiceFactory.New()
	repositoryObj, err := repositoryService.GetByName(ctx, req.Repository)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Str("repository", req.Repository).Msg("Repository not found")
			return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeNotFound, err.Error())
		}
		log.Error().Err(err).Msg("Repository find failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeInternalError, err.Error())
	}
	if repositoryObj.NamespaceID != namespaceObj.ID {
		log.Error().Interface("repositoryObj", repositoryObj).Interface("namespaceObj", namespaceObj).Msg("Repository's namespace ref id not equal namespace id")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeNotFound)
	}

	tagService := h.tagServiceFactory.New()
	tags, total, err := tagService.ListTag(ctx, repositoryObj.ID, req.Name, req.Type, req.Pagination, req.Sortable)
	if err != nil {
		log.Error().Err(err).Msg("List tag from db failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeInternalError, err.Error())
	}

	var resp = make([]any, 0, len(tags))
	for _, tag := range tags {
		if tag.Artifact == nil {
			log.Error().Str("image", fmt.Sprintf("%s:%s", repositoryObj.Name, tag.Name)).Msg("Some tag's artifact reference invalid")
			continue
		}
		var artifacts []types.TagItemArtifact
		for _, item := range tag.Artifact.ArtifactIndexes {
			artifacts = append(artifacts, types.TagItemArtifact{
				ID:              item.ID,
				Digest:          item.Digest,
				MediaType:       item.ContentType,
				Raw:             string(item.Raw),
				ConfigMediaType: ptr.To(item.ConfigMediaType),
				ConfigRaw:       string(item.ConfigRaw),
				Type:            string(item.Type),
				Size:            item.Size,
				BlobSize:        item.BlobsSize,
				LastPull:        item.LastPull.Time.Format(consts.DefaultTimePattern),
				PushedAt:        item.PushedAt.Format(consts.DefaultTimePattern),
				Vulnerability:   string(item.Vulnerability.Result),
				Sbom:            string(item.Sbom.Result),
				CreatedAt:       item.CreatedAt.Format(consts.DefaultTimePattern),
				UpdatedAt:       item.UpdatedAt.Format(consts.DefaultTimePattern),
			})
		}
		resp = append(resp, types.TagItem{
			ID:   tag.ID,
			Name: tag.Name,
			Artifact: types.TagItemArtifact{
				ID:              tag.Artifact.ID,
				Digest:          tag.Artifact.Digest,
				MediaType:       tag.Artifact.ContentType,
				Raw:             string(tag.Artifact.Raw),
				ConfigMediaType: ptr.To(tag.Artifact.ConfigMediaType),
				ConfigRaw:       string(tag.Artifact.ConfigRaw),
				Type:            string(tag.Artifact.Type),
				Size:            tag.Artifact.Size,
				BlobSize:        tag.Artifact.BlobsSize,
				LastPull:        tag.Artifact.LastPull.Time.Format(consts.DefaultTimePattern),
				PushedAt:        tag.Artifact.PushedAt.Format(consts.DefaultTimePattern),
				Vulnerability:   string(tag.Artifact.Vulnerability.Result),
				Sbom:            string(tag.Artifact.Sbom.Result),
				CreatedAt:       tag.Artifact.CreatedAt.Format(consts.DefaultTimePattern),
				UpdatedAt:       tag.Artifact.UpdatedAt.Format(consts.DefaultTimePattern),
			},
			Artifacts: artifacts,
			PullTimes: tag.PullTimes,
			PushedAt:  tag.PushedAt.Format(consts.DefaultTimePattern),
			CreatedAt: tag.CreatedAt.Format(consts.DefaultTimePattern),
			UpdatedAt: tag.UpdatedAt.Format(consts.DefaultTimePattern),
		})
	}

	return c.JSON(http.StatusOK, types.CommonList{Total: total, Items: resp})
}
