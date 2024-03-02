package repo

import (
	"database/sql"
	"log"
)

type RepInterface interface {
	UserInterface
}

type Repository struct {
	log *log.Logger
	db  *sql.DB
}

func CreateRepository(db *sql.DB, l *log.Logger) RepInterface {
	return Repository{db: db, log: l}
}
