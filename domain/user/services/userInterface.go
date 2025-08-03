package services

import (
	requestdto "go-jwt/controllers/requestDTO"
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"
	"go-jwt/models/user/repositories"
)

func NewUserDomainService(
	userRepository repositories.UserRepository,
) UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepository repositories.UserRepository
}

type UserDomainService interface {
	CreateUser(domain.UserDomainInterface) *errorhandler.ErrorHandler
	Login(requestdto.LoginDTO) (domain.UserDomainInterface, string, *errorhandler.ErrorHandler)
	FindUser(string) (domain.UserDomainInterface, *errorhandler.ErrorHandler)
}
