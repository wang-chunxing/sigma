// Copyright 2023 XImager
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

package namespaces

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/ximager/ximager/pkg/types"
	"github.com/ximager/ximager/pkg/utils"
	"github.com/ximager/ximager/pkg/xerrors"
)

// DeleteNamespace handles the delete namespace request
// @Summary Delete namespace
// @security BasicAuth
// @Tags Namespace
// @Accept json
// @Produce json
// @Router /namespaces/{id} [delete]
// @Success 204
func (h *handlers) DeleteNamespace(c echo.Context) error {
	ctx := log.Logger.WithContext(c.Request().Context())

	var req types.DeleteNamespaceRequest
	err := utils.BindValidate(c, &req)
	if err != nil {
		log.Error().Err(err).Msg("Bind and validate request body failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeBadRequest, err.Error())
	}

	namespaceService := h.namespaceServiceFactory.New()
	err = namespaceService.DeleteByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error().Err(err).Msg("Delete namespace from db failed")
			return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeNotFound, err.Error())
		}
		log.Error().Err(err).Msg("Delete namespace from db failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeInternalError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}