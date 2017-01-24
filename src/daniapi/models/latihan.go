package models

type Latihan struct {
	Nama string `json:"nama"`
	Kelas string `json:"kelas"`
	Gender string `json:"gender"`
	Tanggal string `json:"tanggal"`
	Hobi []Category `json:"hobi"`
}

type Category struct {
	Hobi1 string `json:"hobi_1"`
	Hobi2 string `json:"hobi_2"`
}

type LatihanRequest struct {
	Token string `json:"token"`
}

type JsonErrorr struct {
	Output string
}