package main

import (
	"fmt"
	"strings"
)

//-- laser type ---//
type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

//-- Talker: interface ---//
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

//-- startship type---//
type starship struct {
	laser
}

func main() {

	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)

}
