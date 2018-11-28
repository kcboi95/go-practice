package practice

import (
	"strconv"
	"fmt"
)

// Pointers is a function show you how to work with pointer in Go
func Pointers() {
	x := 10

	fmt.Println("------------------------------------------------------------")
	fmt.Println("|                        POINTERS                          |")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Value before changed:"+strconv.Itoa(x))
	changeValue(&x)
	fmt.Println("Value after changed:"+strconv.Itoa(x))
}

func changeValue(x *int) {
	*x = 5
}
