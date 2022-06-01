package tokenmanager

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/resyahrial/go-commerce/pkg/inspect"
	"github.com/resyahrial/go-commerce/pkg/transformers"
	log "github.com/sirupsen/logrus"
)

type JwtTokenManager struct {
	keyAccess        []byte
	keyRefresh       []byte
	expiryAgeAccess  time.Duration
	expiryAgeRefresh time.Duration
}

type JwtTokenManagerOpts struct {
	KeyAccess        string
	KeyRefresh       string
	ExpiryAgeAccess  time.Duration
	ExpiryAgeRefresh time.Duration
}

func NewJwtTokenManager(opts JwtTokenManagerOpts) TokenManager {
	return &JwtTokenManager{
		keyAccess:        []byte(opts.KeyAccess),
		keyRefresh:       []byte(opts.KeyRefresh),
		expiryAgeAccess:  opts.ExpiryAgeAccess,
		expiryAgeRefresh: opts.ExpiryAgeRefresh}
}

func (t *JwtTokenManager) GenerateAccess(claims Claims) (string, bool) {
	return t.generate(claims, t.keyAccess, t.expiryAgeAccess)
}

func (t *JwtTokenManager) GenerateRefresh(claims Claims) (string, bool) {
	return t.generate(claims, t.keyRefresh, t.expiryAgeRefresh)
}

func (t JwtTokenManager) ParseAccess(tokenString string) (claims Claims, err error) {
	return
}

func (t JwtTokenManager) ParseRefresh(tokenString string) (claims Claims, err error) {
	return
}

func (t *JwtTokenManager) generate(claims Claims, key []byte, expiryAge time.Duration) (string, bool) {
	var tokenStr string
	var claimsMap map[string]interface{}
	var err error
	defer func(err error) {
		if err != nil {
			log.Error(err)
		}
	}(err)

	jwtMapClaims := jwt.MapClaims{
		"exp": time.Now().Add(expiryAge).Unix(),
	}

	if claimsMap, err = transformers.StructToMap(claims); err != nil {
		return "", false
	}

	for k, v := range claimsMap {
		jwtMapClaims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtMapClaims)
	if tokenStr, err = token.SignedString(key); err != nil {
		inspect.Do(err)
		return "", false
	}

	return tokenStr, true
}

/*
	- fix error statment
	- cast claims
*/
func Parse(key string, tokenString string) (claims Claims, err error) {
	var token *jwt.Token
	if token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return key, nil
	}); err != nil {
		return
	}

	if token.Valid {
		err = fmt.Errorf("you look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			err = fmt.Errorf("that's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			err = fmt.Errorf("timing is everything")
		} else {
			err = fmt.Errorf("couldn't handle this token: %v", err)
		}
	} else {
		err = fmt.Errorf("couldn't handle this token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	return
}
