package main

import (
	"fmt"
	"go-practice/practice"
)

func main() {
	for true {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("| Please choose one value in the list below                |")
		fmt.Println("| to see the result or use `Ctrl + C` to exit the program  |")
		fmt.Println("------------------------------------------------------------")
		fmt.Println("1 . Hello World")
		fmt.Println("2 . Data types")
		fmt.Println("3 . Variables")
		fmt.Println("4 . Constants")
		fmt.Println("5 . Decision making")
		fmt.Println("6 . Loop")
		fmt.Println("7 . Pointers")
		fmt.Println("8 . Array")
		fmt.Println("9 . Struct")
		fmt.Println("10 . Interface")
		fmt.Println("11 . Goroutine")
		fmt.Print("Value:")

		var value int
		_, err := fmt.Scanf("%d\n", &value)

		// Capture error
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		switch value {
		case 1:
			practice.HelloWorld()
		case 2:
			practice.DataTypes()
		case 3:
			practice.Variables()
		case 4:
			practice.Constants()
		case 5:
			practice.DecisionMaking()
		case 6:
			practice.Loop()
		case 7:
			practice.Pointers()
		case 8:
			practice.Array()
		case 9:
			practice.Struct()
		case 10:
			practice.Interface()
		case 11:
			practice.Goroutine()
		}
	}
}
