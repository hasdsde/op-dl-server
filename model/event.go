package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameEvent = "event"

// Event mapped from table <event>
type Event struct {
	ID        int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string" form:"id"`
	Title     string         `gorm:"column:title;type:varchar(100);comment:标题" json:"title" form:"title"`                 // 标题
	Primogems string         `gorm:"column:primogems;type:varchar(10);comment:原石数量" json:"primogems" form:"primogems"`    // 原石数量
	Award     string         `gorm:"column:award;type:varchar(100);comment:其他奖励" json:"award" form:"award"`               // 其他奖励
	PreTask   int64          `gorm:"column:pre_task;type:int(11);comment:前置任务，数字越大越长，0表示无" json:"preTask" form:"preTask"` // 前置任务，数字越大越长，0表示无
	CoPlay    int64          `gorm:"column:co_play;type:int(11);comment:联机游戏 0为否" json:"coPlay" form:"coPlay"`            // 联机游戏 0为否
	StartTime time.Time      `gorm:"column:start_time;type:datetime;comment:开始时间" json:"startTime" form:"startTime"`      // 开始时间
	EndTime   time.Time      `gorm:"column:end_time;type:datetime;comment:结束时间" json:"endTime" form:"endTime"`            // 结束时间
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	TodoNum   int64          `gorm:"column:todo_num;type:int(11);comment:todo空位" json:"todoNum" form:"todoNum"` // todo空位

	EventTag    []*EventTag    `gorm:"foreignKey:event_id;references:id" json:"eventTag"`
	EventDetail []*EventDetail `gorm:"foreignKey:event_id;reference:id" json:"eventDetail"`
}

// TableName Event's table name
func (*Event) TableName() string {
	return TableNameEvent
}
