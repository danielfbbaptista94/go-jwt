package services

import (
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"
)

func (ud *userDomainService) CreateUser(
	userDomain domain.UserDomainInterface,
) *errorhandler.ErrorHandler {

	userDomain.EncryptPassword()
	_, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return err
	}
	return nil
}
