package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDaily = "daily"

// Daily mapped from table <daily>
type Daily struct {
	ID            int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	UserID        int64          `gorm:"column:user_id;type:int(11);comment:用户ID" json:"userId"`             // 用户ID
	UID           string         `gorm:"column:uid;type:varchar(10);comment:游戏ID" json:"uid"`                // 游戏ID
	HomeCoin      int64          `gorm:"column:home_coin;type:int(11);comment:洞天宝钱" json:"homeCoin"`         // 洞天宝钱
	TaskNum       int64          `gorm:"column:task_num;type:int(11);comment:每日委托数量" json:"taskNum"`         // 每日委托数量
	TaskFinished  int64          `gorm:"column:task_finished;type:int(11);comment:委托奖励" json:"taskFinished"` // 委托奖励
	Resin         int64          `gorm:"column:resin;type:int(11);comment:树脂" json:"resin"`                  // 树脂
	Expeditions   int64          `gorm:"column:expeditions;type:int(11);comment:派遣" json:"expeditions"`      // 派遣
	Transformer   int64          `gorm:"column:transformer;type:int(11);comment:参变" json:"transformer"`      // 参变
	ResinDiscount int64          `gorm:"column:resin_discount;type:int(11);comment:周本" json:"resinDiscount"` // 周本
	CreatedAt     time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName Daily's table name
func (*Daily) TableName() string {
	return TableNameDaily
}
