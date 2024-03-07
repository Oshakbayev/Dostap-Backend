package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"hellowWorldDeploy/pkg/entity"
	"time"
)

func (s *Service) SignUp(user *entity.User) error {
	//Encrypting passsword
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.EncryptedPass), bcrypt.DefaultCost)
	if err != nil {
		s.log.Printf("error while hashing password at the service level:user - #{user}, error - #{err}")
		return fmt.Errorf("error while hashing password: %v, error: %s", user, err)
	}
	user.EncryptedPass = string(hashedPass)
	err = s.repo.CreateUser(user)
	if err != nil {
		//s.log.Printf("Error while Inserting user into the table at the service level")
		return fmt.Errorf("error while insert new user: %v, error: %s", user, err)
	}
	return nil
}

func (s *Service) LogIn(phoneNum, pass string) (string, error) {
	user, err := s.repo.GetUserByPhoneNum(phoneNum)
	if err != nil {
		return entity.EmtpyString, fmt.Errorf("there is no user with this number %s", phoneNum)
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPass), []byte(pass)); err != nil {
		fmt.Println(hashedPass, "----", user.EncryptedPass)
		s.log.Printf("given password of %s is incorrect: %s", phoneNum, pass)
		return entity.EmtpyString, fmt.Errorf("given password of %s is incorrect: %s", phoneNum, pass)
	}
	expTime := time.Now().Add(time.Minute * 100)
	claims := &entity.Claims{
		PhoneNum: phoneNum,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(entity.JWTKey)
	if err != nil {
		s.log.Printf("Error while signing jwt token")
		return entity.EmtpyString, fmt.Errorf("error while signing jwt token")
	}
	return signedToken, nil
}
