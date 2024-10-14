package main

import (
	"log"
	"simple-clean-architecture/internal/handler"
	"simple-clean-architecture/internal/infra/config"
	"simple-clean-architecture/internal/repository"
	"simple-clean-architecture/internal/repository/database"
	"simple-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.NewConfig(".toml")
	if err != nil {
		log.Panicf("failed new config: %v\n", err)
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Panicf("failed new database: %v\n", err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	app := fiber.New(fiber.Config{
		ErrorHandler: fiber.DefaultErrorHandler,
	})
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			userHandler.Route(v1)
		}
	}

	log.Fatal(app.Listen(":8080"))
}
