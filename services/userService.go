package services

import (
	"go-jwt/dtos"
	"go-jwt/models"
	"go-jwt/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return UserService{
		userRepository: repositories.NewUserRepository(),
	}
}

func (u *UserService) Signup(userDTO *dtos.UserDTO) error {
	user := models.UserModel{Email: userDTO.Email, Password: userDTO.Password}

	err := u.userRepository.Save(&user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) FindByEmail(email string) (dtos.UserDTO, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return dtos.UserDTO{}, err
	}

	userDTO := dtos.UserDTO{Id: int(user.ID), Email: user.Email, Password: user.Password}

	return userDTO, nil
}
