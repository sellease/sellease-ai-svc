package models

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	TeamName string `json:"team_name"`
	Passcode string `json:"passcode"`
	IsActive bool   `json:"is_active"`
}
