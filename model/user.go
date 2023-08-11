package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id,string"`
	Name      string         `gorm:"column:name;type:varchar(100);comment:名字" json:"name"`         // 名字
	Password  string         `gorm:"column:password;type:varchar(100);comment:密码" json:"password"` // 密码
	Token     string         `gorm:"column:token;type:text;comment:token" json:"token"`            // token
	Role      string         `gorm:"column:role;type:varchar(100)" json:"role"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"-"`

	UserEvent        []*UserEvent        `gorm:"foreignKey:user_id;references:id" json:"userEvent"`
	UserPrivate      []*UserPrivate      `gorm:"foreignKey:user_id;references:id" json:"userPrivate"`
	UserVersionEvent []*UserVersionEvent `gorm:"foreignKey:user_id;references:id" json:"userVersionEvent"`
	Daily            []*Daily            `gorm:"foreignKey:user_id;references:id" json:"daily"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
