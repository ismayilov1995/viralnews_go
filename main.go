package main

import (
	"fiber_news/models"
	"fiber_news/routers"
	"fiber_news/utils"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	htmlt "github.com/gofiber/template/html"
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
	engine := htmlt.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	initDB()

	routers.StaticRouters(app)
	routers.ApiRouters(app)

	app.Get("/reset", func(c *fiber.Ctx) error {
		resetDB()
		return c.SendString("db is rested")
	})

	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if err := app.Listen(":" + port); err != nil {
		log.Panicf("server not started")
	}
}
