package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"hellowWorldDeploy/pkg/entity"
	"math/rand"
	"net/smtp"
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
	fmt.Println("USEEERIDDD", user.ID)
	if err != nil {
		//s.log.Printf("Error while Inserting user into the table at the service level")
		return fmt.Errorf("error while insert new user: %v, error: %s", user, err)
	}
	verificationLink, err := s.VerificationLinkGenerator(user.Email)
	if err != nil {
		return err
	}
	err = s.VerificationEmailGenerator(user.Email, verificationLink, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) LogIn(email, pass string) error {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("there is no user with this number %s", email)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPass), []byte(pass)); err != nil {
		s.log.Printf("given password of %s is incorrect: %s", email, pass)
		return fmt.Errorf("given password of %s is incorrect: %s", email, pass)
	}
	return nil
}

func (s *Service) TokenGenerator(email string) (string, error) {
	expTime := time.Now().Add(time.Minute * 100)
	claims := &entity.Claims{
		Email: email,
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

func (s *Service) VerificationLinkGenerator(email string) (string, error) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	allowedChars := "0123456789"
	var result string
	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(len(allowedChars))
		result += string(allowedChars[randomIndex])
	}
	secretCode, err := bcrypt.GenerateFromPassword([]byte(result+email), bcrypt.DefaultCost)
	if err != nil {
		s.log.Printf("error while hashing secretCode for email verification: error error - #{err}")
		return entity.EmtpyString, fmt.Errorf("error while hashing secretCode for email verification: error: %s", err)
	}
	return string(secretCode), nil
}

func (s *Service) VerificationEmailGenerator(emailAddress, verificationLink string, userID int64) error {
	expTime := time.Now().Add(time.Hour * 48)
	to := []string{emailAddress}
	subject := "Subject: DostApp registration link\r\n"
	emailContent := "Hello, thanks for registration on our app! Please follow the link attached below to complete your registration " + "http://localhost:8080/auth/confirmUserAccount?link=" + verificationLink + " Reminder: the link is valid for 2 days"
	auth := smtp.PlainAuth("", "rakatan228322@gmail.com", "zgjw nlyp zyhk bczp", "smtp.gmail.com")
	msg := []byte(subject + "\r\n" + emailContent)
	err := smtp.SendMail("smtp.gmail.com:587", auth, "rakatan228322@gmail.com", to, msg)
	if err != nil {
		s.log.Printf("error while sending verification email: error: #{err}")
		return fmt.Errorf("error while sending verification email: error: %s", err)
	}
	email := entity.Email{
		UserID:     userID,
		Email:      emailContent,
		SecretCode: verificationLink,
		ExpTime:    expTime,
	}
	err = s.repo.CreateEmail(&email)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) VerifyAccount(secretCode string) error {
	userId, expTime, err := s.repo.VerifyEmail(secretCode)
	if err != nil {
		return err
	}
	//link is invalid
	if expTime.Before(time.Now()) {
		return errors.New("Link is invalid")
	}
	fmt.Println("HERERERERER ", userId)
	user, err := s.repo.GetUserByID(userId)
	if err != nil {
		return err
	}

	err = s.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	err = s.repo.UpdateEmail(secretCode)
	if err != nil {
		return err
	}
	return nil
}
