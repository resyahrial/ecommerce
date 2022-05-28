package authentication

import "github.com/resyahrial/go-commerce/pkg/gvalidator"

type Login struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"-"`
}

func (l Login) Validate() (string, bool) {
	return gvalidator.Validate(l)
}

type Token struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}
