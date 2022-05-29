package authentication

import "github.com/resyahrial/go-commerce/pkg/gvalidator"

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"-" validate:"required"`
}

func (l Login) Validate() (string, bool) {
	return gvalidator.Validate(l)
}

type Token struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}
