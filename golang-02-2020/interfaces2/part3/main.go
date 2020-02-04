package main

import "fmt"

//I :
type I interface {
	walk()
}

// A : struct type
type A struct {
	name string
}

func (a A) walk() {}

//B :
type B struct {
	name string
}

func (b B) walk() {}

func main() {
	var i I
	i = A{name: "foo"}

	// ---, ok syntex :  To handle failures gracefully
	valA, okA := i.(A)
	fmt.Printf("%#v %#v\n", valA, okA)

	// ---, ok syntex :  To handle failures gracefully
	valB, okB := i.(B)
	fmt.Printf("%#v %#v\n", valB, okB)
}
