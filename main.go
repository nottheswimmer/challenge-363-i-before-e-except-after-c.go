package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(check("a"))
	fmt.Println(check("zombie"))
	fmt.Println(check("transceiver"))
	fmt.Println(check("veil"))
	fmt.Println(check("icier"))
}

func check(s string) bool {
	if strings.Contains(s, "cie") {
		return false
	} else if strings.Count(s, "ei") != strings.Count(s, "cei") {
		return false
	} else {
		return true
	};
}