package repositories

import (
	"go-jwt/configuration/initializers"
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"
	"go-jwt/models/user/entities"

	"gorm.io/gorm"
)

func NewUserRepository() UserRepository {
	return &userRepository{
		db: initializers.DB,
	}
}

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	CreateUser(
		userDomain domain.UserDomainInterface,
	) (domain.UserDomainInterface, *errorhandler.ErrorHandler)
	FindUserByEmail(email string) (domain.UserDomainInterface, *errorhandler.ErrorHandler)
}

func (u *userRepository) toModel(userDomain domain.UserDomainInterface) *entities.UserModel {
	return &entities.UserModel{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
	}
}
