package main

import (
        "fmt"
)

func main() {
        hello, _ := Asset("data/hello.txt")
        fmt.Printf("%s\n", hello)
}
