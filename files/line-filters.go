package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// $ echo 'hello'   > /tmp/lines
// $ echo 'filter' >> /tmp/lines
// Then use the line filter to get uppercase lines.

// $ cat /tmp/lines | go run line-filters.go
func main() {

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {

        ucl := strings.ToUpper(scanner.Text())

        fmt.Println(ucl)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
