package practice

import (
	"fmt"
)

// Change age value to see what happen
var age = 23

// DecisionMaking show how to use if else statements
func DecisionMaking() {

	fmt.Println("------------------------------------------------------------")
	fmt.Println("|                     DECISION MAKING                      |")
	fmt.Println("------------------------------------------------------------")

	// if - else if - else
	if age == 18 {
		fmt.Println("Your age is 18")
	} else if age > 18 {
		fmt.Println("Your age is greater than 18")
	} else {
		fmt.Println("Your age is smaller than 18")
	}

	// switch
	switch {
	case age == 18:
		fmt.Println("Eighteen")
	case age > 18:
		fmt.Println("> 18")
	case age < 18:
		fmt.Println("< 18")
	default:
		fmt.Println("Not sure")
	}

	// switch another style
	switch age {
	case 18:
		fmt.Println("Ha")
	case 19:
		fmt.Println("Comma")
	case 20:
		fmt.Println("Ho")
	default:
		fmt.Println("Yolo")
	}
}
