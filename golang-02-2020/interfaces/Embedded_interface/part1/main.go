package main

import "fmt"

//Mover : some comment on  interface
type Mover interface {
	Move()
}

//Car : some struct
type Car struct {
	price int
	Mover
	speed int
}

func main() {
	var c Car
	fmt.Printf("%v\n", c)

	/* 	As it printed out, Car has three fields.
	=>  {0 <nil> 0}
	Embedded interface is just an anonymous interface fields.
	*/
}
