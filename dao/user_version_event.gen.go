// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"hasdsd.cn/op-dl-server/model"
)

func newUserVersionEvent(db *gorm.DB, opts ...gen.DOOption) userVersionEvent {
	_userVersionEvent := userVersionEvent{}

	_userVersionEvent.userVersionEventDo.UseDB(db, opts...)
	_userVersionEvent.userVersionEventDo.UseModel(&model.UserVersionEvent{})

	tableName := _userVersionEvent.userVersionEventDo.TableName()
	_userVersionEvent.ALL = field.NewAsterisk(tableName)
	_userVersionEvent.ID = field.NewInt64(tableName, "id")
	_userVersionEvent.UserID = field.NewInt64(tableName, "user_id")
	_userVersionEvent.VersionDetailID = field.NewInt64(tableName, "version_detail_id")
	_userVersionEvent.Todo = field.NewInt64(tableName, "todo")
	_userVersionEvent.CreatedAt = field.NewTime(tableName, "created_at")
	_userVersionEvent.DeletedAt = field.NewField(tableName, "deleted_at")
	_userVersionEvent.UpdatedAt = field.NewTime(tableName, "updated_at")

	_userVersionEvent.fillFieldMap()

	return _userVersionEvent
}

type userVersionEvent struct {
	userVersionEventDo

	ALL             field.Asterisk
	ID              field.Int64
	UserID          field.Int64 // 用户id
	VersionDetailID field.Int64 // 版本活动id
	Todo            field.Int64 // 已完成数量
	CreatedAt       field.Time
	DeletedAt       field.Field
	UpdatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (u userVersionEvent) Table(newTableName string) *userVersionEvent {
	u.userVersionEventDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userVersionEvent) As(alias string) *userVersionEvent {
	u.userVersionEventDo.DO = *(u.userVersionEventDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userVersionEvent) updateTableName(table string) *userVersionEvent {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UserID = field.NewInt64(table, "user_id")
	u.VersionDetailID = field.NewInt64(table, "version_detail_id")
	u.Todo = field.NewInt64(table, "todo")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")

	u.fillFieldMap()

	return u
}

func (u *userVersionEvent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userVersionEvent) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 7)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["version_detail_id"] = u.VersionDetailID
	u.fieldMap["todo"] = u.Todo
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
}

func (u userVersionEvent) clone(db *gorm.DB) userVersionEvent {
	u.userVersionEventDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userVersionEvent) replaceDB(db *gorm.DB) userVersionEvent {
	u.userVersionEventDo.ReplaceDB(db)
	return u
}

type userVersionEventDo struct{ gen.DO }

type IUserVersionEventDo interface {
	gen.SubQuery
	Debug() IUserVersionEventDo
	WithContext(ctx context.Context) IUserVersionEventDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserVersionEventDo
	WriteDB() IUserVersionEventDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserVersionEventDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserVersionEventDo
	Not(conds ...gen.Condition) IUserVersionEventDo
	Or(conds ...gen.Condition) IUserVersionEventDo
	Select(conds ...field.Expr) IUserVersionEventDo
	Where(conds ...gen.Condition) IUserVersionEventDo
	Order(conds ...field.Expr) IUserVersionEventDo
	Distinct(cols ...field.Expr) IUserVersionEventDo
	Omit(cols ...field.Expr) IUserVersionEventDo
	Join(table schema.Tabler, on ...field.Expr) IUserVersionEventDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserVersionEventDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserVersionEventDo
	Group(cols ...field.Expr) IUserVersionEventDo
	Having(conds ...gen.Condition) IUserVersionEventDo
	Limit(limit int) IUserVersionEventDo
	Offset(offset int) IUserVersionEventDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserVersionEventDo
	Unscoped() IUserVersionEventDo
	Create(values ...*model.UserVersionEvent) error
	CreateInBatches(values []*model.UserVersionEvent, batchSize int) error
	Save(values ...*model.UserVersionEvent) error
	First() (*model.UserVersionEvent, error)
	Take() (*model.UserVersionEvent, error)
	Last() (*model.UserVersionEvent, error)
	Find() ([]*model.UserVersionEvent, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserVersionEvent, err error)
	FindInBatches(result *[]*model.UserVersionEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserVersionEvent) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserVersionEventDo
	Assign(attrs ...field.AssignExpr) IUserVersionEventDo
	Joins(fields ...field.RelationField) IUserVersionEventDo
	Preload(fields ...field.RelationField) IUserVersionEventDo
	FirstOrInit() (*model.UserVersionEvent, error)
	FirstOrCreate() (*model.UserVersionEvent, error)
	FindByPage(offset int, limit int) (result []*model.UserVersionEvent, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserVersionEventDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userVersionEventDo) Debug() IUserVersionEventDo {
	return u.withDO(u.DO.Debug())
}

func (u userVersionEventDo) WithContext(ctx context.Context) IUserVersionEventDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userVersionEventDo) ReadDB() IUserVersionEventDo {
	return u.Clauses(dbresolver.Read)
}

func (u userVersionEventDo) WriteDB() IUserVersionEventDo {
	return u.Clauses(dbresolver.Write)
}

func (u userVersionEventDo) Session(config *gorm.Session) IUserVersionEventDo {
	return u.withDO(u.DO.Session(config))
}

func (u userVersionEventDo) Clauses(conds ...clause.Expression) IUserVersionEventDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userVersionEventDo) Returning(value interface{}, columns ...string) IUserVersionEventDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userVersionEventDo) Not(conds ...gen.Condition) IUserVersionEventDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userVersionEventDo) Or(conds ...gen.Condition) IUserVersionEventDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userVersionEventDo) Select(conds ...field.Expr) IUserVersionEventDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userVersionEventDo) Where(conds ...gen.Condition) IUserVersionEventDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userVersionEventDo) Order(conds ...field.Expr) IUserVersionEventDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userVersionEventDo) Distinct(cols ...field.Expr) IUserVersionEventDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userVersionEventDo) Omit(cols ...field.Expr) IUserVersionEventDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userVersionEventDo) Join(table schema.Tabler, on ...field.Expr) IUserVersionEventDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userVersionEventDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserVersionEventDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userVersionEventDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserVersionEventDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userVersionEventDo) Group(cols ...field.Expr) IUserVersionEventDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userVersionEventDo) Having(conds ...gen.Condition) IUserVersionEventDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userVersionEventDo) Limit(limit int) IUserVersionEventDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userVersionEventDo) Offset(offset int) IUserVersionEventDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userVersionEventDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserVersionEventDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userVersionEventDo) Unscoped() IUserVersionEventDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userVersionEventDo) Create(values ...*model.UserVersionEvent) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userVersionEventDo) CreateInBatches(values []*model.UserVersionEvent, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userVersionEventDo) Save(values ...*model.UserVersionEvent) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userVersionEventDo) First() (*model.UserVersionEvent, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVersionEvent), nil
	}
}

func (u userVersionEventDo) Take() (*model.UserVersionEvent, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVersionEvent), nil
	}
}

func (u userVersionEventDo) Last() (*model.UserVersionEvent, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVersionEvent), nil
	}
}

func (u userVersionEventDo) Find() ([]*model.UserVersionEvent, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserVersionEvent), err
}

func (u userVersionEventDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserVersionEvent, err error) {
	buf := make([]*model.UserVersionEvent, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userVersionEventDo) FindInBatches(result *[]*model.UserVersionEvent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userVersionEventDo) Attrs(attrs ...field.AssignExpr) IUserVersionEventDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userVersionEventDo) Assign(attrs ...field.AssignExpr) IUserVersionEventDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userVersionEventDo) Joins(fields ...field.RelationField) IUserVersionEventDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userVersionEventDo) Preload(fields ...field.RelationField) IUserVersionEventDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userVersionEventDo) FirstOrInit() (*model.UserVersionEvent, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVersionEvent), nil
	}
}

func (u userVersionEventDo) FirstOrCreate() (*model.UserVersionEvent, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserVersionEvent), nil
	}
}

func (u userVersionEventDo) FindByPage(offset int, limit int) (result []*model.UserVersionEvent, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userVersionEventDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userVersionEventDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userVersionEventDo) Delete(models ...*model.UserVersionEvent) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userVersionEventDo) withDO(do gen.Dao) *userVersionEventDo {
	u.DO = *do.(*gen.DO)
	return u
}
