package main

import "fmt"

// I : interface defination //
type I interface {
	walk()
}

// A : struct
type A struct{}

//walk: method of A
func (a A) walk() {}

// B : struct
type B struct{}

//walk: method of B
func (b B) walk() {}

func main() {
	var i I
	i = A{} // dynamic type of i is A
	fmt.Printf("%T \n", i.(A))
	i = B{} // dynamic type of i is B
	fmt.Printf("%T \n", i.(B))
}
