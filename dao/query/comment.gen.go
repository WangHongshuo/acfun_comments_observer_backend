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

	"github.com/WangHongshuo/acfuncommentsspider-go/dao/model"
)

func newComment(db *gorm.DB, opts ...gen.DOOption) comment {
	_comment := comment{}

	_comment.commentDo.UseDB(db, opts...)
	_comment.commentDo.UseModel(&model.Comment{})

	tableName := _comment.commentDo.TableName()
	_comment.ALL = field.NewAsterisk(tableName)
	_comment.Cid = field.NewInt64(tableName, "cid")
	_comment.Aid = field.NewInt64(tableName, "aid")
	_comment.FloorNumber = field.NewInt32(tableName, "floor_number")
	_comment.Comment = field.NewString(tableName, "comment")
	_comment.IsDel = field.NewBool(tableName, "is_del")
	_comment.HarmInfoReportCnt = field.NewInt64(tableName, "harm_info_report_cnt")

	_comment.fillFieldMap()

	return _comment
}

type comment struct {
	commentDo commentDo

	ALL               field.Asterisk
	Cid               field.Int64  // Comment ID
	Aid               field.Int64  // Article ID
	FloorNumber       field.Int32  // Floor Number
	Comment           field.String // Comment
	IsDel             field.Bool   // Is Delete
	HarmInfoReportCnt field.Int64  // The number of reports of harmful information

	fieldMap map[string]field.Expr
}

func (c comment) Table(newTableName string) *comment {
	c.commentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c comment) As(alias string) *comment {
	c.commentDo.DO = *(c.commentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *comment) updateTableName(table string) *comment {
	c.ALL = field.NewAsterisk(table)
	c.Cid = field.NewInt64(table, "cid")
	c.Aid = field.NewInt64(table, "aid")
	c.FloorNumber = field.NewInt32(table, "floor_number")
	c.Comment = field.NewString(table, "comment")
	c.IsDel = field.NewBool(table, "is_del")
	c.HarmInfoReportCnt = field.NewInt64(table, "harm_info_report_cnt")

	c.fillFieldMap()

	return c
}

func (c *comment) WithContext(ctx context.Context) *commentDo { return c.commentDo.WithContext(ctx) }

func (c comment) TableName() string { return c.commentDo.TableName() }

func (c comment) Alias() string { return c.commentDo.Alias() }

func (c comment) Columns(cols ...field.Expr) gen.Columns { return c.commentDo.Columns(cols...) }

func (c *comment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *comment) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["cid"] = c.Cid
	c.fieldMap["aid"] = c.Aid
	c.fieldMap["floor_number"] = c.FloorNumber
	c.fieldMap["comment"] = c.Comment
	c.fieldMap["is_del"] = c.IsDel
	c.fieldMap["harm_info_report_cnt"] = c.HarmInfoReportCnt
}

func (c comment) clone(db *gorm.DB) comment {
	c.commentDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c comment) replaceDB(db *gorm.DB) comment {
	c.commentDo.ReplaceDB(db)
	return c
}

type commentDo struct{ gen.DO }

func (c commentDo) Debug() *commentDo {
	return c.withDO(c.DO.Debug())
}

func (c commentDo) WithContext(ctx context.Context) *commentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentDo) ReadDB() *commentDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentDo) WriteDB() *commentDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentDo) Session(config *gorm.Session) *commentDo {
	return c.withDO(c.DO.Session(config))
}

func (c commentDo) Clauses(conds ...clause.Expression) *commentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentDo) Returning(value interface{}, columns ...string) *commentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentDo) Not(conds ...gen.Condition) *commentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentDo) Or(conds ...gen.Condition) *commentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentDo) Select(conds ...field.Expr) *commentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentDo) Where(conds ...gen.Condition) *commentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentDo) Order(conds ...field.Expr) *commentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentDo) Distinct(cols ...field.Expr) *commentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentDo) Omit(cols ...field.Expr) *commentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentDo) Join(table schema.Tabler, on ...field.Expr) *commentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *commentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentDo) RightJoin(table schema.Tabler, on ...field.Expr) *commentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentDo) Group(cols ...field.Expr) *commentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentDo) Having(conds ...gen.Condition) *commentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentDo) Limit(limit int) *commentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentDo) Offset(offset int) *commentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *commentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentDo) Unscoped() *commentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentDo) Create(values ...*model.Comment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentDo) CreateInBatches(values []*model.Comment, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentDo) Save(values ...*model.Comment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentDo) First() (*model.Comment, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) Take() (*model.Comment, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) Last() (*model.Comment, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) Find() ([]*model.Comment, error) {
	result, err := c.DO.Find()
	return result.([]*model.Comment), err
}

func (c commentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Comment, err error) {
	buf := make([]*model.Comment, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentDo) FindInBatches(result *[]*model.Comment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentDo) Attrs(attrs ...field.AssignExpr) *commentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentDo) Assign(attrs ...field.AssignExpr) *commentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentDo) Joins(fields ...field.RelationField) *commentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentDo) Preload(fields ...field.RelationField) *commentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentDo) FirstOrInit() (*model.Comment, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) FirstOrCreate() (*model.Comment, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Comment), nil
	}
}

func (c commentDo) FindByPage(offset int, limit int) (result []*model.Comment, count int64, err error) {
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

func (c commentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentDo) Delete(models ...*model.Comment) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentDo) withDO(do gen.Dao) *commentDo {
	c.DO = *do.(*gen.DO)
	return c
}
