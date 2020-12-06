package transform

import (
	"golang-skeleton/domain/user/entity"
	"golang-skeleton/dto"
)

func UserEntityAsDTO(user *dto.User) *entity.User {
	return &entity.User{
		UserID:   user.ID,
		Username: user.Name,
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
	}
}

func UserDTOAsEntity(user *entity.User) *dto.User {
	return &dto.User{
		ID:       0,
		Name:     "",
		Email:    "",
		IsActive: false,
	}
}

func ToGetUserResponseDTO(u *entity.User) *dto.GetUserResponse {
	return &dto.GetUserResponse{
		User: &dto.User{
			ID:       u.UserID,
			Name:     u.Name,
			Email:    u.Email,
			IsActive: u.IsActive,
		},
	}
}

func ToCreateUserResponseDTO(u *entity.User) *dto.CreateUserResponse {
	return &dto.CreateUserResponse{
		User: &dto.User{
			ID:       u.UserID,
			Name:     u.Name,
			Email:    u.Email,
			IsActive: u.IsActive,
		},
	}
}

func ToUserLoginResponseDTO(u *entity.User, token string) *dto.UserLoginResponse {
	return &dto.UserLoginResponse{
		Username: u.Username,
		Email:    u.Email,
		Name:     u.Name,
		Token:    token,
	}
}
