package api

import (
    "fmt"
)

func JohnDoe() string {
    return "John Doe"
}

func main() {
    fmt.Println("Hello, " + JohnDoe() + "!")
}