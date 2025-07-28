package initializers

import (
	"go-jwt/models/user/entities"
)

func SyncDatabase() {
	DB.AutoMigrate(&entities.UserModel{})
}
