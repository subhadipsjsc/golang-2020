package main

import "fmt"

//I1 : interface
type I1 interface {
	M1()
}

//I2 : interface
type I2 interface {
	M1()
}

// T : struct
type T struct{}

//M1 : method
func (T) M1() {}

func main() {
	var v1 I1 = T{}

	// fmt.Println(T(v1)) =>  cannot convert v1 (type I1) to type T: need type assertion

	var v2 I2 = v1
	var x = v2
	fmt.Printf("%T\n", x)
}
