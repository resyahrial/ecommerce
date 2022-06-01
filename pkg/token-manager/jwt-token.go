package tokenmanager

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
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

func (t *JwtTokenManager) ParseAccess(tokenString string) (claims Claims, err error) {
	return t.parse(t.keyAccess, tokenString)
}

func (t *JwtTokenManager) ParseRefresh(tokenString string) (claims Claims, err error) {
	return t.parse(t.keyRefresh, tokenString)
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

func (t *JwtTokenManager) parse(key []byte, tokenString string) (claims Claims, err error) {
	var token *jwt.Token
	if token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	}); err != nil {
		return
	}

	if !token.Valid {
		err = fmt.Errorf("invalid token")
		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			err = fmt.Errorf("token is either expired or not active yet")
		} else {
			err = fmt.Errorf("couldn't handle this token: %v", err)
		}
	} else {
		err = fmt.Errorf("couldn't handle this token: %v", err)
	}

	if mapClaims, ok := token.Claims.(jwt.MapClaims); ok {
		if err = mapstructure.Decode(mapClaims, &claims); err != nil {
			return
		}
	}

	return
}
