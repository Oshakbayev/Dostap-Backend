package repo

import (
	"database/sql"
	"hellowWorldDeploy/pkg/entity"
)

type EventInterface interface {
	CreateEvent(event *entity.Event) error
}

func (r *Repository) CreateEvent(event *entity.Event) error {
	stmt, err := r.db.Prepare(`INSERT INTO events (organizer_id, event_name, format_id, address, coordinatex, coordinatey, capacity, link, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`)
	if err != nil {
		r.log.Printf("\nError at the stage of preparing data to Insert CreateEvent(repo):%s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError at the stage of closing stmt CreateEvent(repo): %s\n", err.Error())
		}
	}(stmt)
	err = stmt.QueryRow(event.OrganizerID, event.EventName, event.FormatID, event.Address, event.CoordinateX, event.CoordinateY, event.Capacity, event.Link, event.Description).Scan(&event.ID)
	if err != nil {
		r.log.Printf("\nError at the stage of data Inserting CreateEvent(repo): %s\n", err.Error())
		return err
	}
	return nil
}
