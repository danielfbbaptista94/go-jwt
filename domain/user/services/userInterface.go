package services

import (
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(domain.UserDomainInterface) *errorhandler.ErrorHandler
	FindUser(string) (domain.UserDomainInterface, *errorhandler.ErrorHandler)
}
