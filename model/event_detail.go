package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameEventDetail = "event_detail"

type EventDetail struct {
	ID        int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string" form:"id"`
	Title     string         `gorm:"column:title;type:varchar(100);comment:标题" json:"title" form:"title"` // 标题
	EventID   int64          `gorm:"column:event_id;type:int(11)" json:"eventId" form:"eventId"`
	Icon      string         `gorm:"column:icon;type:varchar(100)" json:"icon"`
	Primogems string         `gorm:"column:primogems;type:varchar(10);comment:原石数量" json:"primogems" form:"primogems"` // 原石数量
	Award     string         `gorm:"column:award;type:varchar(100);comment:其他奖励" json:"award" form:"award"`            // 其他奖励
	StartTime time.Time      `gorm:"column:start_time;type:datetime;comment:开始时间" json:"startTime" form:"startTime"`   // 开始时间
	EndTime   time.Time      `gorm:"column:end_time;type:datetime;comment:结束时间" json:"endTime" form:"endTime"`         // 结束时间
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

func (*EventDetail) TableName() string {
	return TableNameEventDetail
}
