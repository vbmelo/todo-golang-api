package models

type Todo struct {
	ID		int    `json:"id"`
	Title string `json:"title"`
	Description string `json:"descritpion"`
	Done bool `json:"done"`
}
