package services

import (
	"byu.edu/hackday-favorite-color/db"
	"byu.edu/hackday-favorite-color/dto"
	"fmt"
)

type FavColorService struct {
	repo *db.FavColorRepo
}

func NewFavColorService(repo *db.FavColorRepo) (*FavColorService, error) {
	return &FavColorService{repo}, nil
}

func (s FavColorService) LoadFavColor(id int) (*dto.FavoriteColor, error) {
	return s.repo.GetFavoriteColor(id)
}

func (s FavColorService) UpdateName(id int, profile dto.Profile) error {
	existing, err := s.repo.GetFavoriteColor(id)
	if err != nil {
		return nil
	}
	name := fmt.Sprintf("%s %s", profile.FirstName, profile.LastName)
	favColor := &dto.FavoriteColor{
		Id:        id,
		ProfileId: existing.ProfileId,
		Name:      name,
		Color:     existing.Color,
	}
	return s.repo.UpdateFavoriteColor(favColor)
}
