package main

import (
	"fmt"
	"os"
)

type Form struct {
	No        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	dataForm := []Form{
		{
			No:        "1",
			Nama:      "Shidqi",
			Alamat:    "Jalan Sudirman",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin Belajar Golang",
		},
		{
			No:        "2",
			Nama:      "Adiatma",
			Alamat:    "Jalan Patimura",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin Belajar Golang",
		},
		{
			No:        "3",
			Nama:      "Putra",
			Alamat:    "Jalan Nelayan",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin Belajar Golang",
		},
	}
	osArgsInput(dataForm)
}

func osArgsInput(dataForm []Form) {
	if len(os.Args) == 2 {
		noAbsen := os.Args[1]
		for _, loop := range dataForm {
			if loop.No == noAbsen {
				fmt.Println("Nama:", loop.Nama)
				fmt.Println("Alamat:", loop.Alamat)
				fmt.Println("Pekerjaan:", loop.Pekerjaan)
				fmt.Println("Alasan:", loop.Alasan)
			}
		}
	} else {
		fmt.Println("Maaf tidak ada data yang ditampilkan, Tolong masukkan no absen anda")
	}
}
