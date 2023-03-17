package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	exampleWords("selamat malam")
}

func exampleWords(text string) {
	mapAwal, mapAkhir := map[string]int{}, make(map[string]int)

	pecahKata := strings.Split(text, "")

	for _, w := range pecahKata {
		fmt.Println(w)
		mapAwal[w]++
	}

	keys := make([]string, 0, len(mapAwal))

	for k := range mapAwal {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		mapAkhir[k] = mapAwal[k]
	}

	fmt.Println(mapAkhir)
}
