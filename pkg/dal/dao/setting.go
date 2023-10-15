package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/go-sigma/sigma/pkg/dal/models"
	"github.com/go-sigma/sigma/pkg/dal/query"
)

//go:generate mockgen -destination=mocks/setting.go -package=mocks github.com/go-sigma/sigma/pkg/dal/dao SettingService
//go:generate mockgen -destination=mocks/setting_factory.go -package=mocks github.com/go-sigma/sigma/pkg/dal/dao SettingServiceFactory

// SettingService is the interface that provides methods to operate on setting model
type SettingService interface {
	// Save save a new cache record in the database
	Save(ctx context.Context, key string, val []byte) error
	// Delete get a cache record
	Delete(ctx context.Context, key string) error
	// Get get a cache record
	Get(ctx context.Context, key string) (*models.Setting, error)
}

type settingService struct {
	tx *query.Query
}

// SettingServiceFactory is the interface that provides the setting service factory methods.
type SettingServiceFactory interface {
	New(txs ...*query.Query) SettingService
}

type settingServiceFactory struct{}

// NewSettingServiceFactory creates a new setting service factory.
func NewSettingServiceFactory() SettingServiceFactory {
	return &settingServiceFactory{}
}

func (f *settingServiceFactory) New(txs ...*query.Query) SettingService {
	tx := query.Q
	if len(txs) > 0 {
		tx = txs[0]
	}
	return &settingService{
		tx: tx,
	}
}

// Create creates a new setting record in the database
func (s settingService) Save(ctx context.Context, key string, val []byte) error {
	var setting = models.Setting{Key: key, Val: val}
	return s.tx.Setting.WithContext(ctx).Clauses(clause.OnConflict{UpdateAll: true}).Create(&setting)
}

// Delete get a cache record
func (s settingService) Delete(ctx context.Context, key string) error {
	matched, err := s.tx.Setting.WithContext(ctx).Unscoped().Where(s.tx.Setting.Key.Eq(key)).Delete()
	if err != nil {
		return err
	}
	if matched.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Get get a cache record
func (s settingService) Get(ctx context.Context, key string) (*models.Setting, error) {
	return s.tx.Setting.WithContext(ctx).Where(s.tx.Setting.Key.Eq(key)).First()
}
