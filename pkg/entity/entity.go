package entity

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	JWTKey                    = []byte("sercet_key")
	EmtpyString               = ""
	VerificationLinkURL       = "http://92.38.48.85:80/auth/confirmUserAccount?link="
	NilID               int64 = -1
)

type User struct {
	ID              int64
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	EncryptedPass   string `json:"password"`
	AvatarLink      string `json:"avatar_link"`
	Gender          string `json:"gender"`
	Age             int64  `json:"age"`
	PhoneNum        string `json:"phone_number"`
	ResidenceCity   string `json:"city_of_residence"`
	Description     string `json:"description"`
	Email           string `json:"email"`
	IsEmailVerified bool
}

type Email struct {
	ID         int64
	UserID     int64
	Email      string
	SecretCode string
	IsUsed     bool
	ExpTime    time.Time
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	Sub   int64  `json:"sub"`
	Level string `json:"level"`
	jwt.StandardClaims
}

// struct for jsopn decoding for update user info
type UpdateJson struct {
	Token    TokenData `json:"jwtToken"`
	UserInfo User      `json:"userInfo"`
}

type TokenData struct {
	Token string
}

type event struct{}

type ResponseJSON struct {
	Message string
}
