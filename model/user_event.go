package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserEvent = "user_event"

// UserEvent mapped from table <user_event>
type UserEvent struct {
	ID        int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string" form:"id"`
	UserID    int64          `gorm:"column:user_id;type:int(11);comment:用户id" json:"userId" form:"userId"`    // 用户id
	EventID   int64          `gorm:"column:event_id;type:int(11);comment:活动id" json:"eventId" form:"eventId"` // 活动id
	Todo      string         `gorm:"column:todo;type:varchar(100);comment:已经完成数量" json:"todo" form:"todo"`    // 已经完成数量
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`

	Event *Event `gorm:"foreignKey:id;references:event_id" json:"event"`
}

// TableName UserEvent's table name
func (*UserEvent) TableName() string {
	return TableNameUserEvent
}
