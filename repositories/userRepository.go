package repositories

import (
	"errors"
	"go-jwt/initializers"
	"go-jwt/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return UserRepository{
		db: initializers.DB,
	}
}

func (u *UserRepository) Save(userModel *models.UserModel) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(userModel).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *UserRepository) FindByEmail(email string) (models.UserModel, error) {
	var user models.UserModel
	u.db.First(&user, "email = ?", email)
	if user.ID == 0 {
		return models.UserModel{}, errors.New("user not found")
	}

	return user, nil
}
