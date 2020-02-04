package main

import "fmt"

// T : type struct
type T struct {
	name string
}

func main() {
	v1 := struct{ name string }{"foo"} // initiate and assign value of struct on the fly
	fmt.Printf("%T\n", v1)             // struct { name string }
	var v2 T                           // v2 is of type T struct
	v2 = v1                            // assign v2 value from v1
	fmt.Printf("%T\n", v2)             // main.T
}
