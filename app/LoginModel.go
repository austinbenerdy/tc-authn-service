package main

import "golang.org/x/crypto/bcrypt"

type LoginModel struct {
	Email    string
	Password string
}

func newLoginModel(email string, password string) LoginModel {
	return LoginModel{
		email,
		password,
	}
}

func (lm *LoginModel) HashPassword() string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(lm.Password), 14)
	return string(bytes)
}
