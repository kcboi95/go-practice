package practice

import (
	"fmt"
)

// Array - how to use
func Array() {
	var names [5]string

	names[0] = "Dat"
	names[1] = "John"
	names[2] = "Krista"
	names[3] = "Lorem"
	names[4] = "Wowy"

	fmt.Println("------------------------------------------------------------")
	fmt.Println("|                         ARRAY                            |")
	fmt.Println("------------------------------------------------------------")
	for _, name := range names {
		fmt.Println(name)
	}
}
