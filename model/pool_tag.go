package model

import (
	"gorm.io/gorm"
	"time"
)

const TablePoolTag = "pool_tag"

type PoolTag struct {
	ID        int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	PoolId    int64          `gorm:"column:pool_id;type:int(11);comment:卡池id" json:"poolId" form:"poolId"` // 版本ID
	TagID     int64          `gorm:"column:tag_id;type:int(11);comment:标签ID" json:"tagId" form:"tagId"`    // 标签ID
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`

	Tag *Tag `gorm:"foreignKey:id;references:tag_id" json:"tag"`
}

func (*PoolTag) TableName() string {
	return TablePoolTag
}
