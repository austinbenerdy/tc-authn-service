package models

import (
	"time"
)

type UserToken struct {
	Id         string
	UserId     string
	Token      string
	Expiration time.Time
	Expired    bool
}

func (u *UserToken) IsExpire() bool {
	return time.Now().After(u.Expiration) || u.Expired
}
