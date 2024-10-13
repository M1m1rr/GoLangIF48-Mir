package main

import "fmt"

func main() {
	var sNum, bNum, i, tNums int
	fmt.Scan(&bNum, &sNum)

	for i = sNum; i <= bNum; i += 1 {
		fmt.Printf("%d ", i)
		tNums += i
	}
	fmt.Printf("\nTotal: %d", tNums)
}
