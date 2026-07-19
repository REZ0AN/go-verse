package main

import (
    "fmt"
    "os"
)

func main() {
    name := "Go"
    fmt.Println("Hello,", name)
    fmt.Println("Args:", os.Args)
}