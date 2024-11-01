package route

import (
	"login-be/internal/app/users/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type RouteConfig struct {
	App             *fiber.App
	UsersController *controller.UsersController
}

func (c *RouteConfig) Setup() {
	c.App.Use(cors.New(cors.Config{
		AllowOrigins:     "https://lawrients.my.id",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization, Origin, Accept",
		AllowCredentials: true,
	}))
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/auth/register", c.UsersController.RegisterEmail)
	c.App.Post("/auth/login", c.UsersController.LoginEmail)
	c.App.Get("/auth/register/google", c.UsersController.RegisterGoogle)
	c.App.Get("/auth/register/google/callback", c.UsersController.RegisterCallbackGoogle)

	c.App.Get("/auth/login/google", c.UsersController.LoginGoogle)
	c.App.Get("/auth/login/google/callback", c.UsersController.LoginCallbackGoogle)

	c.App.Get("/bye", c.UsersController.Logout)
}

func (c *RouteConfig) SetupAuthRoute() {

	c.App.Get("/api/users", c.UsersController.GetUser)
}
