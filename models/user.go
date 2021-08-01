package models

import (
	"fiber_news/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Avatar   string
	Rating   int
}

func (u *User) Create() (*User, error) {
	if res := utils.DbConn.Create(&u); res.Error != nil {
		return nil, res.Error
	}
	return u, nil
}

func (u *User) LoadAll() *[]User {
	var users *[]User
	utils.DbConn.Find(&users)
	return users
}

func (u *User) Load(id interface{}) (*User, error) {
	if res := utils.DbConn.First(&u, id); res.Error != nil {
		return nil, res.Error
	} else {
		return u, nil
	}
}

func (u *User) Delete(id interface{}) (string, error) {
	if err := utils.DbConn.First(&u, id).Error; err != nil {
		return "", err
	}
	utils.DbConn.Delete(&u)
	return "user deleted", nil
}
