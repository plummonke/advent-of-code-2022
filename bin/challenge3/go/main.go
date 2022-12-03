package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    
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

    textArr := strings.Split(text, "\n")

    var priorities []int
    for _, each := range textArr {
        compartments := splitCompartments(each)
        if len(each) > 0 {
            priorities = append(
                priorities, 
                prioritizeChar(compareCompartments(compartments[0], compartments[1])[0]),
            )
        }
    }

    fmt.Printf("Part 1 answer: %d\n", sum[int](priorities...))

    var badges []int
    for i := 0; i < len(textArr); i += 3 {
        if len(textArr) - 1 < i + 3 {
            break
        }
        sacks := textArr[i:i + 3]
        // To find the common character between all three lines, use compareCompartments to assess the
        // first two lines and then create a new []string with the output of compareCompartments and the
        // third line and run compareCompartments again.
        common := compareCompartments(
            string(compareCompartments(sacks[0], sacks[1])), 
            sacks[2],
        )
        badges = append(badges, prioritizeChar(common[0]))
    }

    fmt.Printf("Part 2 answer: %d\n", sum[int](badges...))
}

func splitCompartments(x string) []string {
    var mid int = len(x) / 2
    return []string{
        x[:mid],
        x[mid:],
    }
}

// compareCompartments returns an array of all runes found in common between the strings.
func compareCompartments(a, b string) []rune {
    var out []rune
    for _, char1 := range a {
        for _, char2 := range b {
            if char1 == char2 {
                out = append(out, char1)
            }
        }
    }
    return out
}

func prioritizeChar(a rune) int {
    alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    return strings.IndexRune(alpha, a) + 1
}

type number interface {
    int|int8|int16|int32|int64|uint|uint8|uint16|uint32|uint64|float32|float64
}

func sum[T number](x ...T) (total T) {
    for _, num := range x {
        total += num
    }
    return
}
