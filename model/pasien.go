package model

type Pasien struct {
	NIK    string `json:"nik" gorm:"primaryKey"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	NoHp   string `json:"no_hp"`
}
