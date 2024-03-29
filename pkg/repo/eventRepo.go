package repo

import (
	"database/sql"
	"hellowWorldDeploy/pkg/entity"
	"strings"
)

type EventInterface interface {
	CreateEvent(event *entity.Event) error
	GetEventsByInterests([]string) ([]entity.Event, error)
	CreateEventInterests(int64, []string) error
	CreateEventOrganizers(eventID int64, organizers []string) error
	GetAllEvents() ([]entity.Event, error)
}

func (r *Repository) CreateEvent(event *entity.Event) error {
	stmt, err := r.db.Prepare(`INSERT INTO events (event_name, format_id, address, coordinatex, coordinatey, capacity, link, description,privacy_id,creator_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10) RETURNING id`)
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
	err = stmt.QueryRow(event.EventName, event.FormatID, event.Address, event.CoordinateX, event.CoordinateY, event.Capacity, event.Link, event.Description, event.PrivacyID, event.CreatorID).Scan(&event.ID)
	if err != nil {
		r.log.Printf("\nError at the stage of data Inserting CreateEvent(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) CreateEventOrganizers(eventID int64, organizers []string) error {
	organizersString := strings.Join(organizers, ",")
	_, err := r.db.Exec(`INSERT INTO event_organizers (event_id, organizer_id) VALUES ($1,unnest(string_to_array($2, ','))::int)`, eventID, organizersString)
	if err != nil {
		r.log.Printf("\nError in CreateEventOrganizers(repo) during inserting: %s\n", err.Error())
		return err
	}
	return nil
}
func (r *Repository) CreateEventInterests(eventID int64, interestList []string) error {
	stmt, err := r.db.Prepare(`INSERT INTO event_interests (event_id, interest_id) VALUES ($1,unnest(string_to_array($2, ','))::int)`)
	if err != nil {
		r.log.Printf("\nError preparing statement for inserting event interests: %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			r.log.Printf("\nError closing statement for inserting event interests: %s\n", err.Error())
		}
	}(stmt)
	interestsString := strings.Join(interestList, ",")
	_, err = stmt.Exec(eventID, interestsString)
	if err != nil {
		r.log.Printf("\nError inserting event interests: %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetAllEvents() ([]entity.Event, error) {
	query := `SELECT 
    a.*,
    b.interest_names
FROM 
    (
        SELECT DISTINCT
            e.*,
            STRING_AGG(u.username, ', ') AS organizer_usernames
        FROM 
            events e
        JOIN 
            event_organizers o ON e.id = o.event_id
        JOIN 
            users u ON o.organizer_id = u.id
        GROUP BY 
            e.id
    ) AS a
LEFT JOIN
    (
        SELECT DISTINCT
            e.id,
            STRING_AGG(ii.interest_name::TEXT, ', ') AS interest_names
        FROM 
            events e
        JOIN 
            event_interests i ON e.id = i.event_id
        JOIN 
            interests ii ON i.interest_id = ii.id
        GROUP BY 
            e.id
    ) AS b ON a.id = b.id

;`
	var allEvents []entity.Event
	rows, err := r.db.Query(query)
	if err != nil {
		r.log.Printf("\nError at the stage of query GetAllEvents(repo): %s\n", err.Error())
		return nil, err
	}
	for rows.Next() {
		event := entity.Event{}
		interestsString := ""
		organizersString := ""
		err = rows.Scan(&event.ID, &event.EventName, &event.FormatID, &event.Address, &event.CoordinateX, &event.CoordinateY, &event.Capacity, &event.Link, &event.Description, &event.PrivacyID, &event.CreatorID, &organizersString, &interestsString)
		if err != nil {
			r.log.Printf("\n error during scanning GetAllEvents(repo): %s\n", err.Error())
			return nil, err
		}
		event.EventInterestIDs = strings.Split(interestsString, ",")
		event.OrganizerIDs = strings.Split(organizersString, ",")
		allEvents = append(allEvents, event)
	}
	return allEvents, nil
}

func (r *Repository) GetEventsByInterests(interestList []string) ([]entity.Event, error) {
	query := `SELECT DISTINCT e.*
	FROM event_interests ei
	JOIN events e ON ei.event_id = e.id
	JOIN unnest(string_to_array($1, ',')) AS interests_list(interest_id) ON ei.interest_id = interests_list.interest_id::int`
	interestsString := strings.Join(interestList, ",")
	rows, err := r.db.Query(query, interestsString)
	if err != nil {
		r.log.Printf("\nError at the stage of Query GetEventByInterests(repo): %s\n", err.Error())
		return nil, err
	}
	filteredEvents := make([]entity.Event, 0)
	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.ID, &event.EventName, &event.FormatID, &event.Address, &event.CoordinateX, &event.CoordinateY, &event.Capacity, &event.Link, &event.Description, &event.PrivacyID)
		if err != nil {
			r.log.Printf("\n error during scanning GetEventByInterests(repo): %s\n", err.Error())
			return nil, err
		}
		filteredEvents = append(filteredEvents, event)
	}
	return filteredEvents, nil
}
