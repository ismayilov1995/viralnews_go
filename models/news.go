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
	Author    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

func (n *News) Load(id interface{}) (*News, error) {
	utils.DbConn.Find(&n, id)
	err := utils.DbConn.Model(n).Association("Author").Find(&n.Author)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (n *News) Delete(id interface{}) (string, error) {
	utils.DbConn.Find(&n, id)
	if err := utils.DbConn.Delete(&n); err != nil {
		return "", err.Error
	}
	return "news is removed", nil
}

func (n *News) Seed() *[]News {
	var news []News
	for i := 0; i < 10; i++ {
		news = append(news, News{
			Title:    fmt.Sprintf("This is dummy title %v", i),
			Content:  fmt.Sprintf("This is content %v", i),
			AuthorID: i + 1})
	}
	utils.DbConn.Create(&news)
	return &news
}

func (n *News) Reset() {
	utils.DbConn.Where("id IS NOT NULL").Delete(&n)
}
