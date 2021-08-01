package main

import (
	"fiber_news/models"
	"fiber_news/routers"
	"fiber_news/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func initDB() {
	var err error
	dsn := "host=localhost user=postgres password=7090698 dbname=fiber port=5430 sslmode=disable TimeZone=Asia/Shanghai"
	if utils.DbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Panicln("not connected to db")
	}
	migrate()
	fmt.Println("connected to db")
}

func migrate()  {
	a := []interface{}{&models.News{}, &models.User{}}
	utils.DbConn.AutoMigrate(a...)
}

func resetDB()  {
	a := []interface{}{&models.News{}, &models.User{}}
	utils.DbConn.Migrator().DropTable(a...)
	utils.DbConn.AutoMigrate(a...)
}

func main() {
	app := fiber.New()
	initDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/reset", func(c *fiber.Ctx) error {
		resetDB()
		return c.SendString("db is rested")
	})

	routers.ApiRouters(app)

	if err := app.Listen(":8080"); err != nil {
		log.Panicf("server not started")
	}
}
