package entity

import "github.com/golang-jwt/jwt"

var JWTKey = []byte("sercet_key")

type User struct {
	ID            int64
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	EncryptedPass string `json:"password"`
	AvatarLink    string `json:"avatar_link"`
	Gender        string `json:"gender"`
	Age           int64  `json:"age"`
	PhoneNum      string `json:"phone_number"`
	ResidenceCity string `json:"city_of_residence"`
	Description   string `json:"description"`
}

type Credentials struct {
	PhoneNum string `json:"phone_number"`
	Password string `json:"password"`
}

type Claims struct {
	PhoneNum string `json:"phone_number"`
	jwt.StandardClaims
}

type event struct{}
