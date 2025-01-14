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
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/go-sigma/sigma/pkg/types"
	"github.com/go-sigma/sigma/pkg/utils"
	"github.com/go-sigma/sigma/pkg/xerrors"
)

// DeleteTag handles the delete tag request
//
//	@Summary	Delete tag
//	@Tags		Tag
//	@security	BasicAuth
//	@Accept		json
//	@Produce	json
//	@Router		/namespaces/{namespace}/tags/{id} [delete]
//	@Param		namespace	path	string	true	"Namespace"
//	@Param		id			path	string	true	"Tag ID"
//	@Param		repository	query	string	false	"repository"
//	@Success	204
//	@Failure	404	{object}	xerrors.ErrCode
//	@Failure	500	{object}	xerrors.ErrCode
func (h *handlers) DeleteTag(c echo.Context) error {
	ctx := log.Logger.WithContext(c.Request().Context())

	var req types.DeleteTagRequest
	err := utils.BindValidate(c, &req)
	if err != nil {
		log.Error().Err(err).Msg("Bind and validate request body failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeBadRequest, err.Error())
	}

	tagService := h.tagServiceFactory.New()
	err = tagService.DeleteByID(ctx, req.ID)
	if err != nil {
		log.Error().Err(err).Msg("Delete tag failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeInternalError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
