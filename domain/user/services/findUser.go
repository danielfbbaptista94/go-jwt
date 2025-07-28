package services

import (
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"
	"go-jwt/models/user/repositories"
)

func (ud *userDomainService) FindUser(
	email string,
) (domain.UserDomainInterface, *errorhandler.ErrorHandler) {

	repository := repositories.NewUserRepository()

	userDomain, err := repository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return userDomain, nil
}
