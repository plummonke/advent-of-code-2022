package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "aoc2022/utility"
)

func main() {
    p, err := filepath.Abs(os.Args[1])
    if err != nil {
        log.Fatal(err.Error())
    }

    text, err := utility.ReadFile(p)
    if err != nil {
        log.Fatal(err.Error())
    }
}
