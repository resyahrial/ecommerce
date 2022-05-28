package user

import (
	"github.com/mitchellh/mapstructure"
	"github.com/segmentio/ksuid"
)

const (
	BUYER  = "BUYER"
	SELLER = "SELLER"
)

type User struct {
	ID       ksuid.KSUID `json:"ID"`
	Email    string      `json:"email"`
	Name     string      `json:"name"`
	Password string      `json:"-"`
	Address  string      `json:"address"`
	Role     string      `json:"-"`
}

func (u User) ToBuyer() (buyer Buyer, ok bool) {
	if u.Role != BUYER {
		return buyer, false
	}

	if err := mapstructure.Decode(u, &buyer); err != nil {
		return buyer, false
	}

	return
}

func (u User) ToSeller() (seller Seller, ok bool) {
	if u.Role != SELLER {
		return seller, false
	}

	if err := mapstructure.Decode(u, &seller); err != nil {
		return seller, false
	}

	return
}

type Buyer struct {
	ID      ksuid.KSUID `json:"ID"`
	Email   string      `json:"email"`
	Name    string      `json:"name"`
	Address string      `json:"address"`
}

type Seller struct {
	ID      ksuid.KSUID `json:"ID"`
	Email   string      `json:"email"`
	Name    string      `json:"name"`
	Address string      `json:"address"`
}
