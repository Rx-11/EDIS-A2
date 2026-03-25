package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;autoIncrement"`
	UserID   string `json:"userId" gorm:"unique;type:varchar(255)"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Phone    string `json:"phone" gorm:"type:varchar(255)"`
	Address  string `json:"address" gorm:"type:varchar(255)"`
	Address2 string `json:"address2" gorm:"type:varchar(255)"`
	City     string `json:"city" gorm:"type:varchar(255)"`
	State    string `json:"state" gorm:"type:varchar(255)"`
	Zipcode  string `json:"zipcode" gorm:"type:varchar(255)"`
}
