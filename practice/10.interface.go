package practice

import (
	"fmt"
)

// Bird - Just bird stuffs
type Bird interface {
	sound() string
}

// Duck - ...
type Duck struct {
}

// Owl - Night? Day?
type Owl struct {
}

func (o *Owl) sound() string {
	return "Kuckooo"
}

func (d *Duck) sound() string {
	return "Quack quack"
}

func soundOf(obj Bird) string {
	return obj.sound()
}

// Interface - what the interface?
func Interface() {
	duck := new(Duck)
	owl := new(Owl)

	fmt.Println("------------------------------------------------------------")
	fmt.Println("|                        INTERFACE                         |")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Duck sound: " + soundOf(duck))
	fmt.Println("Owl sound: " + soundOf(owl))
}
