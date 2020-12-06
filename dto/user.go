package dto

import "github.com/dgrijalva/jwt-go"

type JWTCustomClaims struct {
	Username string `json:"username"`
	ID       uint64 `json:"id"`
	jwt.StandardClaims
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TODO: add refresh token and expire
type UserLoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

type CreateUserResponse struct {
	User *User `json:"user"`
}
