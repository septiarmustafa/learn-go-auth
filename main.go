package main

import (
	"belajar-auth/internal/api"
	"belajar-auth/internal/component"
	"belajar-auth/internal/config"
	"belajar-auth/internal/middleware"
	"belajar-auth/internal/repository"
	"belajar-auth/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)
	cacheConnection := component.GetCacheConnection()

	userRespository := repository.NewUser(dbConnection)

	emailService := service.NewEmail(cnf)
	userService := service.NewUser(userRespository, cacheConnection, emailService)

	authMid := middleware.Authenticate(userService)

	app := fiber.New()
	api.NewAuth(app, userService, authMid)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
