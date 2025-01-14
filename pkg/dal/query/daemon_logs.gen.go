// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/go-sigma/sigma/pkg/dal/models"
)

func newDaemonLog(db *gorm.DB, opts ...gen.DOOption) daemonLog {
	_daemonLog := daemonLog{}

	_daemonLog.daemonLogDo.UseDB(db, opts...)
	_daemonLog.daemonLogDo.UseModel(&models.DaemonLog{})

	tableName := _daemonLog.daemonLogDo.TableName()
	_daemonLog.ALL = field.NewAsterisk(tableName)
	_daemonLog.CreatedAt = field.NewTime(tableName, "created_at")
	_daemonLog.UpdatedAt = field.NewTime(tableName, "updated_at")
	_daemonLog.DeletedAt = field.NewUint(tableName, "deleted_at")
	_daemonLog.ID = field.NewInt64(tableName, "id")
	_daemonLog.NamespaceID = field.NewInt64(tableName, "namespace_id")
	_daemonLog.Type = field.NewField(tableName, "type")
	_daemonLog.Action = field.NewField(tableName, "action")
	_daemonLog.Resource = field.NewString(tableName, "resource")
	_daemonLog.Status = field.NewField(tableName, "status")
	_daemonLog.Message = field.NewBytes(tableName, "message")

	_daemonLog.fillFieldMap()

	return _daemonLog
}

type daemonLog struct {
	daemonLogDo daemonLogDo

	ALL         field.Asterisk
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Uint
	ID          field.Int64
	NamespaceID field.Int64
	Type        field.Field
	Action      field.Field
	Resource    field.String
	Status      field.Field
	Message     field.Bytes

	fieldMap map[string]field.Expr
}

func (d daemonLog) Table(newTableName string) *daemonLog {
	d.daemonLogDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d daemonLog) As(alias string) *daemonLog {
	d.daemonLogDo.DO = *(d.daemonLogDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *daemonLog) updateTableName(table string) *daemonLog {
	d.ALL = field.NewAsterisk(table)
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewUint(table, "deleted_at")
	d.ID = field.NewInt64(table, "id")
	d.NamespaceID = field.NewInt64(table, "namespace_id")
	d.Type = field.NewField(table, "type")
	d.Action = field.NewField(table, "action")
	d.Resource = field.NewString(table, "resource")
	d.Status = field.NewField(table, "status")
	d.Message = field.NewBytes(table, "message")

	d.fillFieldMap()

	return d
}

func (d *daemonLog) WithContext(ctx context.Context) *daemonLogDo {
	return d.daemonLogDo.WithContext(ctx)
}

func (d daemonLog) TableName() string { return d.daemonLogDo.TableName() }

func (d daemonLog) Alias() string { return d.daemonLogDo.Alias() }

func (d daemonLog) Columns(cols ...field.Expr) gen.Columns { return d.daemonLogDo.Columns(cols...) }

func (d *daemonLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *daemonLog) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 10)
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["id"] = d.ID
	d.fieldMap["namespace_id"] = d.NamespaceID
	d.fieldMap["type"] = d.Type
	d.fieldMap["action"] = d.Action
	d.fieldMap["resource"] = d.Resource
	d.fieldMap["status"] = d.Status
	d.fieldMap["message"] = d.Message
}

func (d daemonLog) clone(db *gorm.DB) daemonLog {
	d.daemonLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d daemonLog) replaceDB(db *gorm.DB) daemonLog {
	d.daemonLogDo.ReplaceDB(db)
	return d
}

type daemonLogDo struct{ gen.DO }

func (d daemonLogDo) Debug() *daemonLogDo {
	return d.withDO(d.DO.Debug())
}

func (d daemonLogDo) WithContext(ctx context.Context) *daemonLogDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d daemonLogDo) ReadDB() *daemonLogDo {
	return d.Clauses(dbresolver.Read)
}

func (d daemonLogDo) WriteDB() *daemonLogDo {
	return d.Clauses(dbresolver.Write)
}

func (d daemonLogDo) Session(config *gorm.Session) *daemonLogDo {
	return d.withDO(d.DO.Session(config))
}

func (d daemonLogDo) Clauses(conds ...clause.Expression) *daemonLogDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d daemonLogDo) Returning(value interface{}, columns ...string) *daemonLogDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d daemonLogDo) Not(conds ...gen.Condition) *daemonLogDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d daemonLogDo) Or(conds ...gen.Condition) *daemonLogDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d daemonLogDo) Select(conds ...field.Expr) *daemonLogDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d daemonLogDo) Where(conds ...gen.Condition) *daemonLogDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d daemonLogDo) Order(conds ...field.Expr) *daemonLogDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d daemonLogDo) Distinct(cols ...field.Expr) *daemonLogDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d daemonLogDo) Omit(cols ...field.Expr) *daemonLogDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d daemonLogDo) Join(table schema.Tabler, on ...field.Expr) *daemonLogDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d daemonLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *daemonLogDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d daemonLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *daemonLogDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d daemonLogDo) Group(cols ...field.Expr) *daemonLogDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d daemonLogDo) Having(conds ...gen.Condition) *daemonLogDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d daemonLogDo) Limit(limit int) *daemonLogDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d daemonLogDo) Offset(offset int) *daemonLogDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d daemonLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *daemonLogDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d daemonLogDo) Unscoped() *daemonLogDo {
	return d.withDO(d.DO.Unscoped())
}

func (d daemonLogDo) Create(values ...*models.DaemonLog) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d daemonLogDo) CreateInBatches(values []*models.DaemonLog, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d daemonLogDo) Save(values ...*models.DaemonLog) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d daemonLogDo) First() (*models.DaemonLog, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonLog), nil
	}
}

func (d daemonLogDo) Take() (*models.DaemonLog, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonLog), nil
	}
}

func (d daemonLogDo) Last() (*models.DaemonLog, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonLog), nil
	}
}

func (d daemonLogDo) Find() ([]*models.DaemonLog, error) {
	result, err := d.DO.Find()
	return result.([]*models.DaemonLog), err
}

func (d daemonLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DaemonLog, err error) {
	buf := make([]*models.DaemonLog, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d daemonLogDo) FindInBatches(result *[]*models.DaemonLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d daemonLogDo) Attrs(attrs ...field.AssignExpr) *daemonLogDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d daemonLogDo) Assign(attrs ...field.AssignExpr) *daemonLogDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d daemonLogDo) Joins(fields ...field.RelationField) *daemonLogDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d daemonLogDo) Preload(fields ...field.RelationField) *daemonLogDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d daemonLogDo) FirstOrInit() (*models.DaemonLog, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonLog), nil
	}
}

func (d daemonLogDo) FirstOrCreate() (*models.DaemonLog, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonLog), nil
	}
}

func (d daemonLogDo) FindByPage(offset int, limit int) (result []*models.DaemonLog, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d daemonLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d daemonLogDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d daemonLogDo) Delete(models ...*models.DaemonLog) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *daemonLogDo) withDO(do gen.Dao) *daemonLogDo {
	d.DO = *do.(*gen.DO)
	return d
}
