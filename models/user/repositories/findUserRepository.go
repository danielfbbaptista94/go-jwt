package repositories

import (
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"
	"go-jwt/models/user/entities"
)

func (u *userRepository) FindUserByEmail(email string) (domain.UserDomainInterface, *errorhandler.ErrorHandler) {
	var model entities.UserModel
	u.db.First(&model, "email = ?", email)
	if model.ID == 0 {
		return nil, errorhandler.NewNotFoundError("user not found")
	}

	user := domain.NewUserDomain(model.Email, model.Password)
	return user, nil
}
