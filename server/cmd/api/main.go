package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/maggickkw/ticket-master/config"
	"github.com/maggickkw/ticket-master/db"
	"github.com/maggickkw/ticket-master/handlers"
	"github.com/maggickkw/ticket-master/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketMaster",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.SeverPort))
}
