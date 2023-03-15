package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	calculateWords("selamat malam") // result text
}

func calculateWords(text string) {
	elementMapAwal, elementMapAkhir := map[string]int{}, make(map[string]int) // variabel awal dan akhir

	pecahPerkata := strings.Split(text, "") // dipotong menjadi slice

	for _, w := range pecahPerkata { // di looping simpan ke variabel awal
		fmt.Println(w)
		elementMapAwal[w]++
	}

	keys := make([]string, 0, len(elementMapAwal))

	for k := range elementMapAwal { //for looping simpan ke variabel keys
		keys = append(keys, k)
	}

	sort.Strings(keys) // mengurutkan dari a sampai z

	for _, k := range keys { // menyimpan variabel akhir
		elementMapAkhir[k] = elementMapAwal[k] //fmt.Printf("%s = %d\n", k, elementMapAwal[k])
	}

	fmt.Println(elementMapAkhir) // menampilkan hasil output
}
