package models

import (
	"golang_gateway/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name     string
	Password string
	Phone    string
}

func (receiver *UserBasic) TableName() string {
	return "user_basic"
}

func GetUser(id int) []*UserBasic {
	userList := make([]*UserBasic, 10)
	utils.DB.First(&userList, "id = ?", id)
	return userList
}
