package repository

import (
	"golang-skeleton/domain/user/entity"
	"golang-skeleton/dto"
	"golang-skeleton/model"
	"golang-skeleton/repository/transform"
)

type IUserRepo interface {
	GetUserByID(userID uint64) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	CreateUser(userDto *dto.CreateUserRequest, password string) (*entity.User, error)
}

func (r *repositoryImpl) GetUserByID(userID uint64) (*entity.User, error) {
	var user = &model.User{}
	err := r.db.First(user, userID).Error
	return transform.UserModelAsEntity(user), err
}

func (r *repositoryImpl) GetUserByUsername(username string) (*entity.User, error) {
	var user = &model.User{}
	err := r.db.Where("username = ?", username).First(user).Error
	return transform.UserModelAsEntity(user), err
}

func (r *repositoryImpl) CreateUser(userDto *dto.CreateUserRequest, password string) (*entity.User, error) {
	user := &model.User{
		Username: userDto.Username,
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: password,
		IsActive: true,
	}
	err := r.db.Create(user).Error
	return transform.UserModelAsEntity(user), err
}
