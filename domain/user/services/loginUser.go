package services

import (
	requestdto "go-jwt/controllers/requestDTO"
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"
)

func (ud *userDomainService) Login(
	loginDTO requestdto.LoginDTO,
) (domain.UserDomainInterface, string, *errorhandler.ErrorHandler) {

	userDomain, err := ud.userRepository.FindUserByEmail(loginDTO.Email)
	if err != nil {
		return nil, "", err
	}

	errPassword := userDomain.ComparePassword(loginDTO.Password)
	if errPassword != nil {
		return nil, "", errPassword
	}

	tokenString, tokenErr := userDomain.GenerateToken()
	if tokenErr != nil {
		return nil, "", tokenErr
	}

	return userDomain, tokenString, nil
}
