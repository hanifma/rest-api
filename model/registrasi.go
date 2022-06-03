package model

type Registrasi struct {
	NIK             string `json:"nik"`
	AdminID         int    `json:"admin_id"`
	DokterID        int    `json:"dokter_id"`
	JenisPenyakitID int    `json:"jenis_penyakit_id"`
	Keluhan         string `json:"keluhan"`
}
