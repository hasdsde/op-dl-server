package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameVersion = "version"

// Version mapped from table <version>
type Version struct {
	ID        int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	Num       int64          `gorm:"column:num;type:int(11);not null;comment:版本号" json:"num"`       // 版本号
	StartTime time.Time      `gorm:"column:start_time;type:datetime;comment:开始时间" json:"startTime"` // 开始时间
	EndTime   time.Time      `gorm:"column:end_time;type:datetime;comment:结束时间" json:"endTime"`     // 结束时间
	Title     string         `gorm:"column:title;type:varchar(100);comment:标题" json:"title"`        // 标题
	Img       string         `gorm:"column:img;type:varchar(255);comment:图片" json:"img"`            // 图片
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`

	VersionTag   []*VersionTag   `gorm:"foreignKey:version_id;reference:num" json:"versionTag"`
	VersionEvent []*VersionEvent `gorm:"foreignKey:VersionNum;reference:num" json:"versionEvent"`
}

// TableName Version's table name
func (*Version) TableName() string {
	return TableNameVersion
}
