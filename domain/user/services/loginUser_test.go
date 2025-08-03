package services

import (
	"github.com/stretchr/testify/assert"
	requestdto "go-jwt/controllers/requestDTO"
	errorhandler "go-jwt/errorHandler"
	mocks "go-jwt/tests/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserDomainService_FindUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	request := requestdto.LoginDTO{
		Email:    "teste@email.com",
		Password: "password",
	}

	var repo = mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repo)

	t.Run("when_exists_an_users_returns_success", func(t *testing.T) {
		userDomainMock := mocks.NewMockUserDomainInterface(ctrl)
		tokenReturn := "fake.jwt.token"

		// These expectations will be used by the service
		userDomainMock.EXPECT().ComparePassword(request.Password).Return(nil)
		userDomainMock.EXPECT().GenerateToken().Return(tokenReturn, nil)
		userDomainMock.EXPECT().GetEmail().Return(request.Email)

		repo.EXPECT().FindUserByEmail(request.Email).Return(userDomainMock, nil)

		user, token, err := service.Login(request)

		assert.Nil(t, err)
		assert.Equal(t, request.Email, user.GetEmail())
		assert.Equal(t, tokenReturn, token)
	})

	t.Run("when_not_found_an_users_returns_error", func(t *testing.T) {
		repo.EXPECT().FindUserByEmail(request.Email).Return(nil, errorhandler.NewNotFoundError("user not found"))
		user, _, err := service.Login(request)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})

	t.Run("when_invalid_password_returns_error", func(t *testing.T) {
		userDomainMock := mocks.NewMockUserDomainInterface(ctrl)

		// These expectations will be used by the service
		userDomainMock.EXPECT().ComparePassword(request.Password).Return(errorhandler.NewNotFoundError("invalid password"))

		repo.EXPECT().FindUserByEmail(request.Email).Return(userDomainMock, nil)
		_, _, err := service.Login(request)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "invalid password")
	})

	t.Run("when_create_token_returns_error", func(t *testing.T) {
		userDomainMock := mocks.NewMockUserDomainInterface(ctrl)

		// These expectations will be used by the service
		userDomainMock.EXPECT().ComparePassword(request.Password).Return(nil)
		userDomainMock.EXPECT().GenerateToken().Return("", errorhandler.NewBadRequestError("failed to create token"))

		repo.EXPECT().FindUserByEmail(request.Email).Return(userDomainMock, nil)
		_, _, err := service.Login(request)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "failed to create token")
	})
}
