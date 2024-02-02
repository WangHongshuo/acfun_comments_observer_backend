// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameArticle = "article"

// Article mapped from table <article>
type Article struct {
	Aid             int64 `gorm:"column:aid;primaryKey;comment:Article ID" json:"aid"`                                          // Article ID
	LastFloorNumber int32 `gorm:"column:last_floor_number;not null;comment:Last Comment Floor Number" json:"last_floor_number"` // Last Comment Floor Number
	IsCompleted     bool  `gorm:"column:is_completed;not null;comment:Is get all comments" json:"is_completed"`                 // Is get all comments
}

// TableName Article's table name
func (*Article) TableName() string {
	return TableNameArticle
}
