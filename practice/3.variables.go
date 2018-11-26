package practice

import (
	"fmt"
	"strconv"
)

// Variables - show variables result
func Variables() {
	// First way
	var name string
	name = "Dat"
	// Second way
	var age = 23
	// Third way
	hometown := "Hanoi"

	fmt.Println("------------------------------------------------------------")
	fmt.Println("|                       VARIABLES                          |")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Name: " + name)
	fmt.Println("Age: " + strconv.Itoa(age))
	fmt.Println("Hometown: " + hometown)
}
