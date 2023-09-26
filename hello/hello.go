package main

import (
	"fmt"
	"strings"
)

var i, j int = 1, 2
var s = "foo bar baz"

func main() {
	fmt.Println("Hello,test!")
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	b := make(map[string]int)

	a := strings.Fields(s)
	fmt.Println(a)
	fmt.Printf("Fields are: %q", strings.Fields("foo bar baz"))
}
