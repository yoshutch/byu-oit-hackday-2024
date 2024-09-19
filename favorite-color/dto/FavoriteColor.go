package dto

type FavoriteColor struct {
	Id        int    `json:"id"`
	ProfileId int    `json:"profile_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
}
