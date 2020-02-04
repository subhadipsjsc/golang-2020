package main

import "fmt"

// I : interface
type I interface {
	M()
}

// T1 : struct
type T1 struct{}

// M : method of T1
func (T1) M() {}

// T2 : struct
type T2 struct{}

// M : method of T2
func (T2) M() {}

func main() {
	var v1 I = T1{}
	v2, ok := v1.(T2)
	if !ok {
		fmt.Printf("ok: %v\n", ok)      // ok: false
		fmt.Printf("%v,  %T\n", v2, v2) // {},  main.T2
	}
}
