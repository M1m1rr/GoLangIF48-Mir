package main

import "fmt"

func main() {
	var N, i, bas, hgt int
	var aTri float64

	fmt.Println("Jumlah Delta")
	fmt.Scan(&N)
	for i = 1; i <= N; i += 1 {
		fmt.Println("Alas Tinggi?")
		fmt.Scan(&bas, &hgt)
		aTri = float64(bas*hgt) / 2
		fmt.Print(aTri)
	}
}
