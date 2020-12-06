package user

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-skeleton/model"
	"golang-skeleton/repository"
	"testing"
)

func TestGetUserByID_Success(t *testing.T) {
	user := &model.User{
		BaseModel: model.BaseModel{
			ID: 1,
		},
		Username: "username",
		Name:     "name",
		Email:    "email",
		IsActive: true,
	}
	mockRepo := new(repository.MockIRepository)
	mockRepo.On("GetUserByID", mock.Anything).Return(user, nil)

	userService := &userServiceImpl{repo: mockRepo}

	result, err := userService.GetUserByID(1)

	assert.NotNil(t, result)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Username, result.Username)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.IsActive, result.IsActive)
}
