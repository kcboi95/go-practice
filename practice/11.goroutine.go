package practice

import (
	"fmt"
)

func routine(number int) {
	for i := 0; i < number; i++ {
		fmt.Println(i)
	}
}

// Goroutine - go go go
func Goroutine() {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("|                       GO ROUTINE                         |")
	fmt.Println("------------------------------------------------------------")
	for i := 0; i < 10; i++ {
		go routine(i)
	}
}
