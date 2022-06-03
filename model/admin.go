package model

type Admin struct {
	Id    int    `json:"id" gorm:"autoIncrement;primaryKey"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
}
