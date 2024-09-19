package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type ProfileRepo struct {
	db *sql.DB
}

type Profile struct {
	Id        int
	FirstName string
	LastName  string
}

func NewProfileRepo(username string, password string, port int, database string) (*ProfileRepo, error) {
	connStr := fmt.Sprintf("user=%s password=%s port=%d database=%s sslmode=disable", username, password, port, database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	//rows, err := db.Query("Select * from profiles")
	//log.Printf("query resuls: $s", rows.)
	return &ProfileRepo{db: db}, nil
}

func (r ProfileRepo) GetProfile(id int) (*Profile, error) {
	row := r.db.QueryRow(`SELECT * from profiles where id = $1`, id)
	var profile Profile
	err := row.Scan(&profile.Id, &profile.FirstName, &profile.LastName)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r ProfileRepo) UpdateProfile(profile *Profile) error {
	stmt, err := r.db.Prepare("UPDATE profiles set first_name = $2, last_name = $3 where id = $1")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(profile.Id, profile.FirstName, profile.LastName)
	//_, err := r.db.Exec("UPDATE profiles set first_name = $2 and last_name = $3 where id = $1", profile.Id, profile.FirstName, profile.LastName)
	if err != nil {
		return err
	}
	log.Printf("exect result: %s", result)
	return nil
}
