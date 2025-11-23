package dto

type Team struct {
	Name    string  `json:"name"`
	Members []*User `json:"members"`
}
