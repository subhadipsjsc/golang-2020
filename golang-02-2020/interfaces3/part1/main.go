package main

import "fmt"

// T1 : struct
type T1 struct {
	name string
}

// T2 : struct
type T2 struct {
	name string
}

func main() {
	vs := []interface{}{T2(T1{"foo"}), string(322), []byte("abł")}

	// -- the first element  T2(T1{"foo"}) => is "foo" in T1 assing into T2 => because of empty interface {}
	for _, v := range vs {
		fmt.Printf("%v %T\n", v, v)
	}
}
