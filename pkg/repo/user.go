package repo

import (
	"database/sql"
	"fmt"
	"hellowWorldDeploy/pkg/entity"
)

type UserInterface interface {
	CreateUser(*entity.User) error
	GetUserByID(int64) (*entity.User, error)
	GetUserByPhoneNum(string) (*entity.User, error)
}

func (r *Repository) CreateUser(user *entity.User) error {
	stmt, err := r.db.Prepare(`INSERT INTO users (first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	if err != nil {
		r.log.Printf("Error at the stage of preparing data to Insert:%s\n", err.Error())
		fmt.Println("kek1", err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("Error at the stage of closing stmt: %s\n", err.Error())
		}
	}(stmt)
	_, err = stmt.Exec(user.FirstName, user.Lastname, user.EncryptedPass, user.AvatarLink, user.Gender, user.Age, user.PhoneNum, user.ResidenceCity, user.Description)
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

func (r Repository) GetUserByID(ID int64) (*entity.User, error) {
	stmt, err := r.db.Prepare("SELECT id,first_name,last_name,encrypted_password,avatar_link,gender,age,phone_number,city_of_residence,description FROM users WHERE id = ?")
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
	err = stmt.QueryRow(ID).Scan(&user.ID, &user.FirstName, &user.Lastname, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description)
	if err != nil {
		r.log.Printf("Error while to query row and scan user to get by id: %s\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (r Repository) GetUserByPhoneNum(phoneNum string) (*entity.User, error) {
	stmt, err := r.db.Prepare("SELECT id,first_name,last_name,encrypted_password,avatar_link,gender,age,phone_number,city_of_residence,description FROM users WHERE phone_number = ?")
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
	err = stmt.QueryRow(phoneNum).Scan(&user.ID, &user.FirstName, &user.Lastname, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description)
	if err != nil {
		r.log.Printf("Error while to query row and scan user to get by id: %s\n", err.Error())
		return nil, err
	}
	return user, nil
}
