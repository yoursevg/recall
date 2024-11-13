package models

type Command struct {
	Command     string `json:"command"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
}
