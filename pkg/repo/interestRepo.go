package repo

import "hellowWorldDeploy/pkg/entity"

type InterestInterface interface {
	GetAllInterests() ([]entity.Interest, error)
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
