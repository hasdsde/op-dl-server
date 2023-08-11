package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameVersionEvent = "version_event"

// VersionEvent mapped from table <version_event>
type VersionEvent struct {
	ID         int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	VersionNum int64          `gorm:"column:version_num;type:int(11);comment:版本号" json:"versionNum"` // 版本号
	Title      string         `gorm:"column:title;type:varchar(100);comment:标题" json:"title"`        // 标题
	Tag        int64          `gorm:"column:tag;type:int(11);comment:标签" json:"tag"`                 // 标签
	Content    string         `gorm:"column:content;type:text;comment:内容" json:"content"`            // 内容
	Level      int64          `gorm:"column:level;type:int(11);comment:级别" json:"level"`             // 级别
	CreatedAt  time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`
	StartTime  time.Time      `gorm:"column:start_time;type:datetime;comment:开始时间" json:"startTime"` // 开始时间
	EndTime    time.Time      `gorm:"column:end_time;type:datetime;comment:结束时间" json:"endTime"`     // 结束时间
	TodoNum    int64          `gorm:"column:todo_num;type:int(11);comment:todo空位" json:"todoNum"`    // todo空位
}

// TableName VersionEvent's table name
func (*VersionEvent) TableName() string {
	return TableNameVersionEvent
}
