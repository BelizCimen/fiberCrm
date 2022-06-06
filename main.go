package main

import (
	"basicCrm/database"
	"basicCrm/lead"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Fail to connect database")
	}
	fmt.Println("Connected to database")
	database.DBConn.AutoMigrate(&lead.Lead{})

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(5432)
	defer database.DBConn.Close()

}
