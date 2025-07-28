package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type userDomain struct {
	email    string
	password string
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	EncryptPassword()
}

func NewUserDomain(email, password string) UserDomainInterface {
	return &userDomain{email, password}
}

func (ud *userDomain) EncryptPassword() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(ud.password), 10)
	ud.password = string(hash)
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}
