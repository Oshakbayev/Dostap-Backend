package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"hellowWorldDeploy/pkg/entity"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"
)

type UserServiceInterface interface {
	SignUp(*entity.User) (int, error)
	LogIn(string, string) (int, error)
	TokenGenerator(string) (string, error)
	VerifyAccount(string) (int, error)
}

func (s *Service) SignUp(user *entity.User) (int, error) {
	//Encrypting passsword
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.EncryptedPass), bcrypt.DefaultCost)
	if err != nil {
		s.log.Printf("Error while hashing password at the service level: user - %v, error - %v", user, err)
		return http.StatusInternalServerError, fmt.Errorf("error while hashing password: %v, error: %s", user, err)
	}
	user.EncryptedPass = string(hashedPass)
	isExist, err, userInDB := s.CheckUserExistence(user)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if isExist && !userInDB.IsEmailVerified {
		//update and resend
		if err = s.repo.UpdateUser(user); err != nil {
			return http.StatusInternalServerError, err
		}
		emailContent, verificationLink, err := s.VerificationEmailGenerator(user.Email)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		err = s.SendVerificationEmail(user.Email, emailContent)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		if err = s.InsertVerificationEmail(user.ID, emailContent, verificationLink); err != nil {
			return http.StatusInternalServerError, err
		}
	} else if isExist && userInDB.IsEmailVerified {
		return http.StatusBadRequest, errors.New("user with this email already exists")
	}
	err = s.repo.CreateUser(user)
	if err != nil {
		//s.log.Printf("Error while Inserting user into the table at the service level")
		return http.StatusInternalServerError, fmt.Errorf("error while insert new user: %v, error: %s", user, err)
	}
	emailContent, verificationLink, err := s.VerificationEmailGenerator(user.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	err = s.SendVerificationEmail(user.Email, emailContent)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if err = s.InsertVerificationEmail(user.ID, emailContent, verificationLink); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *Service) LogIn(email, pass string) (int, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, fmt.Errorf("given email %s is incorrect", email)
		}
		return http.StatusInternalServerError, err
	}
	if !user.IsEmailVerified {
		return http.StatusBadRequest, fmt.Errorf("email %s did not verified", email)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPass), []byte(pass)); err != nil {
		s.log.Printf("given password of %s is incorrect: %s", email, pass)
		return http.StatusBadRequest, fmt.Errorf("given password of %s is incorrect: %s", email, pass)
	}
	return http.StatusOK, nil
}

func (s *Service) TokenGenerator(email string) (string, error) {
	expTime := time.Now().Add(time.Minute * 100)
	claims := &entity.Claims{
		Email: email,
		Level: "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(entity.JWTKey)
	if err != nil {
		s.log.Printf("Error while signing jwt token: %v", err)
		return entity.EmtpyString, err
	}
	return signedToken, nil
}

func (s *Service) VerificationEmailGenerator(email string) (string, string, error) {
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
		return entity.EmtpyString, entity.EmtpyString, fmt.Errorf("error while hashing secretCode for email verification: error: %s", err)
	}
	emailContent := "Hello, thanks for registration on our app! Please follow the link attached below to complete your registration " + "http://localhost:80/auth/confirmUserAccount?link=" + string(secretCode) + " Reminder: the link is valid for 2 days"
	return emailContent, string(secretCode), nil
}

func (s *Service) SendVerificationEmail(emailAddress, emailContent string) error {
	to := []string{emailAddress}
	subject := "Subject: DostApp registration link\r\n"
	auth := smtp.PlainAuth("", "rakatan228322@gmail.com", "zgjw nlyp zyhk bczp", "smtp.gmail.com")
	msg := []byte(subject + "\r\n" + emailContent)
	err := smtp.SendMail("smtp.gmail.com:587", auth, "rakatan228322@gmail.com", to, msg)
	if err != nil {
		s.log.Printf("error while sending verification email: error: #{err}")
		return fmt.Errorf("error while sending verification email: error: %s", err)
	}
	return nil
}

func (s *Service) VerifyAccount(secretCode string) (int, error) {
	userId, expTime, err := s.repo.GetVerifyEmailBySecretCode(secretCode)
	if err != nil {
		fmt.Println(err, "HERERERERERERE")
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, errors.New("emergency: secretCode does not exist")
		}
		return http.StatusInternalServerError, err
	}
	fmt.Println(userId, "----------")
	//link expired
	if expTime.Before(time.Now()) {
		fmt.Println(err, "HERERERERERERE2")
		return http.StatusBadRequest, errors.New("link expired")
	}

	user, err := s.repo.GetUserByID(userId)
	if err != nil {
		fmt.Println(err, "HERERERERERERE3")
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, errors.New("user does not exist")
		}
		return http.StatusInternalServerError, err
	}
	fmt.Println(user, "++++++")
	fmt.Println(user.Email, "------++++++")
	user.IsEmailVerified = true
	err = s.repo.UpdateUser(user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, errors.New("user does not exist")
		}
		return http.StatusInternalServerError, err
	}
	err = s.repo.UpdateVerifyEmail(secretCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, errors.New("email does not exist")
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *Service) CheckUserExistence(user *entity.User) (bool, error, *entity.User) {
	sameUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return false, err, nil
		}
		return false, nil, nil
	}
	return true, nil, sameUser
}

func (s *Service) InsertVerificationEmail(userID int64, emailContent, verificationLink string) error {
	expTime := time.Now().Add(time.Hour * 48)
	email := entity.Email{
		UserID:     userID,
		Email:      emailContent,
		SecretCode: verificationLink,
		ExpTime:    expTime,
	}
	err := s.repo.CreateEmail(&email)
	if err != nil {
		return err
	}
	return nil
}
