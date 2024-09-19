package services

import (
	"byu.edu/hackday-profile/db"
	"byu.edu/hackday-profile/dto"
	"byu.edu/hackday-profile/events"
	"log"
)

type ProfileService struct {
	repo         *db.ProfileRepo
	eventAdapter *events.EventAdapter
}

func NewProfileService(repo *db.ProfileRepo, adapter *events.EventAdapter) (*ProfileService, error) {
	return &ProfileService{repo, adapter}, nil
}

func (s ProfileService) LoadProfile(id int) (*dto.Profile, error) {
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
		updated := &dto.Profile{
			Id:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		err := s.repo.UpdateProfile(updated)
		if err != nil {
			return err
		}
		// send event
		err = s.eventAdapter.SendProfileUpdatedEvent(*updated)
		if err != nil {
			return err
		}
		return nil
	}
	// TODO send internal event to eventbus
}
