package services

import (
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"
)

func (ud *userDomainService) FindUser(
	email string,
) (domain.UserDomainInterface, *errorhandler.ErrorHandler) {

	userDomain, err := ud.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return userDomain, nil
}
