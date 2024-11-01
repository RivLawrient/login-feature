package users

import "time"

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,max=100,min=8"`
}

type UserResponse struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Token     string    `json:"token,omitempty"`
	IsGoogle  bool      `json:"is_google,omitempty"`
	IsGithub  bool      `json:"is_github,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

type RegisterUserGoogle struct {
	Name  string `validate:"required,max=100"`
	Email string `validate:"required,email"`
}

type LoginUserGoogle struct {
	Email string `validate:"required, email"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}
