package main

import (
	"fmt"
	"strings"
)

//-- Talker: interface ---//
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

//-- rover : string  type ---//
type rover string

func (r rover) talk() string {
	return string(r)
}

//--- main function ----//
func main() {
	r := rover("whir whir") // rover is a talker
	shout(r)                // shout is a part of talker interface ,
}
