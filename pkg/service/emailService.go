package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"hellowWorldDeploy/pkg/entity"
	"net/smtp"
	"time"
)

func (s *Service) VerificationEmailGenerator(email string) (string, string, error) {
	result := s.generateRandomKey(entity.VerificationSecretCodeLength)
	secretCode, err := bcrypt.GenerateFromPassword([]byte(result+email), bcrypt.DefaultCost)
	if err != nil {
		s.log.Printf("error while hashing secretCode for email verification: error error - #{err}")
		return entity.EmtpyString, entity.EmtpyString, fmt.Errorf("error while hashing secretCode for email verification: error: %s", err)
	}
	emailContent := "Hello, thanks for registration on our app! Please follow the link attached below to complete your registration " + entity.VerificationLinkURL + string(secretCode) + " Reminder: the link is valid for 2 days"
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

func (s *Service) CreateVerifyEmail(userID int, emailContent, verificationLink string) error {
	expTime := time.Now().Add(time.Hour * 48)
	email := entity.Email{
		UserID:     userID,
		Email:      emailContent,
		SecretCode: verificationLink,
		ExpTime:    expTime,
	}
	err := s.repo.CreateEmail(&email)
	if err != nil {
		s.log.Printf("Error in CreateVerifyEmail %s", err.Error())
		return err
	}
	return nil
}
