package models

type Barang struct {
	Id   int64  `json:"id"`
	Nama string `json: "name"` //definisi tipe model
	Kode string `json: "kode"` //definisi tipe model
}
