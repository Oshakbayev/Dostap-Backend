package repo

import (
	"github.com/lib/pq"
	"hellowWorldDeploy/pkg/entity"
)

type InterestInterface interface {
	GetAllInterests() ([]entity.Interest, error)
	GetUserInterests(int) ([]entity.Interest, error)
	CreateUserInterests(int, []int) error
	DeleteUserInterests(int) error
}

func (r *Repository) CreateUserInterests(userId int, interests []int) error {
	query := `INSERT INTO user_interests (user_id, interest_id) VALUES  ($1, unnest($2::int[]))`
	_, err := r.db.Exec(query, userId, pq.Array(interests))
	if err != nil {
		r.log.Printf("\nError in CreateUserInterests(repo): %s\n", err.Error())

	}
	return err
}

func (r *Repository) GetAllInterests() ([]entity.Interest, error) {
	rows, err := r.db.Query(`SELECT  * FROM interests`)
	if err != nil {
		r.log.Printf("\nError in GetAllInterests(repo) during selecting all interests: %s\n", err.Error())
		return nil, err
	}
	var allInterests []entity.Interest
	for rows.Next() {
		interest := entity.Interest{}
		if err := rows.Scan(&interest.ID, &interest.Name); err != nil {
			r.log.Printf("\nError in GetAllInterests(repo) during scanning  interest: %s\n", err.Error())
			return nil, err
		}
		allInterests = append(allInterests, interest)
	}
	return allInterests, nil
}

func (r *Repository) GetUserInterests(userId int) ([]entity.Interest, error) {
	query := `SELECT * FROM user_intersts WHERE user_id = $1`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		r.log.Printf("\n error in GetUserInterests(repo): %s\n", err.Error())
		return nil, err
	}
	var userInterests []entity.Interest
	for rows.Next() {
		var interest entity.Interest
		err := rows.Scan(&interest.ID, &interest.Name, &interest.Category)
		if err != nil {
			r.log.Printf("\n error in GetUserInterests(repo): %s\n", err.Error())
			return nil, err
		}
		userInterests = append(userInterests, interest)
	}
	return userInterests, err
}

func (r *Repository) DeleteUserInterests(userId int) error {
	query := `DELETE FROM user_interests WHERE user_id = $1`
	_, err := r.db.Exec(query, userId)
	if err != nil {
		r.log.Printf("\n error in DeleteUserInterests(repo): %s\n", err.Error())
	}
	return err
}
