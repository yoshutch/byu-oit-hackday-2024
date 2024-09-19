package db

import (
	"byu.edu/hackday-favorite-color/dto"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type FavColorRepo struct {
	db *sql.DB
}

func NewFavColorRepo(username string, password string, port int, database string) (*FavColorRepo, error) {
	connStr := fmt.Sprintf("user=%s password=%s port=%d database=%s sslmode=disable", username, password, port, database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &FavColorRepo{db: db}, nil
}

func (r FavColorRepo) GetFavoriteColor(id int) (*dto.FavoriteColor, error) {
	row := r.db.QueryRow(`SELECT * from hackday_colors.favorite_colors where id = $1`, id)
	var favoriteColor dto.FavoriteColor
	err := row.Scan(&favoriteColor.Id, &favoriteColor.ProfileId, &favoriteColor.Name, &favoriteColor.Color)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &favoriteColor, nil
}

func (r FavColorRepo) UpdateFavoriteColor(favColor *dto.FavoriteColor) error {
	stmt, err := r.db.Prepare("UPDATE hackday_colors.favorite_colors set name = $2, color = $3 where id = $1")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(favColor.Id, favColor.Name, favColor.Color)
	if err != nil {
		return err
	}
	log.Printf("exect result: %s", result)
	return nil
}
