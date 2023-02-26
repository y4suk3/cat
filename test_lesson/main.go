package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func greeting(s string) string {
	ss := "Hi, " + s
	return ss
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(greeting("gopher"))
}
