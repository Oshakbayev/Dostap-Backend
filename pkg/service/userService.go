package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
	"time"
)

type UserServiceInterface interface {
	SignUp(*entity.User) (int, error)
	LogIn(*entity.Credentials) (*entity.User, error)
	TokenGenerator(int64, string, string) (string, error)
	VerifyAccount(string) (int, error)
	TokenChecker(string) (*entity.Claims, int, error)
	DeleteAccount(int64) error
}

func (s *Service) SignUp(user *entity.User) (int, error) {
	//Encrypting passsword
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.EncryptedPass), bcrypt.DefaultCost)
	if err != nil {
		s.log.Printf("Error while hashing password at the service level: user - %v, error - %v", user, err)
		return http.StatusInternalServerError, fmt.Errorf("error while hashing password: %v, error: %s", user, err)
	}
	user.EncryptedPass = string(hashedPass)
	//if user has already registered but not verified
	err, userInDB := s.CheckUserExistence(user)
	if err != nil {
		s.log.Printf("Error in SignUp %s", err.Error())
		return http.StatusInternalServerError, err
	} else if userInDB != nil && !userInDB.IsEmailVerified {
		//update and resend
		user.ID = userInDB.ID
		if err = s.repo.UpdateUserByID(user); err != nil {
			return http.StatusInternalServerError, err
		}
	} else if userInDB != nil && userInDB.IsEmailVerified {
		return http.StatusBadRequest, fmt.Errorf("498 user with this email already exists")
	} else if userInDB == nil {
		//user registration
		err = s.repo.CreateUser(user)
		if err != nil {
			//s.log.Printf("Error while Inserting user into the table at the service level")
			return http.StatusInternalServerError, fmt.Errorf("error while CreateUser new user: %v, error: %s", user, err)
		}
	}
	emailContent, verificationLink, err := s.VerificationEmailGenerator(user.Email)
	if err != nil {
		s.log.Printf("Error in SignUp %s", err.Error())
		return http.StatusInternalServerError, err
	}
	err = s.SendVerificationEmail(user.Email, emailContent)
	if err != nil {
		s.log.Printf("Error in SignUp %s", err.Error())
		return http.StatusInternalServerError, err
	}
	if err = s.CreateVerifyEmail(user.ID, user.Email, verificationLink); err != nil {
		s.log.Printf("Error in SignUp %s", err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *Service) LogIn(credentials *entity.Credentials) (*entity.User, error) {
	user := &entity.User{}
	var err error
	if credentials.Username != "" {
		user, err = s.repo.GetUserByUsername(credentials.Username)
	} else if credentials.Email != "" {
		user, err = s.repo.GetUserByEmail(credentials.Email)
	} else {
		return nil, fmt.Errorf("492")
	}
	if err != nil {
		return nil, err
	}
	if !user.IsEmailVerified {
		return nil, fmt.Errorf(" 496 email %s is not verified", user.Email)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPass), []byte(credentials.Password)); err != nil {
		s.log.Printf("given password  is incorrect: %s", credentials.Password)
		return nil, fmt.Errorf(" 495 given password is incorrect: %s", credentials.Password)
	}
	return user, nil
}

func (s *Service) TokenGenerator(userID int64, email, username string) (string, error) {
	expTime := time.Now().Add(time.Hour * 48)
	claims := &entity.Claims{
		Email:    email,
		Level:    "user",
		Sub:      userID,
		Username: username,
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

func (s *Service) TokenChecker(tokenStr string) (*entity.Claims, int, error) {
	claims := &entity.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {

		return entity.JWTKey, nil
	})
	if err != nil {
		if err.Error() == jwt.ErrSignatureInvalid.Error() {
			s.log.Printf("Error in TokenChecker(Service): %v", err)
			return claims, http.StatusUnauthorized, err
		}
		s.log.Printf("Error in TokenChecker(Service): %v", err)
		return claims, http.StatusBadRequest, err
	}
	if !tkn.Valid {
		s.log.Printf("Error in TokenChecker(Service): %v", err)
		return claims, http.StatusBadRequest, err
	}
	decodedClaims := tkn.Claims.(*entity.Claims)
	return decodedClaims, http.StatusOK, nil
}

func (s *Service) VerifyAccount(secretCode string) (int, error) {
	verifyEmail, err := s.repo.GetVerifyEmailBySecretCode(secretCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.log.Printf("ErrNoRows in VerifyAccount %s", err.Error())
			return http.StatusBadRequest, errors.New("emergency: secretCode does not exist")
		}
		s.log.Printf("Error in VerifyAccount %s", err.Error())
		return http.StatusInternalServerError, err
	}
	//link expired
	if verifyEmail.ExpTime.Before(time.Now()) {
		s.log.Printf("Error in VerifyAccount %s", err.Error())
		return http.StatusBadRequest, errors.New("link expired")
	}

	err = s.repo.UpdateUserEmailStatus(verifyEmail.Email, true)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.log.Printf("Error in VerifyAccount %s", err.Error())
			return http.StatusBadRequest, errors.New("user does not exist")
		}
		s.log.Printf("Error in VerifyAccount %s", err.Error())
		return http.StatusInternalServerError, err
	}
	err = s.repo.UpdateVerifyEmail(secretCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.log.Printf("Error in VerifyAccount %s", err.Error())
			return http.StatusBadRequest, errors.New("email does not exist")
		}
		s.log.Printf("Error in VerifyAccount %s", err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *Service) CheckUserExistence(user *entity.User) (error, *entity.User) {
	sameUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			s.log.Printf("ErrNoRows in CheckUserExistence %s", err.Error())
			return nil, nil
		}
		s.log.Printf("Error in CheckUserExistence %s", err.Error())
		return err, nil
	}
	return nil, sameUser
}

func (s *Service) DeleteAccount(userID int64) error {
	err := s.repo.DeleteUserByID(userID)
	if err != nil {
		s.log.Printf("ErrNoRows in DeleteAccount %s", err.Error())
		return err
	}
	return nil
}
