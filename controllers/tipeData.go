package controllers

func tipeDataSesi1() {
	var namaHewan string = "kupu-kupu"
	var sisiPersegi int 
	sisiPersegi = 5
	nama, umur, alamat  := "Dicky", 25, "Jakarta"
	namaUser := 5.5
	
	fmt.Println(namaHewan)
	fmt.Println(sisiPersegi)
	fmt.Println(len(namaHewan))
	fmt.Println("1+1 = ", 1+1)
	fmt.Println(nama)

	slice1 := []int{1,4,5,6,7}
	//udah ditentuin panjang array
	array1 := [2]int{0,1}

	//[tipedata], value nya apa
	map1 := map[string]int{
		"januari": 10,
		"februari": 5,
	}
	//tipe data berbeda beda
	map2 := map[string]interface{}{
		"januari": 10,
		"februari": 5,
	}

	map3 := map[string]map[string]int{
		"indonesia": map[string]int{
			"beras": 100,
			"jagung": 5,
		}
		"malaysia": map[string]int{
			"beras": 200,
			"jagung": 5,
		}
	}
	fmt.Println(slice1)
	fmt.Println(array1)
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(profile)
	fmt.Println(map3)

}