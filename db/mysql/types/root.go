package types

type User struct {
	UserName    string   `json:"user_name"`
	Image       []string `json:"image"`
	Description string   `json:"description"`
	Hobby       []string `json:"hobby"`
	Latitude    float64  `json:"latitude"`
	Hardness    float64  `json:"hardness"`
}
