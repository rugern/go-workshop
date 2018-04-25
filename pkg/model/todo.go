package model

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Checked     bool   `json:"checked"`
}
