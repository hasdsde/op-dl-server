package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameVersionTag = "version_tag"

// VersionTag mapped from table <version_tag>
type VersionTag struct {
	ID        int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	VersionID int64          `gorm:"column:version_id;type:int(11);comment:版本ID" json:"versionId"` // 版本ID
	TagID     int64          `gorm:"column:tag_id;type:int(11);comment:标签ID" json:"tagId"`         // 标签ID
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`
}

// TableName VersionTag's table name
func (*VersionTag) TableName() string {
	return TableNameVersionTag
}
