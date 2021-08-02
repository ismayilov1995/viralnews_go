package main

import (
	"fiber_news/models"
	"fiber_news/routers"
	"fiber_news/utils"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic("enviroment not loaded")
	}
}

func initDB() {
	var err error
	dbp := os.Getenv("DB_PORT")
	dbh := os.Getenv("DB_HOST")
	dbu := os.Getenv("DB_USER")
	dbpwd := os.Getenv("DB_PWD")
	dbn := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Baku", dbh, dbu, dbpwd, dbn, dbp)
	if utils.DbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Panicln("not connected to db")
	}
	migrate()
	fmt.Println("connected to db")
}

func migrate() {
	a := []interface{}{&models.News{}, &models.User{}}
	utils.DbConn.AutoMigrate(a...)
}

func resetDB() {
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

	if err := app.Listen(":3000"); err != nil {
		log.Panicf("server not started")
	}
}
