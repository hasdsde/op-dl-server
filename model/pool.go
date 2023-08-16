package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameCat = "pool"

type Pool struct {
	ID         int64     `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string" form:"id"`
	VersionNum int64     `gorm:"column:version_num;type:int(11);comment:版本号" json:"versionNum" form:"versionNum"` // 版本号
	Name       string    `gorm:"column:name;type:varchar(100);comment:名字" json:"name" form:"name"`                // 名字
	Img        string    `gorm:"column:img;type:varchar(256);comment:图片" json:"img" form:"img"`                   // 图片
	Star       string    `gorm:"column:star;type:int(11);comment:星级" json:"star" form:"star"`                     // 星级
	Type       string    `gorm:"column:type;varchar(11);comment:type" json:"type" form:"type"`                    // 类型
	StartTime  time.Time `gorm:"column:start_time;type:datetime;comment:开始时间" json:"startTime" form:"startTime"`  // 开始时间
	EndTime    time.Time `gorm:"column:end_time;type:datetime;comment:结束时间" json:"endTime" form:"endTime"`        // 结束时间

	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`
}

// TableName Version's table name
func (*Pool) TableName() string {
	return TableNameCat
}
