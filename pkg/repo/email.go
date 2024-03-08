package repo

import (
	"database/sql"
	"fmt"
	"hellowWorldDeploy/pkg/entity"
	"time"
)

type EmailInterface interface {
	VerifyEmail(string) (int64, time.Time, error)
	CreateEmail(*entity.Email) error
	UpdateEmail(string) error
}

func (r *Repository) CreateEmail(email *entity.Email) error {
	stmt, err := r.db.Prepare(`INSERT INTO verify_emails (user_id, email, secret_code, expired_at) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		r.log.Printf("Error at the stage of preparing data to Insert:%s\n", err.Error())
		fmt.Println("kekEmail1", err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("Error at the stage of closing stmt: %s\n", err.Error())
		}
	}(stmt)
	if _, err := stmt.Exec(email.UserID, email.Email, email.SecretCode, email.ExpTime); err != nil {
		r.log.Printf("Error at the stage of data Inserting to email: %s\n", err.Error())
		fmt.Println("emailKek2", err)
		return err
	}
	return nil
}

func (r *Repository) VerifyEmail(secretCode string) (int64, time.Time, error) {
	var userID int64
	var expTime time.Time
	stmt, err := r.db.Prepare("SELECT user_id,expired_at FROM verify_emails WHERE secret_code = $1 ")
	if err != nil {
		r.log.Printf("Error while to prepare data to get email by secretCode from verify_email table: %s\n", err.Error())
		return -1, expTime, err
	}
	err = stmt.QueryRow(secretCode).Scan(&userID, &expTime)
	if err != nil {
		r.log.Printf("Error while selecting email data: %s\n", err.Error())
		return -1, expTime, err
	}
	return userID, expTime, nil
	//stmt, err = r.db.Prepare("UPDATE verify_emails SET is_used = true WHERE secret_code = $1")
	//if err != nil {
	//	r.log.Printf("Error while to prepare data to get user by email from user table: %s\n", err.Error())
	//	//return nil, err
	//}
	//
	//defer func(stmt *sql.Stmt) {
	//	err := stmt.Close()
	//	if err != nil {
	//		r.log.Printf("Error at the stage of closing stmt: %s\n", err.Error())
	//	}
	//}(stmt)
}

func (r *Repository) UpdateEmail(secretCode string) error {
	stmt, err := r.db.Prepare("UPDATE verify_emails SET is_used = true WHERE secret_code = $1")
	if err != nil {
		r.log.Printf("Error while to prepare data to update email by secretCode from verify_emails table: %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("Error at the stage of closing stmt: %s\n", err.Error())
		}
	}(stmt)
	_, err = stmt.Query(secretCode)
	if err != nil {
		r.log.Printf("Error while exec update email by secretCode: %s\n", err.Error())
		return err
	}
	return nil
}
