package models

import "github.com/segmentio/ksuid"

type User struct {
	ID       ksuid.KSUID `json:"ID" gorm:"primaryKey;type:varchar(30)" `
	Email    string      `json:"email"`
	Name     string      `json:"name"`
	Password string      `json:"-"`
	Address  string      `json:"address"`
	Role     string      `json:"-"`
}
