package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserPrivate = "user_private"

// UserPrivate mapped from table <user_private>
type UserPrivate struct {
	ID         int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	UserID     int64          `gorm:"column:user_id;type:int(11)" json:"userId"`
	TodoNum    int64          `gorm:"column:todo_num;type:int(11);comment:需要数量" json:"todoNum"`     // 需要数量
	Todo       int64          `gorm:"column:todo;type:int(11);comment:已完成数量" json:"todo"`           // 已完成数量
	StartTime  int64          `gorm:"column:start_time;type:int(11);comment:开始时间" json:"startTime"` // 开始时间
	EndTime    int64          `gorm:"column:end_time;type:int(11);comment:结束时间" json:"endTime"`     // 结束时间
	CreatedAt  time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	AutoUpdate string         `gorm:"column:auto_update;type:varchar(10);default:0;comment:自动更新周期" json:"autoUpdate"` // 自动更新周期
	UpdateHour int64          `gorm:"column:update_hour;type:int(11);comment:自动更新时间 0或4" json:"updateHour"`           // 自动更新时间 0或4
}

// TableName UserPrivate's table name
func (*UserPrivate) TableName() string {
	return TableNameUserPrivate
}
