package main

import "fmt"

func main() {
	type NoKTP string
	type Maried bool

	var noKTP NoKTP = "12398123981293872183"
	var mariedStatus Maried = false
	fmt.Println(noKTP)
	fmt.Println(mariedStatus)
}
