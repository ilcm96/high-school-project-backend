package db

type User struct {
	ID   string `json:"id"`
	PW   string `json:"pw"`
	Name string `json:"name"`
}
