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
	LogIn(string, string) (int, int64, error)
	TokenGenerator(int64, string) (string, error)
	VerifyAccount(string) (int, error)
	TokenChecker(string) (*entity.Claims, int, error)
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
	isExist, err, userInDB := s.CheckUserExistence(user)
	if err != nil {
		return http.StatusInternalServerError, err
	} else if isExist && !userInDB.IsEmailVerified {
		//update and resend
		if err = s.repo.UpdateUser(user); err != nil {
			return http.StatusInternalServerError, err
		}
	} else if isExist && userInDB.IsEmailVerified {
		return http.StatusBadRequest, errors.New("user with this email already exists")
	} else if !isExist {
		//user registration
		err = s.repo.CreateUser(user)
		if err != nil {
			//s.log.Printf("Error while Inserting user into the table at the service level")
			return http.StatusInternalServerError, fmt.Errorf("error while insert new user: %v, error: %s", user, err)
		}
	}
	emailContent, verificationLink, err := s.VerificationEmailGenerator(user.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	err = s.SendVerificationEmail(user.Email, emailContent)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if err = s.CreateVerifyEmail(userInDB.ID, user.Email, verificationLink); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *Service) LogIn(email, pass string) (int, int64, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, entity.NilID, fmt.Errorf("no user exist with given email %s", email)
		}
		return http.StatusInternalServerError, entity.NilID, err
	}
	if !user.IsEmailVerified {
		return http.StatusBadRequest, -entity.NilID, fmt.Errorf("email %s did not verified", email)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPass), []byte(pass)); err != nil {
		s.log.Printf("given password of %s is incorrect: %s", email, pass)
		return http.StatusBadRequest, entity.NilID, fmt.Errorf("given password of %s is incorrect: %s", email, pass)
	}
	return http.StatusOK, user.ID, nil
}

func (s *Service) TokenGenerator(userID int64, email string) (string, error) {
	expTime := time.Now().Add(time.Minute * 100)
	claims := &entity.Claims{
		Email: email,
		Level: "user",
		Sub:   userID,
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
			return http.StatusBadRequest, errors.New("emergency: secretCode does not exist")
		}
		return http.StatusInternalServerError, err
	}
	//link expired
	if verifyEmail.ExpTime.Before(time.Now()) {
		return http.StatusBadRequest, errors.New("link expired")
	}

	user, err := s.repo.GetUserByID(verifyEmail.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return http.StatusBadRequest, errors.New("user does not exist")
		}
		return http.StatusInternalServerError, err
	}
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
