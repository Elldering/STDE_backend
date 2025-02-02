package models

type Permisson struct {
	ID          int    `json:"id"`
	Codename    string `json:"codename"`
	Description string `json:"description"`
}
