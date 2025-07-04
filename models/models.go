package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Active   bool   `json:"active"`
	Role     int    `json:"role"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
