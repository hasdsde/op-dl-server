package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserVersionEvent = "user_version_event"

// UserVersionEvent mapped from table <user_version_event>
type UserVersionEvent struct {
	ID             int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string" form:"id"`
	UserID         int64          `gorm:"column:user_id;type:int(11);comment:用户id" json:"userId" form:"userId"`                            // 用户id
	VersionEventID int64          `gorm:"column:version_event_id;type:int(11);comment:版本活动id" json:"versionEventId" form:"versionEventId"` // 版本活动id
	Todo           int64          `gorm:"column:todo;type:int(11);comment:已完成数量" json:"todo" form:"todo"`                                  // 已完成数量
	CreatedAt      time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName UserVersionEvent's table name
func (*UserVersionEvent) TableName() string {
	return TableNameUserVersionEvent
}
