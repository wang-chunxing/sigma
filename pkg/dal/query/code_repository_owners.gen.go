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

func newCodeRepositoryOwner(db *gorm.DB, opts ...gen.DOOption) codeRepositoryOwner {
	_codeRepositoryOwner := codeRepositoryOwner{}

	_codeRepositoryOwner.codeRepositoryOwnerDo.UseDB(db, opts...)
	_codeRepositoryOwner.codeRepositoryOwnerDo.UseModel(&models.CodeRepositoryOwner{})

	tableName := _codeRepositoryOwner.codeRepositoryOwnerDo.TableName()
	_codeRepositoryOwner.ALL = field.NewAsterisk(tableName)
	_codeRepositoryOwner.CreatedAt = field.NewTime(tableName, "created_at")
	_codeRepositoryOwner.UpdatedAt = field.NewTime(tableName, "updated_at")
	_codeRepositoryOwner.DeletedAt = field.NewUint(tableName, "deleted_at")
	_codeRepositoryOwner.ID = field.NewInt64(tableName, "id")
	_codeRepositoryOwner.User3rdPartyID = field.NewInt64(tableName, "user_3rdparty_id")
	_codeRepositoryOwner.OwnerID = field.NewString(tableName, "owner_id")
	_codeRepositoryOwner.Owner = field.NewString(tableName, "owner")
	_codeRepositoryOwner.IsOrg = field.NewBool(tableName, "is_org")
	_codeRepositoryOwner.User3rdParty = codeRepositoryOwnerBelongsToUser3rdParty{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User3rdParty", "models.User3rdParty"),
		User: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User3rdParty.User", "models.User"),
		},
	}

	_codeRepositoryOwner.fillFieldMap()

	return _codeRepositoryOwner
}

type codeRepositoryOwner struct {
	codeRepositoryOwnerDo codeRepositoryOwnerDo

	ALL            field.Asterisk
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Uint
	ID             field.Int64
	User3rdPartyID field.Int64
	OwnerID        field.String
	Owner          field.String
	IsOrg          field.Bool
	User3rdParty   codeRepositoryOwnerBelongsToUser3rdParty

	fieldMap map[string]field.Expr
}

func (c codeRepositoryOwner) Table(newTableName string) *codeRepositoryOwner {
	c.codeRepositoryOwnerDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c codeRepositoryOwner) As(alias string) *codeRepositoryOwner {
	c.codeRepositoryOwnerDo.DO = *(c.codeRepositoryOwnerDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *codeRepositoryOwner) updateTableName(table string) *codeRepositoryOwner {
	c.ALL = field.NewAsterisk(table)
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewUint(table, "deleted_at")
	c.ID = field.NewInt64(table, "id")
	c.User3rdPartyID = field.NewInt64(table, "user_3rdparty_id")
	c.OwnerID = field.NewString(table, "owner_id")
	c.Owner = field.NewString(table, "owner")
	c.IsOrg = field.NewBool(table, "is_org")

	c.fillFieldMap()

	return c
}

func (c *codeRepositoryOwner) WithContext(ctx context.Context) *codeRepositoryOwnerDo {
	return c.codeRepositoryOwnerDo.WithContext(ctx)
}

func (c codeRepositoryOwner) TableName() string { return c.codeRepositoryOwnerDo.TableName() }

func (c codeRepositoryOwner) Alias() string { return c.codeRepositoryOwnerDo.Alias() }

func (c codeRepositoryOwner) Columns(cols ...field.Expr) gen.Columns {
	return c.codeRepositoryOwnerDo.Columns(cols...)
}

func (c *codeRepositoryOwner) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *codeRepositoryOwner) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["id"] = c.ID
	c.fieldMap["user_3rdparty_id"] = c.User3rdPartyID
	c.fieldMap["owner_id"] = c.OwnerID
	c.fieldMap["owner"] = c.Owner
	c.fieldMap["is_org"] = c.IsOrg

}

func (c codeRepositoryOwner) clone(db *gorm.DB) codeRepositoryOwner {
	c.codeRepositoryOwnerDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c codeRepositoryOwner) replaceDB(db *gorm.DB) codeRepositoryOwner {
	c.codeRepositoryOwnerDo.ReplaceDB(db)
	return c
}

type codeRepositoryOwnerBelongsToUser3rdParty struct {
	db *gorm.DB

	field.RelationField

	User struct {
		field.RelationField
	}
}

func (a codeRepositoryOwnerBelongsToUser3rdParty) Where(conds ...field.Expr) *codeRepositoryOwnerBelongsToUser3rdParty {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a codeRepositoryOwnerBelongsToUser3rdParty) WithContext(ctx context.Context) *codeRepositoryOwnerBelongsToUser3rdParty {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a codeRepositoryOwnerBelongsToUser3rdParty) Session(session *gorm.Session) *codeRepositoryOwnerBelongsToUser3rdParty {
	a.db = a.db.Session(session)
	return &a
}

func (a codeRepositoryOwnerBelongsToUser3rdParty) Model(m *models.CodeRepositoryOwner) *codeRepositoryOwnerBelongsToUser3rdPartyTx {
	return &codeRepositoryOwnerBelongsToUser3rdPartyTx{a.db.Model(m).Association(a.Name())}
}

type codeRepositoryOwnerBelongsToUser3rdPartyTx struct{ tx *gorm.Association }

func (a codeRepositoryOwnerBelongsToUser3rdPartyTx) Find() (result *models.User3rdParty, err error) {
	return result, a.tx.Find(&result)
}

func (a codeRepositoryOwnerBelongsToUser3rdPartyTx) Append(values ...*models.User3rdParty) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a codeRepositoryOwnerBelongsToUser3rdPartyTx) Replace(values ...*models.User3rdParty) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a codeRepositoryOwnerBelongsToUser3rdPartyTx) Delete(values ...*models.User3rdParty) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a codeRepositoryOwnerBelongsToUser3rdPartyTx) Clear() error {
	return a.tx.Clear()
}

func (a codeRepositoryOwnerBelongsToUser3rdPartyTx) Count() int64 {
	return a.tx.Count()
}

type codeRepositoryOwnerDo struct{ gen.DO }

func (c codeRepositoryOwnerDo) Debug() *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Debug())
}

func (c codeRepositoryOwnerDo) WithContext(ctx context.Context) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c codeRepositoryOwnerDo) ReadDB() *codeRepositoryOwnerDo {
	return c.Clauses(dbresolver.Read)
}

func (c codeRepositoryOwnerDo) WriteDB() *codeRepositoryOwnerDo {
	return c.Clauses(dbresolver.Write)
}

func (c codeRepositoryOwnerDo) Session(config *gorm.Session) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Session(config))
}

func (c codeRepositoryOwnerDo) Clauses(conds ...clause.Expression) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c codeRepositoryOwnerDo) Returning(value interface{}, columns ...string) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c codeRepositoryOwnerDo) Not(conds ...gen.Condition) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c codeRepositoryOwnerDo) Or(conds ...gen.Condition) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c codeRepositoryOwnerDo) Select(conds ...field.Expr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c codeRepositoryOwnerDo) Where(conds ...gen.Condition) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c codeRepositoryOwnerDo) Order(conds ...field.Expr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c codeRepositoryOwnerDo) Distinct(cols ...field.Expr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c codeRepositoryOwnerDo) Omit(cols ...field.Expr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c codeRepositoryOwnerDo) Join(table schema.Tabler, on ...field.Expr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c codeRepositoryOwnerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c codeRepositoryOwnerDo) RightJoin(table schema.Tabler, on ...field.Expr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c codeRepositoryOwnerDo) Group(cols ...field.Expr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c codeRepositoryOwnerDo) Having(conds ...gen.Condition) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c codeRepositoryOwnerDo) Limit(limit int) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c codeRepositoryOwnerDo) Offset(offset int) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c codeRepositoryOwnerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c codeRepositoryOwnerDo) Unscoped() *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Unscoped())
}

func (c codeRepositoryOwnerDo) Create(values ...*models.CodeRepositoryOwner) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c codeRepositoryOwnerDo) CreateInBatches(values []*models.CodeRepositoryOwner, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c codeRepositoryOwnerDo) Save(values ...*models.CodeRepositoryOwner) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c codeRepositoryOwnerDo) First() (*models.CodeRepositoryOwner, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryOwner), nil
	}
}

func (c codeRepositoryOwnerDo) Take() (*models.CodeRepositoryOwner, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryOwner), nil
	}
}

func (c codeRepositoryOwnerDo) Last() (*models.CodeRepositoryOwner, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryOwner), nil
	}
}

func (c codeRepositoryOwnerDo) Find() ([]*models.CodeRepositoryOwner, error) {
	result, err := c.DO.Find()
	return result.([]*models.CodeRepositoryOwner), err
}

func (c codeRepositoryOwnerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.CodeRepositoryOwner, err error) {
	buf := make([]*models.CodeRepositoryOwner, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c codeRepositoryOwnerDo) FindInBatches(result *[]*models.CodeRepositoryOwner, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c codeRepositoryOwnerDo) Attrs(attrs ...field.AssignExpr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c codeRepositoryOwnerDo) Assign(attrs ...field.AssignExpr) *codeRepositoryOwnerDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c codeRepositoryOwnerDo) Joins(fields ...field.RelationField) *codeRepositoryOwnerDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c codeRepositoryOwnerDo) Preload(fields ...field.RelationField) *codeRepositoryOwnerDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c codeRepositoryOwnerDo) FirstOrInit() (*models.CodeRepositoryOwner, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryOwner), nil
	}
}

func (c codeRepositoryOwnerDo) FirstOrCreate() (*models.CodeRepositoryOwner, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryOwner), nil
	}
}

func (c codeRepositoryOwnerDo) FindByPage(offset int, limit int) (result []*models.CodeRepositoryOwner, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c codeRepositoryOwnerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c codeRepositoryOwnerDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c codeRepositoryOwnerDo) Delete(models ...*models.CodeRepositoryOwner) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *codeRepositoryOwnerDo) withDO(do gen.Dao) *codeRepositoryOwnerDo {
	c.DO = *do.(*gen.DO)
	return c
}
