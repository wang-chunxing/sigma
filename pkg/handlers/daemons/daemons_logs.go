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

package daemons

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/go-sigma/sigma/pkg/consts"
	"github.com/go-sigma/sigma/pkg/types"
	"github.com/go-sigma/sigma/pkg/utils"
	"github.com/go-sigma/sigma/pkg/utils/ptr"
	"github.com/go-sigma/sigma/pkg/xerrors"
)

// Logs get the specific daemon task logs
// @Summary Get logs
// @Tags Daemon
// @Accept json
// @Produce json
// @Router /daemons/{daemon}/logs [get]
// @Param daemon path string true "Daemon name"
// @Param limit query int64 false "limit" minimum(10) maximum(100) default(10)
// @Param page query int64 false "page" minimum(1) default(1)
// @Param sort query string false "sort field"
// @Param method query string false "sort method" Enums(asc, desc)
// @Param namespace_id query string false "Namespace ID"
// @Success 200 {object} types.CommonList{items=[]types.DaemonLogItem}
// @Failure 404 {object} xerrors.ErrCode
// @Failure 500 {object} xerrors.ErrCode
func (h *handlers) Logs(c echo.Context) error {
	ctx := log.Logger.WithContext(c.Request().Context())

	var req types.GetDaemonLogsRequest
	err := utils.BindValidate(c, &req)
	if err != nil {
		log.Error().Err(err).Msg("Bind and validate request body failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeBadRequest, err.Error())
	}
	req.Pagination = utils.NormalizePagination(req.Pagination)

	daemonService := h.daemonServiceFactory.New()
	daemonLogObjs, total, err := daemonService.List(ctx, req.Pagination, req.Sortable)
	if err != nil {
		log.Error().Err(err).Msg("List daemon log failed")
		return xerrors.NewHTTPError(c, xerrors.HTTPErrCodeInternalError, err.Error())
	}
	var resp = make([]any, 0, len(daemonLogObjs))
	for _, l := range daemonLogObjs {
		resp = append(resp, types.DaemonLogItem{
			ID:        l.ID,
			Resource:  l.Resource,
			Action:    l.Action,
			Status:    l.Status,
			Message:   ptr.Of(string(l.Message)),
			CreatedAt: l.CreatedAt.Format(consts.DefaultTimePattern),
			UpdatedAt: l.UpdatedAt.Format(consts.DefaultTimePattern),
		})
	}

	return c.JSON(http.StatusOK, types.CommonList{Total: total, Items: resp})
}