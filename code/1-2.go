package main

import (
	"fmt"
	"os"
	"strings"
)

func one() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func two() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func three() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func four() {
	fmt.Println(os.Args[1:])
}

func main() {
	three()
	four()
}
