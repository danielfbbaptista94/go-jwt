package services

import (
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"
	"go-jwt/models/user/repositories"
)

func (ud *userDomainService) CreateUser(
	userDomain domain.UserDomainInterface,
) *errorhandler.ErrorHandler {

	repository := repositories.NewUserRepository()

	userDomain.EncryptPassword()
	_, err := repository.CreateUser(userDomain)
	if err != nil {
		return err
	}
	return nil
}
