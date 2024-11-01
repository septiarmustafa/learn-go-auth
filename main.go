package main

import (
	"belajar-auth/internal/component"
	"belajar-auth/internal/config"
	"belajar-auth/internal/repository"
	"belajar-auth/internal/service"
)

func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)
	cacheConnection := component.GetCacheConnection()

	userRespository := repository.NewUser(dbConnection)
	userService := service.NewUser(userRespository, cacheConnection)	
}
