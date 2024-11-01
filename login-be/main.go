package main

import (
	"fmt"
	"login-be/internal/app/users"
	"login-be/internal/app/users/controller"
	"login-be/internal/config"
	"login-be/internal/route"
)

func main() {
	viperConf := config.NewViper()
	app := config.NewFiber(viperConf)
	db := config.NewDatabase(viperConf)
	validate := config.NewValidator(viperConf)

	repoUsers := users.NewUsersRepository()
	usecaseUsers := users.NewUsersUsecase(db, validate, repoUsers, viperConf)
	controllerUsers := controller.NewUsersController(usecaseUsers, viperConf)
	// authMiddleware := middleware.NewAuth(usecaseUsers)

	routeConfig := route.RouteConfig{
		App:             app,
		UsersController: controllerUsers,
		// AuthMiddleware:  authMiddleware,
	}

	routeConfig.Setup()
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://127.0.0.1:3000/",      // Mengizinkan semua origin
	// 	AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS", // Metode yang diizinkan
	// 	AllowHeaders:     "Content-Type, Content-Length, Date, Authorization, Origin, Accept",
	// 	AllowCredentials: true,
	// }))

	app.Listen(fmt.Sprintf(":%d", viperConf.GetInt("web.port")))
}
