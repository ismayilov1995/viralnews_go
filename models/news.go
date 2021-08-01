package models

import (
	"fiber_news/utils"
	"fmt"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Title     string `gorm:"not null;unique"`
	Content   string `gorm:"not null"`
	Thumbnail string
	AuthorID  int
	Author    User
	Read      int
}

func (n *News) Create() (*News, error) {
	if res := utils.DbConn.Create(&n); res.Error != nil {
		return nil, res.Error
	}
	return n, nil
}

func (n *News) LoadAll() []News {
	var news []News
	utils.DbConn.Find(&news)
	for i, _ := range news {
		utils.DbConn.Model(news[i]).Association("Author").Find(&news[i].Author)
	}
	return news
}

func (n *News) Load(id interface{}) *News {
	utils.DbConn.Find(&n, id)
	err := utils.DbConn.Model(n).Association("Author").Find(&n.Author)
	if err != nil {
		fmt.Println(err)
	}
	return n
}
