package models

type Book struct {
	ISBN        string  `json:"ISBN" gorm:"primary_key;type:varchar(255)"`
	Title       string  `json:"title" gorm:"type:varchar(255)"`
	Author      string  `json:"Author" gorm:"type:varchar(255)"`
	Description string  `json:"description" gorm:"type:text"`
	Genre       string  `json:"genre" gorm:"type:varchar(255)"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2)"`
	Quantity    int     `json:"quantity" gorm:"type:int"`
	Summary     *string `json:"summary,omitempty" gorm:"type:text"`
}
