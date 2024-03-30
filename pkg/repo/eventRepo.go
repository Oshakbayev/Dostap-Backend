package repo

import (
	"database/sql"
	"github.com/lib/pq"
	"hellowWorldDeploy/pkg/entity"
	"strconv"
	"strings"
)

type EventInterface interface {
	CreateEvent(event *entity.Event) error
	GetEventsByInterests([]int) ([]entity.Event, error)
	CreateEventInterests(int, []int) error
	CreateEventOrganizers(int, []string) error
	GetAllEvents() ([]entity.Event, error)
}

func (r *Repository) CreateEvent(event *entity.Event) error {
	stmt, err := r.db.Prepare(`INSERT INTO events (event_name, format_id, address, coordinatex, coordinatey, capacity, link, description,privacy_id,creator_id,start_time,end_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10,$11,$12) RETURNING id`)
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
	err = stmt.QueryRow(event.EventName, event.FormatID, event.Address, event.CoordinateX, event.CoordinateY, event.Capacity, event.Link, event.Description, event.PrivacyID, event.CreatorID, event.StartTime, event.EndTime).Scan(&event.ID)
	if err != nil {
		r.log.Printf("\nError at the stage of data Inserting CreateEvent(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) CreateEventOrganizers(eventID int, organizers []string) error {
	organizersArray := pq.Array(organizers)
	_, err := r.db.Exec(`INSERT INTO event_organizers (event_id, organizer_id)
SELECT $1 AS event_id, u.id AS organizer_id
FROM users u
WHERE u.username = ANY($2)
`, eventID, organizersArray)
	if err != nil {
		r.log.Printf("\nError in CreateEventOrganizers(repo) during inserting: %s\n", err.Error())
		return err
	}
	return nil
}
func (r *Repository) CreateEventInterests(eventID int, interestList []int) error {
	interestArray := pq.Array(interestList)
	_, err := r.db.Exec(`
   INSERT INTO event_interests (event_id, interest_id)
        VALUES ( $1, unnest($2::int[]))
`, eventID, interestArray)
	if err != nil {
		r.log.Printf("\nError in CreateEventInterests(repo): %s\n", err.Error())
		return err
	}
	//log.Println(res.RowsAffected())
	return nil
}

func (r *Repository) GetAllEvents() ([]entity.Event, error) {
	query := `
SELECT t1.*, COALESCE(interests, '{NULL}') AS interests
FROM (
  SELECT events.*, ARRAY_AGG(u.username) as organizers
  FROM events
  LEFT JOIN event_organizers AS eo ON eo.event_id = events.id
  LEFT JOIN users AS u ON u.id = eo.organizer_id
  GROUP BY events.id
) t1 LEFT JOIN (
  SELECT event_id, ARRAY_AGG(interest_id::bigint) as interests
  FROM event_interests
  JOIN interests ON interests.id = event_interests.interest_id
  GROUP BY event_id
) t2
ON t1.id = t2.event_id
ORDER BY t1.id;
`
	var allEvents []entity.Event
	rows, err := r.db.Query(query)
	if err != nil {
		r.log.Printf("\nError at the stage of query GetAllEvents(repo): %s\n", err.Error())
		return nil, err
	}
	for rows.Next() {
		var orgArr string
		var intrsArr string
		event := entity.Event{}
		err = rows.Scan(&event.ID, &event.EventName, &event.FormatID, &event.Address, &event.CoordinateX, &event.CoordinateY, &event.Capacity, &event.Link, &event.Description, &event.PrivacyID, &event.CreatorID, &event.StartTime, &event.EndTime, &orgArr, &intrsArr)
		if err != nil {
			r.log.Printf("\n error during scanning GetAllEvents(repo): %s\n", err.Error())
			return nil, err
		}
		orgArr = strings.Trim(orgArr, "{}NULL")
		if orgArr != "" {
			event.OrganizerIDs = strings.Split(orgArr, ",")
		}

		intrsArr = strings.Trim(intrsArr, "{}NULL")
		for _, val := range strings.Split(intrsArr, ",") {
			if id, err := strconv.Atoi(string(val)); err == nil {
				event.InterestIDs = append(event.InterestIDs, id)
			}
		}
		//event.EventInterestIDs = strings.Split(intrsArr, ",")
		allEvents = append(allEvents, event)
	}
	return allEvents, nil
}

func (r *Repository) GetEventsByInterests(interestList []int) ([]entity.Event, error) {
	interestArray := pq.Array(interestList)
	query := `SELECT t1.*, COALESCE(interests, '{NULL}') AS interests
FROM (
  SELECT events.*, ARRAY_AGG(u.username) as organizers
  FROM events
  LEFT JOIN event_organizers AS eo ON eo.event_id = events.id
  LEFT JOIN users AS u ON u.id = eo.organizer_id
  GROUP BY events.id
) t1 LEFT JOIN (
  SELECT event_id, ARRAY_AGG(interest_id::bigint) as interests
  FROM event_interests
  JOIN interests ON interests.id = event_interests.interest_id
  GROUP BY event_id
) t2
ON t1.id = t2.event_id
ORDER BY t1.id`

	rows, err := r.db.Query(query, interestArray)
	if err != nil {
		r.log.Printf("\nError at the stage of Query GetEventByInterests(repo): %s\n", err.Error())
		return nil, err
	}
	filteredEvents := make([]entity.Event, 0)
	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.ID, &event.EventName, &event.FormatID, &event.Address, &event.CoordinateX, &event.CoordinateY, &event.Capacity, &event.Link, &event.Description, &event.PrivacyID, &event.CreatorID, &event.StartTime, &event.EndTime)
		if err != nil {
			r.log.Printf("\n error during scanning GetEventByInterests(repo): %s\n", err.Error())
			return nil, err
		}
		filteredEvents = append(filteredEvents, event)
	}
	return filteredEvents, nil
}
