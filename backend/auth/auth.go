package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECERETE_KEY = []byte("supersecretkey")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(email string, username string, role string) (tokenString string, err error) {

	// set time to expiry token
	expirationTime := time.Now().Add(5 * time.Hour)

	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Sign and get teh complete encoded token as a string using secrete key
	tokenString, err = token.SignedString(SECERETE_KEY)
	return
}
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECERETE_KEY), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
