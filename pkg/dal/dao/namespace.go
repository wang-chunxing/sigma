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

package dao

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/go-sigma/sigma/pkg/dal/models"
	"github.com/go-sigma/sigma/pkg/dal/query"
	"github.com/go-sigma/sigma/pkg/types"
	"github.com/go-sigma/sigma/pkg/types/enums"
	"github.com/go-sigma/sigma/pkg/utils"
	"github.com/go-sigma/sigma/pkg/utils/ptr"
)

//go:generate mockgen -destination=mocks/namespace.go -package=mocks github.com/go-sigma/sigma/pkg/dal/dao NamespaceService
//go:generate mockgen -destination=mocks/namespace_factory.go -package=mocks github.com/go-sigma/sigma/pkg/dal/dao NamespaceServiceFactory

// NamespaceService is the interface that provides the namespace service methods.
type NamespaceService interface {
	// Create creates a new namespace.
	Create(ctx context.Context, namespace *models.Namespace) error
	// FindAll ...
	FindAll(ctx context.Context) ([]*models.Namespace, error)
	// UpdateQuota updates the namespace quota.
	UpdateQuota(ctx context.Context, namespaceID, limit int64) error
	// Get gets the namespace with the specified namespace ID.
	Get(ctx context.Context, id int64) (*models.Namespace, error)
	// GetByName gets the namespace with the specified namespace name.
	GetByName(ctx context.Context, name string) (*models.Namespace, error)
	// ListNamespace lists all namespaces.
	ListNamespace(ctx context.Context, name *string, pagination types.Pagination, sort types.Sortable) ([]*models.Namespace, int64, error)
	// CountNamespace counts all namespaces.
	CountNamespace(ctx context.Context, name *string) (int64, error)
	// DeleteByID deletes the namespace with the specified namespace ID.
	DeleteByID(ctx context.Context, id int64) error
	// UpdateByID updates the namespace with the specified namespace ID.
	UpdateByID(ctx context.Context, id int64, updates map[string]interface{}) error
}

type namespaceService struct {
	tx *query.Query
}

// NamespaceServiceFactory is the interface that provides the namespace service factory methods.
type NamespaceServiceFactory interface {
	New(txs ...*query.Query) NamespaceService
}

type namespaceServiceFactory struct{}

// NewNamespaceServiceFactory creates a new namespace service factory.
func NewNamespaceServiceFactory() NamespaceServiceFactory {
	return &namespaceServiceFactory{}
}

// New creates a new namespace service.
func (f *namespaceServiceFactory) New(txs ...*query.Query) NamespaceService {
	tx := query.Q
	if len(txs) > 0 {
		tx = txs[0]
	}
	return &namespaceService{
		tx: tx,
	}
}

// Create creates a new namespace.
func (s *namespaceService) Create(ctx context.Context, namespaceObj *models.Namespace) error {
	return s.tx.Namespace.WithContext(ctx).Create(namespaceObj)
}

// FindAll ...
func (s *namespaceService) FindAll(ctx context.Context) ([]*models.Namespace, error) {
	return s.tx.Namespace.WithContext(ctx).Find()
}

// UpdateQuota updates the namespace quota.
func (s *namespaceService) UpdateQuota(ctx context.Context, namespaceID, limit int64) error {
	result, err := s.tx.Namespace.WithContext(ctx).Where(s.tx.Namespace.ID.Eq(namespaceID)).Update(s.tx.Namespace.SizeLimit, limit)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return err
}

// Get gets the namespace with the specified namespace ID.
func (s *namespaceService) Get(ctx context.Context, id int64) (*models.Namespace, error) {
	return s.tx.Namespace.WithContext(ctx).Where(s.tx.Namespace.ID.Eq(id)).First()
}

// GetByName gets the namespace with the specified namespace name.
func (s *namespaceService) GetByName(ctx context.Context, name string) (*models.Namespace, error) {
	return s.tx.Namespace.WithContext(ctx).Where(s.tx.Namespace.Name.Eq(name)).First()
}

// ListNamespace lists all namespaces.
func (s *namespaceService) ListNamespace(ctx context.Context, name *string, pagination types.Pagination, sort types.Sortable) ([]*models.Namespace, int64, error) {
	pagination = utils.NormalizePagination(pagination)
	query := s.tx.Namespace.WithContext(ctx)
	if name != nil {
		query = query.Where(s.tx.Namespace.Name.Like(fmt.Sprintf("%%%s%%", ptr.To(name))))
	}
	field, ok := s.tx.Namespace.GetFieldByName(ptr.To(sort.Sort))
	if ok {
		switch ptr.To(sort.Method) {
		case enums.SortMethodDesc:
			query = query.Order(field.Desc())
		case enums.SortMethodAsc:
			query = query.Order(field)
		default:
			query = query.Order(s.tx.Namespace.UpdatedAt.Desc())
		}
	} else {
		query = query.Order(s.tx.Namespace.UpdatedAt.Desc())
	}
	return query.FindByPage(ptr.To(pagination.Limit)*(ptr.To(pagination.Page)-1), ptr.To(pagination.Limit))
}

// CountNamespace counts all namespaces.
func (s *namespaceService) CountNamespace(ctx context.Context, name *string) (int64, error) {
	query := s.tx.Namespace.WithContext(ctx)
	if name != nil {
		query = query.Where(s.tx.Namespace.Name.Like(fmt.Sprintf("%%%s%%", ptr.To(name))))
	}
	return query.Count()
}

// DeleteByID deletes the namespace with the specified namespace ID.
func (s *namespaceService) DeleteByID(ctx context.Context, id int64) error {
	matched, err := s.tx.Namespace.WithContext(ctx).Where(s.tx.Namespace.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	if matched.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// UpdateByID updates the namespace with the specified namespace ID.
func (s *namespaceService) UpdateByID(ctx context.Context, id int64, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	matched, err := s.tx.Namespace.WithContext(ctx).Where(s.tx.Namespace.ID.Eq(id)).Updates(updates)
	if err != nil {
		return err
	}
	if matched.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
