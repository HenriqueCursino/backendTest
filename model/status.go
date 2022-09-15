package model

type Status struct {
	ID   int    `gorm:"primaryKey; autoIncrement" json:"id"`
	Name string `gorm:"size:15" json:"name"`
}
