package user

import "github.com/segmentio/ksuid"

type User struct {
	ID       ksuid.KSUID `json:"ID"`
	Email    string      `json:"email"`
	Name     string      `json:"name"`
	Password string      `json:"-"`
	Address  string      `json:"address"`
}
