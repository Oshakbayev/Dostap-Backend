package repo

import (
	"database/sql"
	"hellowWorldDeploy/pkg/entity"
)

type EmailInterface interface {
	GetVerifyEmailBySecretCode(string) (*entity.Email, error)
	CreateEmail(*entity.Email) error
	UpdateVerifyEmail(string) error
}

func (r *Repository) CreateEmail(email *entity.Email) error {
	stmt, err := r.db.Prepare(`INSERT INTO verify_emails (user_id, email, secret_code, expired_at) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		r.log.Printf("\nError at the stage of preparing verify_email data to Insert CreateEmail(repo):%s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError at the stage of closing stmt of email insert CreateEmail(repo): %s\n", err.Error())
		}
	}(stmt)
	if _, err := stmt.Exec(email.UserID, email.Email, email.SecretCode, email.ExpTime); err != nil {
		r.log.Printf("\nError at the stage of verify_email data Inserting CreateEmail(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetVerifyEmailBySecretCode(secretCode string) (*entity.Email, error) {
	verifyEmail := entity.Email{}
	stmt, err := r.db.Prepare("SELECT * FROM verify_emails WHERE secret_code = $1 ")
	if err != nil {
		r.log.Printf("\nError while to prepare data to select verify_email by secretCode GetVerifyEmail(repo): %s\n", err.Error())
		return &verifyEmail, err
	}
	err = stmt.QueryRow(secretCode).Scan(&verifyEmail.ID, &verifyEmail.UserID, &verifyEmail.Email, &verifyEmail.SecretCode, &verifyEmail.IsUsed, &verifyEmail.ExpTime)
	if err != nil {
		r.log.Printf("\nError while selecting verify_email data GetVerifyEmail(repo): %s\n", err.Error())
		return &verifyEmail, err
	}
	return &verifyEmail, nil
}

func (r *Repository) UpdateVerifyEmail(secretCode string) error {
	stmt, err := r.db.Prepare("UPDATE verify_emails SET is_used = true WHERE secret_code = $1")
	if err != nil {
		r.log.Printf("\nError while to prepare data to update verify_email by secretCode UpdateEmail(repo): %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError at the stage of closing stmt of verify_email table  UpdateEmail(repo): %s\n", err.Error())
		}
	}(stmt)
	_, err = stmt.Query(secretCode)
	if err != nil {
		r.log.Printf("\nError while exec update email by secretCode UpdateEmail(repo): %s\n", err.Error())
		return err
	}
	return nil
}
