package model

type Dokter struct {
	Id          int    `json:"id" gorm:"autoIncrement;primaryKey"`
	Nama        string `json:"nama"`
	Email       string `json:"email"`
	SpesialisID int    `json:"spesialis_id"`
}
