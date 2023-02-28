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

package distribution

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	dtspecv1 "github.com/opencontainers/distribution-spec/specs-go/v1"
	"gorm.io/gorm"

	services "github.com/ximager/ximager/pkg/dal/dao"
	"github.com/ximager/ximager/pkg/dal/models"
	"github.com/ximager/ximager/pkg/xerrors"
)

// ListRepositories handles the list repositories request
func (h *handlers) ListRepositories(c echo.Context) error {
	var nStr = c.QueryParam("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return xerrors.GenDsResponseError(c, xerrors.ErrorCodePaginationNumberInvalid)
	}

	ctx := c.Request().Context()

	lastFound := false
	var lastID uint64 = 0

	repositoryService := services.NewRepositoryService()
	var last = c.QueryParam("last")
	if last != "" {
		tagObj, err := repositoryService.GetByName(ctx, last)
		if err != nil && err != gorm.ErrRecordNotFound {
			return xerrors.GenDsResponseError(c, xerrors.ErrorCodeUnknown)
		}
		lastFound = true
		lastID = tagObj.ID
	}

	var repositories []*models.Repository
	if !lastFound {
		repositories, err = repositoryService.ListByDtPagination(ctx, n)
	} else {
		repositories, err = repositoryService.ListByDtPagination(ctx, n, lastID)
	}
	if err != nil {
		return xerrors.GenDsResponseError(c, xerrors.ErrorCodeUnknown)
	}
	var names []string
	for _, repository := range repositories {
		names = append(names, repository.Name)
	}

	var repositoryList = dtspecv1.RepositoryList{
		Repositories: names,
	}

	location := fmt.Sprintf("%s://%s%s", c.Scheme(), c.Request().Host, c.Request().URL.Path)
	values := url.Values{}
	values.Set("n", nStr)
	if len(repositories) > 0 {
		values.Set("last", repositories[len(repositories)-1].Name)
		c.Response().Header().Set("Link", fmt.Sprintf("<%s?%s>; rel=\"next\"", location, values.Encode()))
	}

	return c.JSON(http.StatusOK, repositoryList)
}
