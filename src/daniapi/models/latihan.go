package models

type Latihan struct {
	Nama string `json:"nama"`
	Kelas string `json:"kelas"`
	Gender string `json:"gender"`
	Tanggal string `json:"tanggal"`
	Hobi []string `json:"hobi"`
}

type LatihanRequest struct {
	Token string `json:"token"`
}

type JsonErrorr struct {
	Output string
}