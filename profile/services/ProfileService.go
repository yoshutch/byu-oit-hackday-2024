package services

import (
	"byu.edu/hackday-profile/db"
	"log"
)

type ProfileService struct {
	repo *db.ProfileRepo
}

func NewProfileService(repo *db.ProfileRepo) (*ProfileService, error) {
	return &ProfileService{repo}, nil
}

func (s ProfileService) LoadProfile(id int) (*db.Profile, error) {
	// authorization? business logic
	profile, err := s.repo.GetProfile(id)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s ProfileService) SaveProfile(id int, firstName string, lastName string) error {
	profile, err := s.LoadProfile(id)
	if err != nil {
		log.Printf("Error loading profile: %s", err)
	}
	if profile == nil {
		// insert
		return nil
	} else {
		updated := &db.Profile{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		return s.repo.UpdateProfile(updated)
	}
	// TODO send internal event to eventbus
}
