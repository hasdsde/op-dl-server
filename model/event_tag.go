package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameEventTag = "event_tag"

// EventTag mapped from table <event_tag>
type EventTag struct {
	ID        int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string" form:"id"`
	EventID   int64          `gorm:"column:event_id;type:int(11)" json:"eventId" form:"eventId"`
	TagID     int64          `gorm:"column:tag_id;type:int(11)" json:"tagId" form:"tagId"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`

	Tag *Tag `gorm:"foreignKey:id;references:tag_id" json:"tag"`
}

// TableName EventTag's table name
func (*EventTag) TableName() string {
	return TableNameEventTag
}
