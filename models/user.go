package models

import (
	"fiber_news/utils"
	"fmt"

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
	// utils.DbConn.Exec("DELETE FROM users WHERE id=?", id)
	utils.DbConn.Delete(&u)
	utils.DbConn.Where("author_id", u.ID).Delete(&News{})
	return "user deleted", nil
}

func (u *User) Seed() *[]User {
	var users []User
	for i := 0; i < 10; i++ {
		users = append(users, User{
			FullName: fmt.Sprintf("Dummy User %v", i),
			Email:    fmt.Sprintf("dummy%v@gmail.com", i),
			Password: "7090698"})
	}
	utils.DbConn.Create(&users)
	return &users
}

func (u *User) Reset() {
	utils.DbConn.Where("id IS NOT NULL").Delete(&u)
}
