package main

import (
    "fmt"
    "os"
)

// go run exit.go
//
// go build exit.go
// ./exit
func main() {

    defer fmt.Println("!")

    os.Exit(3)
}
