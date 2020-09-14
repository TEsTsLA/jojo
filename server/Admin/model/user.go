package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"default:'galeone'"`
	Age      int    `gorm:"default:123"`
	Birthday time.Time
	//Email        string  `gor:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	//Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	//Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	//IgnoreMe     int     `gorm:m:"type:varchar(100);unique_index"`
	//Role         string  `gorm:"size:255"`        // 设置字段大小为255
	//MemberNumber *string `gorm"-"`               // 忽略本字段
}

func NewUser() *User {
	return &User{}
}
func Create() {

}
