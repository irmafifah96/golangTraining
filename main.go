package main

import (
	"fmt"
	"training-go/sesi1/models"
	"encoding/json"
)

func main() {

	profile := models.Profile{
		Nama: "irma",
		Umur: 23,
		Alamat: models.Alamat{
			RT: 01,
			AlamatLengkap: "Slipi 3",
		},
	}

	responAlamat, _ := json.Marshal(profile)
	fmt.Println(string(responAlamat))
	
}