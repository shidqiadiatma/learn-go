package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// result text
	calculateWords("selamat malam")
}

func calculateWords(text string) {
	// variabel awal dan akhir
	elementMapAwal, elementMapAkhir := map[string]int{}, make(map[string]int)

	// dipotong menjadi slice
	pecahPerkata := strings.Split(text, "")

	// di looping simpan ke variabel awal
	for _, w := range pecahPerkata {
		fmt.Println(w)
		elementMapAwal[w]++
	}

	keys := make([]string, 0, len(elementMapAwal))

	//for looping simpan ke variabel keys
	for k := range elementMapAwal {
		keys = append(keys, k)
	}

	// mengurutkan dari a sampai z
	sort.Strings(keys)

	// menyimpan variabel akhir
	for _, k := range keys {
		//fmt.Printf("%s = %d\n", k, elementMapAwal[k])
		elementMapAkhir[k] = elementMapAwal[k]
	}

	// menampilkan hasil output
	fmt.Println(elementMapAkhir)
}
