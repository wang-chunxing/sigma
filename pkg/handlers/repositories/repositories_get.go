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

package repositories

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/go-sigma/sigma/pkg/consts"
	"github.com/go-sigma/sigma/pkg/types"
	"github.com/go-sigma/sigma/pkg/types/enums"
	"github.com/go-sigma/sigma/pkg/utils"
	"github.com/go-sigma/sigma/pkg/utils/ptr"
	"github.com/go-sigma/sigma/pkg/xerrors"
)

// GetRepository handles the get repository request
//
//	@Summary	Get repository
//	@Tags		Repository
//	@security	BasicAuth
//	@Accept		json
//	@Produce	json
//	@Router		/namespaces/{namespace}/repositories/{id} [get]
//	@Param		namespace	path		string	true	"Namespace"
//	@Param		id			path		string	true	"Repository ID"
//	@Success	200			{object}	types.RepositoryItem
//	@Failure	404			{object}	xerrors.ErrCode
//	@Failure	500			{object}	xerrors.ErrCode
func (h *handlers) GetRepository(c echo.Context) error {
	ctx := log.Logger.WithContext(c.Request().Context())

	var req types.GetRepositoryRequest
	err := utils.BindValidate(c, &req)
	if err != nil {
		log.Error().Err(err).Msg("Bind and validate request body failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeBadRequest, err.Error())
	}

	repositoryService := h.repositoryServiceFactory.New()
	repositoryObj, err := repositoryService.Get(ctx, req.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Error().Err(err).Msg("Get repository from db failed")
			return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeNotFound, err.Error())
		}
		log.Error().Err(err).Msg("Get repository from db failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeInternalError, err.Error())
	}

	var builderItemObj *types.BuilderItem
	if repositoryObj.Builder != nil {
		platforms := []enums.OciPlatform{}
		for _, p := range strings.Split(repositoryObj.Builder.BuildkitPlatforms, ",") {
			platforms = append(platforms, enums.OciPlatform(p))
		}

		builderItemObj = &types.BuilderItem{
			ID:           repositoryObj.Builder.ID,
			RepositoryID: repositoryObj.Builder.RepositoryID,

			Source: repositoryObj.Builder.Source,

			CodeRepositoryID: repositoryObj.Builder.CodeRepositoryID,

			Dockerfile: ptr.Of(string(repositoryObj.Builder.Dockerfile)),

			ScmRepository:     repositoryObj.Builder.ScmRepository,
			ScmCredentialType: repositoryObj.Builder.ScmCredentialType,
			ScmSshKey:         repositoryObj.Builder.ScmSshKey,
			ScmToken:          repositoryObj.Builder.ScmToken,
			ScmUsername:       repositoryObj.Builder.ScmUsername,
			ScmPassword:       repositoryObj.Builder.ScmPassword,

			ScmBranch: repositoryObj.Builder.ScmBranch,

			ScmDepth:     repositoryObj.Builder.ScmDepth,
			ScmSubmodule: repositoryObj.Builder.ScmSubmodule,

			CronRule:        repositoryObj.Builder.CronRule,
			CronBranch:      repositoryObj.Builder.CronBranch,
			CronTagTemplate: repositoryObj.Builder.CronTagTemplate,

			WebhookBranchName:        repositoryObj.Builder.WebhookBranchName,
			WebhookBranchTagTemplate: repositoryObj.Builder.WebhookBranchTagTemplate,
			WebhookTagTagTemplate:    repositoryObj.Builder.WebhookTagTagTemplate,

			BuildkitInsecureRegistries: strings.Split(repositoryObj.Builder.BuildkitInsecureRegistries, ","),
			BuildkitContext:            repositoryObj.Builder.BuildkitContext,
			BuildkitDockerfile:         repositoryObj.Builder.BuildkitDockerfile,
			BuildkitPlatforms:          platforms,
			BuildkitBuildArgs:          repositoryObj.Builder.BuildkitBuildArgs,
		}
	}

	return c.JSON(http.StatusOK, types.RepositoryItem{
		ID:          repositoryObj.ID,
		NamespaceID: repositoryObj.NamespaceID,
		Name:        repositoryObj.Name,
		Description: repositoryObj.Description,
		Overview:    ptr.Of(string(repositoryObj.Overview)),
		Visibility:  repositoryObj.Visibility,
		SizeLimit:   ptr.Of(repositoryObj.SizeLimit),
		Size:        ptr.Of(repositoryObj.Size),
		TagCount:    repositoryObj.TagCount,
		Builder:     builderItemObj,
		CreatedAt:   repositoryObj.CreatedAt.Format(consts.DefaultTimePattern),
		UpdatedAt:   repositoryObj.UpdatedAt.Format(consts.DefaultTimePattern),
	})
}
