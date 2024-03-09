package repo

import (
	"database/sql"
	"hellowWorldDeploy/pkg/entity"
	"time"
)

type EmailInterface interface {
	GetVerifyEmailBySecretCode(string) (int64, time.Time, error)
	CreateEmail(*entity.Email) error
	UpdateVerifyEmail(string) error
}

func (r *Repository) CreateEmail(email *entity.Email) error {
	stmt, err := r.db.Prepare(`INSERT INTO verify_emails (user_id, email, secret_code, expired_at) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		r.log.Printf("Error at the stage of preparing verify_email data to Insert CreateEmail(repo):%s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("Error at the stage of closing stmt of email insert CreateEmail(repo): %s\n", err.Error())
		}
	}(stmt)
	if _, err := stmt.Exec(email.UserID, email.Email, email.SecretCode, email.ExpTime); err != nil {
		r.log.Printf("Error at the stage of verify_email data Inserting CreateEmail(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetVerifyEmailBySecretCode(secretCode string) (int64, time.Time, error) {
	var userID int64
	var expTime time.Time
	stmt, err := r.db.Prepare("SELECT user_id,expired_at FROM verify_emails WHERE secret_code = $1 ")
	if err != nil {
		r.log.Printf("Error while to prepare data to select verify_email by secretCode GetVerifyEmail(repo): %s\n", err.Error())
		return -1, expTime, err
	}
	err = stmt.QueryRow(secretCode).Scan(&userID, &expTime)
	if err != nil {
		r.log.Printf("Error while selecting verify_email data GetVerifyEmail(repo): %s\n", err.Error())
		return -1, expTime, err
	}
	return userID, expTime, nil
}

func (r *Repository) UpdateVerifyEmail(secretCode string) error {
	stmt, err := r.db.Prepare("UPDATE verify_emails SET is_used = true WHERE secret_code = $1")
	if err != nil {
		r.log.Printf("Error while to prepare data to update verify_email by secretCode UpdateEmail(repo): %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("Error at the stage of closing stmt of verify_email table  UpdateEmail(repo): %s\n", err.Error())
		}
	}(stmt)
	_, err = stmt.Query(secretCode)
	if err != nil {
		r.log.Printf("Error while exec update email by secretCode UpdateEmail(repo): %s\n", err.Error())
		return err
	}
	return nil
}
