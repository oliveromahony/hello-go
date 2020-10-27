package main

import "fmt"

func vals() (int, int, int) {
    return 3, 7, 5
}

func main() {

    a, b, d := vals()
    fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)

	// original example only used one underscore
	_, _, c := vals()

	fmt.Println(c)
}