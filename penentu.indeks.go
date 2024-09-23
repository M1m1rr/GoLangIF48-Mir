package main

import (
	"fmt"
)

//PENENTU INDEKS NILAI//

func main() {
	var nilai int
	var index, pon string

	fmt.Printf("Grade: ")
	fmt.Scan(&nilai)

	if nilai > 80 && nilai <= 100 {
		index = "A"
	} else if nilai > 70 && nilai <= 80 {
		index = "AB"
	} else if nilai > 65 && nilai <= 70 {
		index = "B"
	} else if nilai > 60 && nilai <= 65 {
		index = "BC"
	} else if nilai > 50 && nilai <= 60 {
		index = "C"
	} else if nilai > 40 && nilai <= 50 {
		index = "D"
	} else if nilai <= 40 {
		index = "E"
	} else {
		index = "INPUT INVALID"
	}

	if nilai > 40 && nilai <= 100 {
		pon = "PASSED"
	} else if nilai <= 40 {
		pon = "DID NOT PASS"
	} else {
		pon = "INPUT INVALID"
	}

	fmt.Printf("You [%v] with the index [%v]", pon, index)
}
