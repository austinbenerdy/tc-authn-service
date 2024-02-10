package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       string
	Email    string
	Password string
}

func (u *User) Auth(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
