package repositories

import (
	domain "go-jwt/domain/user"
	errorhandler "go-jwt/errorHandler"

	"gorm.io/gorm"
)

// CreateUser implements UserRepository.
func (u *userRepository) CreateUser(
	userDomain domain.UserDomainInterface,
) (domain.UserDomainInterface, *errorhandler.ErrorHandler) {
	model := u.toModel(userDomain)
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(model).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, errorhandler.NewInternalError("error creating user")
	}

	return userDomain, nil
}
