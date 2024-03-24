package repo

import (
	"database/sql"
	"fmt"
	"hellowWorldDeploy/pkg/entity"
)

type UserInterface interface {
	CreateUser(*entity.User) error
	GetUserByID(int64) (*entity.User, error)
	GetUserByEmail(string) (*entity.User, error)
	UpdateUserByID(*entity.User) error
	UpdateUserEmailStatus(string, bool) error
	DeleteUserByID(int64) error
}

func (r *Repository) CreateUser(user *entity.User) error {
	stmt, err := r.db.Prepare(`INSERT INTO users (first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description,email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10) RETURNING id`)
	if err != nil {
		r.log.Printf("\nError at the stage of preparing data to Insert CreateUser(repo):%s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError at the stage of closing stmt CreateUser(repo): %s\n", err.Error())
		}
	}(stmt)
	err = stmt.QueryRow(user.FirstName, user.LastName, user.EncryptedPass, user.AvatarLink, user.Gender, user.Age, user.PhoneNum, user.ResidenceCity, user.Description, user.Email).Scan(&user.ID)
	if err != nil {
		r.log.Printf("\nError at the stage of data Inserting CreateUser(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetUserByID(ID int64) (*entity.User, error) {
	stmt, err := r.db.Prepare("SELECT id,first_name,last_name,password,avatar_link,gender,age,phone_number,city_of_residence,description,email,is_email_verified FROM users WHERE id = $1")
	if err != nil {
		r.log.Printf("\nError while to prepare data to GetUserByID(repo) from user table: %s\n", err.Error())
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError at the stage of closing stmt GetUserByID(repo): %s\n", err.Error())
		}
	}(stmt)
	user := &entity.User{}
	err = stmt.QueryRow(ID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description, &user.Email, &user.IsEmailVerified)
	if err != nil {
		r.log.Printf("\nError while to query row and scan user to GetUserByID(repo): %s\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (*entity.User, error) {
	stmt, err := r.db.Prepare("SELECT id,first_name,last_name,password,avatar_link,gender,age,phone_number,city_of_residence,description,email,is_email_verified FROM users WHERE email = $1")
	if err != nil {
		r.log.Printf("\nError while to prepare data to GetUserByEmail(repo) from user table: %s\n", err.Error())
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError at the stage of closing stmt in GetUserByEmail(repo): %s\n", err.Error())
		}
	}(stmt)
	user := &entity.User{}
	err = stmt.QueryRow(email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description, &user.Email, &user.IsEmailVerified)
	if err != nil {
		r.log.Printf("\nError while to query row and scan GetUserByEmail(repo): %s\n", err.Error())
		return nil, err
	}
	//fmt.Println(user)
	return user, nil
}

func (r *Repository) UpdateUserByID(user *entity.User) error {
	stmt, err := r.db.Prepare(`
    UPDATE users
    SET 
        first_name = $1,
        last_name = $2,
        avatar_link = $3,
        gender = $4,
        age = $5,
        phone_number = $6,
        city_of_residence = $7,
        description = $8,
        is_email_verified = $9
    WHERE id = $10
`)
	if err != nil {
		r.log.Printf("\nError while preparing data to UpdateUserByID(repo) by id: %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError at the stage of closing stmt in UpdateUserByID(repo): %s\n", err.Error())
		}
	}(stmt)

	_, err = stmt.Exec(
		user.FirstName,
		user.LastName,
		user.AvatarLink,
		user.Gender,
		user.Age,
		user.PhoneNum,
		user.ResidenceCity,
		user.Description,
		user.IsEmailVerified,
		user.ID, // Assuming ID is the 12th parameter in your query
	)
	if err != nil {
		r.log.Printf("\nError while executing UpdateUserByID(repo) by id: %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) UpdateUserEmailStatus(email string, status bool) error {
	stmt, err := r.db.Prepare(`
    UPDATE users
    SET
        is_email_verified = $1
    WHERE email = $2
        `)
	if err != nil {
		r.log.Printf("\nError while preparing data to UpdateUserEmailStatus(repo) by id: %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError at the stage of closing stmt in UpdateUserEmailStatus(repo): %s\n", err.Error())
		}
	}(stmt)
	_, err = stmt.Exec(status, email)
	if err != nil {
		r.log.Printf("\nError while executing UpdateUserEmailStatus(repo) by id: %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) DeleteUserByID(userID int64) error {
	stmt, err := r.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		r.log.Printf("\nError while preparing data to DeleteUserByID(repo) by id: %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError at the stage of closing stmt in DeleteUserByID(repo): %s\n", err.Error())
		}
	}(stmt)

	result, err := stmt.Exec(userID)
	if err != nil {
		r.log.Printf("\nError while executing DeleteUserByID(repo) by id: %s\n", err.Error())
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.log.Printf("\nerror getting rows affected; DeleteUserByID(repo): %s\n", err.Error())
		return fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		r.log.Printf("\nno user found with ID; DeleteUserByID(repo): %d\n", userID)
		// No rows were affected, which means no user with the given ID exists
		return fmt.Errorf("no user found with ID %d", userID)
	}
	return nil
}
