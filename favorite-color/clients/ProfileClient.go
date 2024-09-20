package clients

import (
	"byu.edu/hackday-favorite-color/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetProfile(id int) (*dto.Profile, error) {
	host := os.Getenv("PROFILE_SERVICE_HOST")
	resp, err := http.Get(fmt.Sprintf("http://%s/api/profile/%d", host, id))
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
