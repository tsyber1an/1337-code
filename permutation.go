package main

import (
	"fmt"
)

var n = 2
var m = 3
var a map[int]int

func out() {
	for _, v := range a {
		fmt.Print(v)
	}
	fmt.Println()
}
func rec(idx int) {
	if idx == n {
		out()
		return
	}

	for i := 1; i <= m; i++ {
		a[idx] = i
		rec(idx + 1)
	}
}

func main() {
	a = make(map[int]int)
	rec(0)
}
