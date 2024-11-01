package controller

import (
	"login-be/internal/app/users"
	"login-be/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type UsersController struct {
	Usecase *users.UsersUsecase
	Viper   *viper.Viper
}

func NewUsersController(usecase *users.UsersUsecase, viper *viper.Viper) *UsersController {
	return &UsersController{
		Usecase: usecase,
		Viper:   viper,
	}
}
func (c *UsersController) GetUser(ctx *fiber.Ctx) error {
	auth := ctx.Cookies("auth-token")

	resposne, err := c.Usecase.Verify(&users.VerifyUserRequest{Token: auth})
	if err != nil {
		error := err
		return ctx.Status(error.Code).JSON(model.WebResponse[any]{
			StatusCode: error.Code,
			Data:       nil,
			Errors:     error.Message,
		})
	}

	return ctx.JSON(model.WebResponse[any]{
		StatusCode: ctx.Response().StatusCode(),
		Data:       resposne,
	})
}

func (c *UsersController) RegisterEmail(ctx *fiber.Ctx) error {
	request := new(users.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		error := fiber.ErrBadRequest
		return ctx.Status(error.Code).JSON(model.WebResponse[any]{
			StatusCode: error.Code,
			Data:       nil,
			Errors:     "invalid request body",
		})
	}

	response, errRes := c.Usecase.CreateByEmail(request)
	if errRes != nil {
		error := errRes
		return ctx.Status(error.Code).JSON(model.WebResponse[any]{
			StatusCode: error.Code,
			Data:       nil,
			Errors:     error.Message,
		})
	}
	cookie := new(fiber.Cookie)

	cookie.Name = "auth-token"
	cookie.Value = response.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = true

	ctx.Cookie(cookie)

	return ctx.JSON(model.WebResponse[*users.UserResponse]{
		StatusCode: ctx.Response().StatusCode(),
		Data:       response,
	})
}

func (c *UsersController) LoginEmail(ctx *fiber.Ctx) error {
	request := new(users.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		error := fiber.ErrBadRequest
		return ctx.Status(error.Code).JSON(model.WebResponse[any]{
			StatusCode: error.Code,
			Data:       nil,
			Errors:     "invalid request body",
		})
	}

	response, errRer := c.Usecase.Login(request)
	if errRer != nil {
		error := errRer
		return ctx.Status(error.Code).JSON(model.WebResponse[any]{
			StatusCode: error.Code,
			Data:       nil,
			Errors:     error.Message,
		})
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "auth-token"
	cookie.Value = response.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = true

	ctx.Cookie(cookie)

	return ctx.JSON(model.WebResponse[*users.UserResponse]{
		StatusCode: ctx.Response().StatusCode(),
		Data:       response,
	})
}

func (c *UsersController) RegisterGoogle(ctx *fiber.Ctx) error {
	url := c.Usecase.RegisterGoogleHandle()
	return ctx.JSON(fiber.Map{
		"url": url,
	})
}

func (c *UsersController) RegisterCallbackGoogle(ctx *fiber.Ctx) error {
	code := ctx.FormValue("code")

	hasil, err := c.Usecase.GoogleCallbackHandle(code)
	if err != nil {
		error := err
		return ctx.Status(error.Code).JSON(model.WebResponse[any]{
			StatusCode: error.Code,
			Data:       nil,
			Errors:     error.Error(),
		})
	}
	request := &users.RegisterUserGoogle{
		Name:  hasil.Name,
		Email: hasil.Email,
	}

	register, err := c.Usecase.CreateByGoogle(request)
	if err != nil {
		error := err
		return ctx.Status(error.Code).JSON(model.WebResponse[any]{
			StatusCode: error.Code,
			Data:       nil,
			Errors:     error.Error(),
		})
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "auth-token"
	cookie.Value = register.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = true

	ctx.Cookie(cookie)

	return ctx.Redirect(c.Viper.GetString("ip.web") + "/home")
}

func (c *UsersController) LoginGoogle(ctx *fiber.Ctx) error {

	url := c.Usecase.LoginGoogleHandle()
	return ctx.JSON(fiber.Map{
		"url": url,
	})

}

func (c *UsersController) LoginCallbackGoogle(ctx *fiber.Ctx) error {
	code := ctx.FormValue("code")

	hasil, err := c.Usecase.LoginGoogleCallbackHandle(code)
	if err != nil {
		error := err
		return ctx.Status(error.Code).JSON(model.WebResponse[any]{
			StatusCode: error.Code,
			Data:       nil,
			Errors:     error.Error(),
		})
	}

	request := &users.LoginUserGoogle{
		Email: hasil.Email,
	}

	register, err := c.Usecase.LoginByGoogle(request)
	if err != nil {
		error := err
		return ctx.Status(error.Code).JSON(model.WebResponse[any]{
			StatusCode: error.Code,
			Data:       nil,
			Errors:     error.Error(),
		})
	}

	cookie := new(fiber.Cookie)

	cookie.Name = "auth-token"
	cookie.Value = register.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = true

	ctx.Cookie(cookie)

	return ctx.Redirect(c.Viper.GetString("ip.web") + "/home")
}

func (c *UsersController) Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie("auth-token")

	return ctx.JSON(fiber.Map{
		"url": c.Viper.GetString("ip.web"),
	})
}
