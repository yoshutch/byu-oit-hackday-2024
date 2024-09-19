package clients

import (
	"byu.edu/hackday-profile/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetProfile(id int) (*dto.Profile, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/api/profile/%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var profile dto.Profile
	err = json.Unmarshal(bytes, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
