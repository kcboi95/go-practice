package practice

import "fmt"

// Loop show how to use for loop
func Loop() {

	fmt.Println("------------------------------------------------------------")
	fmt.Println("|                           LOOP                           |")
	fmt.Println("------------------------------------------------------------")

	max := 100

	for i := 0; i < max; i++ {
		fmt.Printf(".")
	}

	fmt.Println("")
}
