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
	UpdateUser(*entity.User) error
}

func (r *Repository) CreateUser(user *entity.User) error {
	fmt.Println("initial userId", user.ID)
	stmt, err := r.db.Prepare(`INSERT INTO users (first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description,email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10) RETURNING id`)
	if err != nil {
		r.log.Printf("Error at the stage of preparing data to Insert:%s\n", err.Error())
		fmt.Println("kek1111", err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("Error at the stage of closing stmt: %s\n", err.Error())
		}
	}(stmt)
	err = stmt.QueryRow(user.FirstName, user.LastName, user.EncryptedPass, user.AvatarLink, user.Gender, user.Age, user.PhoneNum, user.ResidenceCity, user.Description, user.Email).Scan(&user.ID)
	if err != nil {
		r.log.Printf("Error at the stage of data Inserting: %s\n", err.Error())
		fmt.Println("kek2", err)
		return err
	}
	//user.ID, err = result.LastInsertId()
	//if err != nil {
	//	r.log.Printf("Error while exec prepared data to Insert to user Table ")
	//	fmt.Println("kek3", err)
	//	return err
	//}
	return nil
}

func (r *Repository) GetUserByID(ID int64) (*entity.User, error) {
	stmt, err := r.db.Prepare("SELECT id,first_name,last_name,password,avatar_link,gender,age,phone_number,city_of_residence,description FROM users WHERE id = $1")
	if err != nil {
		r.log.Printf("Error while to prepare data to get user by id from user table: %s\n", err.Error())
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("Error at the stage of closing stmt: %s\n", err.Error())
		}
	}(stmt)
	user := &entity.User{}
	err = stmt.QueryRow(ID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description)
	if err != nil {
		r.log.Printf("Error while to query row and scan user to get by id: %s\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (*entity.User, error) {
	stmt, err := r.db.Prepare("SELECT id,first_name,last_name,password,avatar_link,gender,age,phone_number,city_of_residence,description FROM users WHERE email = $1")
	if err != nil {
		r.log.Printf("Error while to prepare data to get user by email from user table: %s\n", err.Error())
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("Error at the stage of closing stmt: %s\n", err.Error())
		}
	}(stmt)
	user := &entity.User{}
	err = stmt.QueryRow(email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description)
	if err != nil {
		r.log.Printf("Error while to query row and scan user to get by id: %s\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (r *Repository) UpdateUser(user *entity.User) error {
	stmt, err := r.db.Prepare(`
    UPDATE users
    SET 
        first_name = $1,
        last_name = $2,
        password = $3,
        avatar_link = $4,
        gender = $5,
        age = $6,
        phone_number = $7,
        city_of_residence = $8,
        description = $9,
        email = $10,
        is_email_verified = $11
    WHERE id = $12
`)
	if err != nil {
		r.log.Printf("Error while preparing data to update user by id: %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("Error at the stage of closing stmt: %s\n", err.Error())
		}
	}(stmt)

	_, err = stmt.Exec(
		user.FirstName,
		user.LastName,
		user.EncryptedPass,
		user.AvatarLink,
		user.Gender,
		user.Age,
		user.PhoneNum,
		user.ResidenceCity,
		user.Description,
		user.Email,
		user.IsEmailVerified,
		user.ID, // Assuming ID is the 12th parameter in your query
	)
	if err != nil {
		r.log.Printf("Error while executing update user by id: %s\n", err.Error())
		return err
	}
	return nil
}
