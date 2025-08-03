package services

import (
	"github.com/stretchr/testify/assert"
	domain "go-jwt/domain/user"
	errorHandler "go-jwt/errorHandler"
	mocks "go-jwt/tests/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUserDomainService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var repo = mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repo)

	t.Run("when_entity_is_valid_returns_success", func(t *testing.T) {
		userDomain := domain.NewUserDomain("test@email.com", "123")

		repo.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		err := service.CreateUser(userDomain)
		if err != nil {
			t.FailNow()
			return
		}

		assert.Nil(t, err)
	})

	t.Run("when_entity_is_invalid_returns_error", func(t *testing.T) {
		userDomain := domain.NewUserDomain("test@email.com", "123")
		repo.EXPECT().CreateUser(userDomain).Return(nil, errorHandler.NewInternalError("error creating user"))

		err := service.CreateUser(userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error creating user")
	})
}
