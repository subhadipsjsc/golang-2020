package main

import "fmt"

// I : interface defination
type I interface {
	walk()
	quack()
}

// S : struct : defination
type S struct{}

func (s S) walk()  {}
func (s S) quack() {}

func main() {
	var i I
	i = S{}

	//Syntax of type assertion is defined as : PrimaryExpression.(Type)
	fmt.Println(i.(interface{ walk() }))

	// error => fmt.Println(S{}.(interface{ walk() }))

}
