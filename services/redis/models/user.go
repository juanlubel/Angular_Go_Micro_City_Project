package models

type UserDTO struct {
	Slug  string `json:"slug"`
	NCard int	 `json:"nCard"`
	Token string `json:"token"`
}