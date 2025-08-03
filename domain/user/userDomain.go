package domain

import (
	"github.com/golang-jwt/jwt/v5"
	errorhandler "go-jwt/errorHandler"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type userDomain struct {
	email    string
	password string
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string

	EncryptPassword()
	ComparePassword(string) *errorhandler.ErrorHandler
	GenerateToken() (string, *errorhandler.ErrorHandler)
}

func NewUserDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func (ud *userDomain) EncryptPassword() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(ud.password), 10)
	ud.password = string(hash)
}

func (ud *userDomain) ComparePassword(password string) *errorhandler.ErrorHandler {
	err := bcrypt.CompareHashAndPassword([]byte(ud.GetPassword()), []byte(password))
	if err != nil {
		return errorhandler.NewBadRequestError("invalid password")
	}
	return nil
}

func (ud *userDomain) GenerateToken() (string, *errorhandler.ErrorHandler) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": ud.GetEmail(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, tokenErr := token.SignedString([]byte(os.Getenv("SECRET")))
	if tokenErr != nil {
		return "", errorhandler.NewBadRequestError("failed to create token")
	}

	return tokenString, nil
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}
