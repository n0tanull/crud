package entity

type User struct {
	ID       int    `json:"Id"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
