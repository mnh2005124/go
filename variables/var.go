package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	name, course = "Manish", "Kaushik"
	module, clip = "4", 2
)

func main() {
	fmt.Println("Name and Course are set to ", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type ", reflect.TypeOf(name))
	fmt.Println("Module is of type ", reflect.TypeOf(module))
	fmt.Println("Memory address of varable cousre", &course)
	iMoudle, err := strconv.Atoi(module)
	if err == nil {
		total := iMoudle + clip
		fmt.Println("Module plus clip equals", total)
	}
}
