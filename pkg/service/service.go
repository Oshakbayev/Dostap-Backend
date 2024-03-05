package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"hellowWorldDeploy/pkg/entity"
)

func (s *Service) SignUp(user *entity.User) error {
	//Encrypting passsword
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.EncryptedPass), bcrypt.DefaultCost)
	if err != nil {
		s.log.Printf("error while hashing password at the service level:user - #{user}, error - #{err}")
		return fmt.Errorf("error while hashing password : user - #{user}, error - #{err}")
	}
	user.EncryptedPass = string(hashedPass)
	err = s.repo.CreateUser(user)
	if err != nil {
		s.log.Printf("Error while Inserting user into the table at the service level")
		return fmt.Errorf("error while insert new user: %v, error: %s", user, err)
	}
	return nil
}

func (s *Service) LogIn(phoneNum, pass string) error {
	user, err := s.repo.GetUserByPhoneNum(phoneNum)
	if err != nil {
		return fmt.Errorf("There is no user with this number: #{user.PhoneNum}")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.EncryptedPass)); err != nil {
		return fmt.Errorf("given pasword is incorrect: #{phoneNum}, {pass}")
	}
	return nil
}
