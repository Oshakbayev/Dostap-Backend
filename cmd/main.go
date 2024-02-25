package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var db *sql.DB

type User struct {
	ID   int
	Name string
	Role string
}

func initDB() {
	var err error
	connectionString := "postgres://owwkdlwj:UqnYqPkMlDPUHBntLlFYpIeLjaXZkCxR@abul.db.elephantsql.com/owwkdlwj"
	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		fmt.Println("kek")
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("kek2")
		log.Fatal(err)
	}
}

func main() {
	initDB()
	defer db.Close()
	fmt.Println("Ansaer Loh and Pidaraz and Mal")
	port := "80"
	router := http.NewServeMux()
	router.HandleFunc("/", MainPage)
	fmt.Println("kek3")

	if err := http.ListenAndServe(":"+port, router); err != nil {

		log.Fatal("Error starting server: ", err)
	}
}

func QueryUsers(db *sql.DB) ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT id, name, role FROM users")
	if err != nil {
		return nil, err // return an error immediately
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Role)
		if err != nil {
			return nil, err // return an error immediately
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil // return the slice of users
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	users, err := QueryUsers(db)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println("QueryUsers error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON encoding error:", err)
	}
}

//func main() {
//	server := pkg.InitServer()
//	server.InitRouter()
//}
