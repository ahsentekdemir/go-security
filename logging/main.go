package main

import (
    "fmt"
    "log"
    "os"
)

func main() {

    f, err := os.Create("log.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString("hi\n")

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("ggs")
}