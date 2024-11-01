package users

import (
	"context"
	"encoding/json"
	"io"
	"login-be/internal/config"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type UsersUsecase struct {
	DB              *gorm.DB
	Validate        *validator.Validate
	UsersRepository *UsersRepository
	Viper           *viper.Viper
}

func NewUsersUsecase(db *gorm.DB, validate *validator.Validate, usersRepository *UsersRepository, viper *viper.Viper) *UsersUsecase {
	return &UsersUsecase{
		DB:              db,
		Validate:        validate,
		UsersRepository: usersRepository,
		Viper:           viper,
	}
}

func (u *UsersUsecase) CreateByEmail(request *RegisterUserRequest) (*UserResponse, *fiber.Error) {
	val := config.ValidateStruct(request)
	if val != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, val[0])
	}

	password := new(string)
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}
	*password = string(hash)

	user := &Users{
		ID:       uuid.New().String(),
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		IsGoogle: false,
		IsGithub: false,
		Token:    uuid.New().String(),
	}
	email := u.UsersRepository.FindByEmail(u.DB, &Users{}, request.Email)
	if email == nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, "Email is already used")
	}

	err = u.UsersRepository.Create(u.DB, user)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, strings.Split(err.Error(), ": ")[1])

	}
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     user.Token,
		IsGoogle:  user.IsGoogle,
		IsGithub:  user.IsGithub,
		CreatedAt: time.UnixMilli(user.CreatedAt),
	}, nil
}

func (u *UsersUsecase) Login(request *LoginUserRequest) (*UserResponse, *fiber.Error) {
	val := config.ValidateStruct(request)
	if val != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, val[0])
	}

	user := new(Users)
	email := u.UsersRepository.FindByEmail(u.DB, user, request.Email)
	if email != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, "Email or Password is invalid")
	}

	pass := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(request.Password))
	if pass != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, "Email or Password is invalid")
	}

	user.Token = uuid.New().String()
	save := u.UsersRepository.Update(u.DB, user)
	if save != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, "Something wrong")
	}

	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     user.Token,
		IsGoogle:  user.IsGoogle,
		IsGithub:  user.IsGithub,
		CreatedAt: time.UnixMilli(user.CreatedAt),
	}, nil
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "https://api.lawrients.my.id/auth/register/google/callback",                 // Ganti dengan redirect URI yang sesuai
		ClientID:     "1095824475598-snu63od692kqkteje6ta9svlcajg7tei.apps.googleusercontent.com", // Ganti dengan Client ID dari Google Console
		ClientSecret: "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",                                       // Ganti dengan Client Secret dari Google Console
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
	oauthStateString = "random"
)

func (u *UsersUsecase) RegisterGoogleHandle() string {
	return googleOauthConfig.AuthCodeURL(oauthStateString)
}

func (u *UsersUsecase) GoogleCallbackHandle(code string) (*GoogleUser, *fiber.Error) {

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, "")
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	body, _ := io.ReadAll(response.Body)
	var apiResponse GoogleUser
	json.Unmarshal(body, &apiResponse)

	return &apiResponse, nil
}

func (u *UsersUsecase) CreateByGoogle(request *RegisterUserGoogle) (*UserResponse, *fiber.Error) {

	check := u.UsersRepository.FindByEmail(u.DB, &Users{}, request.Email)
	if check == nil {
		user := new(Users)

		if request.Email != "" {
			user.Email = request.Email
			user.Token = uuid.New().String()
		}

		err := u.UsersRepository.UpdateByEmail(u.DB, user, request.Email, user.Token)
		if err != nil {
			return nil, fiber.NewError(fiber.ErrBadRequest.Code, strings.Split(err.Error(), ": ")[1])

		}

		return &UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Token:     user.Token,
			IsGoogle:  user.IsGoogle,
			IsGithub:  user.IsGithub,
			CreatedAt: time.UnixMilli(user.CreatedAt),
		}, nil
	}

	user := &Users{
		ID:       uuid.New().String(),
		Name:     request.Name,
		Email:    request.Email,
		IsGoogle: true,
		IsGithub: false,
		Token:    uuid.New().String(),
	}

	err := u.UsersRepository.Create(u.DB, user)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, strings.Split(err.Error(), ": ")[1])

	}
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     user.Token,
		IsGoogle:  user.IsGoogle,
		IsGithub:  user.IsGithub,
		CreatedAt: time.UnixMilli(user.CreatedAt),
	}, nil
}

var (
	googleOauthConfig2 = &oauth2.Config{
		RedirectURL:  "https://api.lawrients.my.id/auth/login/google/callback",                    // Ganti dengan redirect URI yang sesuai
		ClientID:     "1095824475598-snu63od692kqkteje6ta9svlcajg7tei.apps.googleusercontent.com", // Ganti dengan Client ID dari Google Console
		ClientSecret: "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",                                       // Ganti dengan Client Secret dari Google Console
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
	oauthStateString2 = "random"
)

func (u *UsersUsecase) LoginGoogleHandle() string {
	return googleOauthConfig2.AuthCodeURL(oauthStateString2)
}
func (u *UsersUsecase) LoginGoogleCallbackHandle(code string) (*GoogleUser, *fiber.Error) {

	token, err := googleOauthConfig2.Exchange(context.Background(), code)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, "")
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	body, _ := io.ReadAll(response.Body)
	var apiResponse GoogleUser
	json.Unmarshal(body, &apiResponse)

	return &apiResponse, nil
}
func (u *UsersUsecase) LoginByGoogle(request *LoginUserGoogle) (*UserResponse, *fiber.Error) {

	check := u.UsersRepository.FindByEmail(u.DB, &Users{}, request.Email)
	if check != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, "Email is not found")
	}

	user := new(Users)

	if request.Email != "" {
		user.Email = request.Email
		user.Token = uuid.New().String()
	}

	err := u.UsersRepository.UpdateByEmail(u.DB, user, request.Email, user.Token)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, strings.Split(err.Error(), ": ")[1])

	}

	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     user.Token,
		IsGoogle:  user.IsGoogle,
		IsGithub:  user.IsGithub,
		CreatedAt: time.UnixMilli(user.CreatedAt),
	}, nil
}

func (c *UsersUsecase) Verify(request *VerifyUserRequest) (*UserResponse, *fiber.Error) {
	err := c.Validate.Struct(request)

	if err != nil {
		validationError := err.(validator.ValidationErrors)
		var message string
		for _, field := range validationError {
			if len(strings.Split(field.Error(), ":")) > 1 {
				message += strings.Split(field.Error(), ":")[2]
				break
			}
		}
		return nil, fiber.NewError(fiber.ErrBadRequest.Code, message)
	}

	user := new(Users)
	err = c.UsersRepository.FindyByToken(c.DB, user, request.Token)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     user.Token,
		CreatedAt: time.UnixMilli(user.CreatedAt),
	}, nil

}
