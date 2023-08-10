// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTag = "tag"

// Tag mapped from table <tag>
type Tag struct {
	ID        int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	Name      string         `gorm:"column:name;type:varchar(100);comment:名字" json:"name"` // 名字
	Icon      string         `gorm:"column:icon;type:varchar(100)" json:"icon"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`
	Sort      int64          `gorm:"column:sort;type:int(11);comment:分类" json:"sort"` // 分类
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
