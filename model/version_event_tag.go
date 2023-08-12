package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameVersionEventTag = "version_event_tag"

// VersionEventTag mapped from table <version_event_tag>
type VersionEventTag struct {
	ID             int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	VersionEventID int64          `gorm:"column:version_event_id;type:int(11)" json:"version_event_id" form:"versionEventId"`
	TagID          int64          `gorm:"column:tag_id;type:int(11)" json:"tagId" form:"tagId"`
	CreatedAt      time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`

	Tag *Tag `gorm:"foreignKey:id;references:tag_id" json:"tag"`
}

// TableName VersionEventTag's table name
func (*VersionEventTag) TableName() string {
	return TableNameVersionEventTag
}
