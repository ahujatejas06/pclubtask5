package models

type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Rollnumber int    `json:"roll_number"`
	Branch     string `json:"branch"`
}
