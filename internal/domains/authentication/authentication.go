package authentication

type Login struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Token struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}
