package main

import "fmt"

func main() {
	var T float64 = 20
	var r float64 = 4
	const phi = 3.14

	LP := (2 * phi * r * r) + (2 * phi * r * T)
	fmt.Println(LP)

}
