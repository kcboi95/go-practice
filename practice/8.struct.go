package practice

import (
	"strconv"
	"fmt"
)

// Hooman - Human
type Hooman struct {
	name       string
	age        uint8
	alive bool
}

// StillAlive will let you know is this human still alive?
func (h *Hooman) StillAlive() bool {
	return h.alive
}

// Struct - show how to use struct
func Struct() {
	var man = new(Hooman)

	man.name = "Dat"
	man.age = 23
	man.alive = true

	fmt.Println("------------------------------------------------------------")
	fmt.Println("|                         STRUCT                           |")
	fmt.Println("------------------------------------------------------------")
	fmt.Println(man)
	fmt.Println("Still alive? - " + strconv.FormatBool(man.StillAlive()))
}
