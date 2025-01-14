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

func newCodeRepositoryCloneCredential(db *gorm.DB, opts ...gen.DOOption) codeRepositoryCloneCredential {
	_codeRepositoryCloneCredential := codeRepositoryCloneCredential{}

	_codeRepositoryCloneCredential.codeRepositoryCloneCredentialDo.UseDB(db, opts...)
	_codeRepositoryCloneCredential.codeRepositoryCloneCredentialDo.UseModel(&models.CodeRepositoryCloneCredential{})

	tableName := _codeRepositoryCloneCredential.codeRepositoryCloneCredentialDo.TableName()
	_codeRepositoryCloneCredential.ALL = field.NewAsterisk(tableName)
	_codeRepositoryCloneCredential.CreatedAt = field.NewTime(tableName, "created_at")
	_codeRepositoryCloneCredential.UpdatedAt = field.NewTime(tableName, "updated_at")
	_codeRepositoryCloneCredential.DeletedAt = field.NewUint(tableName, "deleted_at")
	_codeRepositoryCloneCredential.ID = field.NewInt64(tableName, "id")
	_codeRepositoryCloneCredential.User3rdPartyID = field.NewInt64(tableName, "user_3rdparty_id")
	_codeRepositoryCloneCredential.Type = field.NewField(tableName, "type")
	_codeRepositoryCloneCredential.SshKey = field.NewString(tableName, "ssh_key")
	_codeRepositoryCloneCredential.Username = field.NewString(tableName, "username")
	_codeRepositoryCloneCredential.Password = field.NewString(tableName, "password")
	_codeRepositoryCloneCredential.Token = field.NewString(tableName, "token")
	_codeRepositoryCloneCredential.User3rdParty = codeRepositoryCloneCredentialBelongsToUser3rdParty{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User3rdParty", "models.User3rdParty"),
		User: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User3rdParty.User", "models.User"),
		},
	}

	_codeRepositoryCloneCredential.fillFieldMap()

	return _codeRepositoryCloneCredential
}

type codeRepositoryCloneCredential struct {
	codeRepositoryCloneCredentialDo codeRepositoryCloneCredentialDo

	ALL            field.Asterisk
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Uint
	ID             field.Int64
	User3rdPartyID field.Int64
	Type           field.Field
	SshKey         field.String
	Username       field.String
	Password       field.String
	Token          field.String
	User3rdParty   codeRepositoryCloneCredentialBelongsToUser3rdParty

	fieldMap map[string]field.Expr
}

func (c codeRepositoryCloneCredential) Table(newTableName string) *codeRepositoryCloneCredential {
	c.codeRepositoryCloneCredentialDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c codeRepositoryCloneCredential) As(alias string) *codeRepositoryCloneCredential {
	c.codeRepositoryCloneCredentialDo.DO = *(c.codeRepositoryCloneCredentialDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *codeRepositoryCloneCredential) updateTableName(table string) *codeRepositoryCloneCredential {
	c.ALL = field.NewAsterisk(table)
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewUint(table, "deleted_at")
	c.ID = field.NewInt64(table, "id")
	c.User3rdPartyID = field.NewInt64(table, "user_3rdparty_id")
	c.Type = field.NewField(table, "type")
	c.SshKey = field.NewString(table, "ssh_key")
	c.Username = field.NewString(table, "username")
	c.Password = field.NewString(table, "password")
	c.Token = field.NewString(table, "token")

	c.fillFieldMap()

	return c
}

func (c *codeRepositoryCloneCredential) WithContext(ctx context.Context) *codeRepositoryCloneCredentialDo {
	return c.codeRepositoryCloneCredentialDo.WithContext(ctx)
}

func (c codeRepositoryCloneCredential) TableName() string {
	return c.codeRepositoryCloneCredentialDo.TableName()
}

func (c codeRepositoryCloneCredential) Alias() string {
	return c.codeRepositoryCloneCredentialDo.Alias()
}

func (c codeRepositoryCloneCredential) Columns(cols ...field.Expr) gen.Columns {
	return c.codeRepositoryCloneCredentialDo.Columns(cols...)
}

func (c *codeRepositoryCloneCredential) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *codeRepositoryCloneCredential) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 11)
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["id"] = c.ID
	c.fieldMap["user_3rdparty_id"] = c.User3rdPartyID
	c.fieldMap["type"] = c.Type
	c.fieldMap["ssh_key"] = c.SshKey
	c.fieldMap["username"] = c.Username
	c.fieldMap["password"] = c.Password
	c.fieldMap["token"] = c.Token

}

func (c codeRepositoryCloneCredential) clone(db *gorm.DB) codeRepositoryCloneCredential {
	c.codeRepositoryCloneCredentialDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c codeRepositoryCloneCredential) replaceDB(db *gorm.DB) codeRepositoryCloneCredential {
	c.codeRepositoryCloneCredentialDo.ReplaceDB(db)
	return c
}

type codeRepositoryCloneCredentialBelongsToUser3rdParty struct {
	db *gorm.DB

	field.RelationField

	User struct {
		field.RelationField
	}
}

func (a codeRepositoryCloneCredentialBelongsToUser3rdParty) Where(conds ...field.Expr) *codeRepositoryCloneCredentialBelongsToUser3rdParty {
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

func (a codeRepositoryCloneCredentialBelongsToUser3rdParty) WithContext(ctx context.Context) *codeRepositoryCloneCredentialBelongsToUser3rdParty {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a codeRepositoryCloneCredentialBelongsToUser3rdParty) Session(session *gorm.Session) *codeRepositoryCloneCredentialBelongsToUser3rdParty {
	a.db = a.db.Session(session)
	return &a
}

func (a codeRepositoryCloneCredentialBelongsToUser3rdParty) Model(m *models.CodeRepositoryCloneCredential) *codeRepositoryCloneCredentialBelongsToUser3rdPartyTx {
	return &codeRepositoryCloneCredentialBelongsToUser3rdPartyTx{a.db.Model(m).Association(a.Name())}
}

type codeRepositoryCloneCredentialBelongsToUser3rdPartyTx struct{ tx *gorm.Association }

func (a codeRepositoryCloneCredentialBelongsToUser3rdPartyTx) Find() (result *models.User3rdParty, err error) {
	return result, a.tx.Find(&result)
}

func (a codeRepositoryCloneCredentialBelongsToUser3rdPartyTx) Append(values ...*models.User3rdParty) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a codeRepositoryCloneCredentialBelongsToUser3rdPartyTx) Replace(values ...*models.User3rdParty) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a codeRepositoryCloneCredentialBelongsToUser3rdPartyTx) Delete(values ...*models.User3rdParty) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a codeRepositoryCloneCredentialBelongsToUser3rdPartyTx) Clear() error {
	return a.tx.Clear()
}

func (a codeRepositoryCloneCredentialBelongsToUser3rdPartyTx) Count() int64 {
	return a.tx.Count()
}

type codeRepositoryCloneCredentialDo struct{ gen.DO }

func (c codeRepositoryCloneCredentialDo) Debug() *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Debug())
}

func (c codeRepositoryCloneCredentialDo) WithContext(ctx context.Context) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c codeRepositoryCloneCredentialDo) ReadDB() *codeRepositoryCloneCredentialDo {
	return c.Clauses(dbresolver.Read)
}

func (c codeRepositoryCloneCredentialDo) WriteDB() *codeRepositoryCloneCredentialDo {
	return c.Clauses(dbresolver.Write)
}

func (c codeRepositoryCloneCredentialDo) Session(config *gorm.Session) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Session(config))
}

func (c codeRepositoryCloneCredentialDo) Clauses(conds ...clause.Expression) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c codeRepositoryCloneCredentialDo) Returning(value interface{}, columns ...string) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c codeRepositoryCloneCredentialDo) Not(conds ...gen.Condition) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c codeRepositoryCloneCredentialDo) Or(conds ...gen.Condition) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c codeRepositoryCloneCredentialDo) Select(conds ...field.Expr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c codeRepositoryCloneCredentialDo) Where(conds ...gen.Condition) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c codeRepositoryCloneCredentialDo) Order(conds ...field.Expr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c codeRepositoryCloneCredentialDo) Distinct(cols ...field.Expr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c codeRepositoryCloneCredentialDo) Omit(cols ...field.Expr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c codeRepositoryCloneCredentialDo) Join(table schema.Tabler, on ...field.Expr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c codeRepositoryCloneCredentialDo) LeftJoin(table schema.Tabler, on ...field.Expr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c codeRepositoryCloneCredentialDo) RightJoin(table schema.Tabler, on ...field.Expr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c codeRepositoryCloneCredentialDo) Group(cols ...field.Expr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c codeRepositoryCloneCredentialDo) Having(conds ...gen.Condition) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c codeRepositoryCloneCredentialDo) Limit(limit int) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c codeRepositoryCloneCredentialDo) Offset(offset int) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c codeRepositoryCloneCredentialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c codeRepositoryCloneCredentialDo) Unscoped() *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Unscoped())
}

func (c codeRepositoryCloneCredentialDo) Create(values ...*models.CodeRepositoryCloneCredential) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c codeRepositoryCloneCredentialDo) CreateInBatches(values []*models.CodeRepositoryCloneCredential, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c codeRepositoryCloneCredentialDo) Save(values ...*models.CodeRepositoryCloneCredential) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c codeRepositoryCloneCredentialDo) First() (*models.CodeRepositoryCloneCredential, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryCloneCredential), nil
	}
}

func (c codeRepositoryCloneCredentialDo) Take() (*models.CodeRepositoryCloneCredential, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryCloneCredential), nil
	}
}

func (c codeRepositoryCloneCredentialDo) Last() (*models.CodeRepositoryCloneCredential, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryCloneCredential), nil
	}
}

func (c codeRepositoryCloneCredentialDo) Find() ([]*models.CodeRepositoryCloneCredential, error) {
	result, err := c.DO.Find()
	return result.([]*models.CodeRepositoryCloneCredential), err
}

func (c codeRepositoryCloneCredentialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.CodeRepositoryCloneCredential, err error) {
	buf := make([]*models.CodeRepositoryCloneCredential, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c codeRepositoryCloneCredentialDo) FindInBatches(result *[]*models.CodeRepositoryCloneCredential, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c codeRepositoryCloneCredentialDo) Attrs(attrs ...field.AssignExpr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c codeRepositoryCloneCredentialDo) Assign(attrs ...field.AssignExpr) *codeRepositoryCloneCredentialDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c codeRepositoryCloneCredentialDo) Joins(fields ...field.RelationField) *codeRepositoryCloneCredentialDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c codeRepositoryCloneCredentialDo) Preload(fields ...field.RelationField) *codeRepositoryCloneCredentialDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c codeRepositoryCloneCredentialDo) FirstOrInit() (*models.CodeRepositoryCloneCredential, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryCloneCredential), nil
	}
}

func (c codeRepositoryCloneCredentialDo) FirstOrCreate() (*models.CodeRepositoryCloneCredential, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.CodeRepositoryCloneCredential), nil
	}
}

func (c codeRepositoryCloneCredentialDo) FindByPage(offset int, limit int) (result []*models.CodeRepositoryCloneCredential, count int64, err error) {
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

func (c codeRepositoryCloneCredentialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c codeRepositoryCloneCredentialDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c codeRepositoryCloneCredentialDo) Delete(models ...*models.CodeRepositoryCloneCredential) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *codeRepositoryCloneCredentialDo) withDO(do gen.Dao) *codeRepositoryCloneCredentialDo {
	c.DO = *do.(*gen.DO)
	return c
}
