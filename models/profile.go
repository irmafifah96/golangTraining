package models

type Profile struct {
	Nama 	string
	Umur 	int
	Alamat 	Alamat
	Pekerjaan string
	Agama 	string
}

type Alamat struct {
	RT	int
	AlamatLengkap string
}