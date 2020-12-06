package user

import (
	"golang-skeleton/domain/user/entity"
	"golang-skeleton/dto"
	"golang-skeleton/model"
	"golang-skeleton/repository"
	"golang-skeleton/utils"
)

type IUserService interface {
	GetUserByID(userID uint64) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	CreateUser(user *dto.CreateUserRequest) (*entity.User, error)
}

type userServiceImpl struct {
	repo repository.IRepository
}

func NewUserDomain(repo repository.IRepository) IUserService {
	return &userServiceImpl{repo: repo}
}

func (u *userServiceImpl) GetUserByID(userID uint64) (*entity.User, error) {
	return u.repo.GetUserByID(userID)
}

func (u *userServiceImpl) GetUserByUsername(username string) (*entity.User, error) {
	return u.repo.GetUserByUsername(username)
}

func (u *userServiceImpl) CreateUser(user *dto.CreateUserRequest) (*entity.User, error) {
	passwordHashed, err := utils.HashBCrypt(user.Password)
	if err != nil {
		return nil, err
	}
	return u.repo.CreateUser(user, passwordHashed)
}
