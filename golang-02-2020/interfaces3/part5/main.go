package main

import "fmt"

// I1 : interface
type I1 interface {
	M()
}

// I2 : interface
type I2 interface {
	N()
}

// T : struct
type T struct{}

// M : method of T
func (T) M() {}

func main() {
	var v1 I1 = T{}
	var v2 I2
	v2, ok := v1.(I2)
	fmt.Printf("%T %v %v\n", v2, v2, ok)
}
