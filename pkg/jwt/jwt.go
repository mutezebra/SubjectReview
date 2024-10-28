package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"

	"github.com/mutezebra/subject-review/config"
	"github.com/mutezebra/subject-review/pkg/constants"
)

var (
	expireTime = constants.JwtExpireTime
	jwtSecret  []byte
)

type Claims struct {
	UID   int64
	Email string
	jwt.StandardClaims
}

func GenerateToken(uid int64, email string) (string, error) {
	if len(jwtSecret) == 0 {
		jwtSecret = []byte(config.Secret.JwtSecret)
	}

	expire := time.Now().Add(expireTime).Unix()
	claim := &Claims{
		UID:   uid,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer:    config.Secret.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	s, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.WithMessage(err, "failed signed token")
	}

	return s, nil
}

func CheckToken(token string) (int64, string, error) {
	if token[0] == 'B' {
		bearer := token[0:6]
		if bearer == "Bearer" {
			token = token[7:]
		}
	}

	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return 0, "", errors.WithMessage(err, "failed parse token")
	}
	claim, ok := tokenClaim.Claims.(*Claims)
	if !ok {
		return 0, "", errors.WithMessage(fmt.Errorf("jurge claim failed"), "")
	}

	return claim.UID, claim.Email, claim.Valid()
}
