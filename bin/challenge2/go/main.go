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

    pairs := strings.Split(text, "\n")

    rock := Throw{"Rock", 1}
    paper := Throw{"Paper", 2}
    scissors := Throw{"Scissors", 3}

    score := Play(
        map[string]Throw {
            "A": rock,
            "B": paper,
            "C": scissors,
            "X": rock,
            "Y": paper,
            "Z": scissors,
        }, 
        pairs)
    fmt.Printf("Final score part 1: %d\n", score)
    
    score = PlayRedux(
        map[string]Throw {
            "A": rock,
            "B": paper,
            "C": scissors,
        }, 
        pairs)
    fmt.Printf("Final score part 2: %d\n", score)
}

type Throw struct {
    Name string
    Score int
}

func (t Throw) Compare(x Throw) int {
    if t.Score == x.Score {
        return 3
    }

    // Determine if self wins
    if x.Score == t.Score - 1 || (t.Score == 1 && x.Score == 3) {
        return 6
    }

    return 0
}

func Play(trans map[string]Throw, input []string) int {
    var score int

    for _, p := range input {
        if len(p) == 3 {
            throws := strings.Split(p, " ")
            other := trans[throws[0]]
            self := trans[throws[1]]

            score += self.Compare(other) + self.Score
        }
    }

    return score
}

func PlayRedux(trans map[string]Throw, input []string) int {
    var score int

    for _, p := range input {
        if len(p) == 3 {
            throws := strings.Split(p, " ")
            other := trans[throws[0]]
            var self Throw

            switch throws[1] {
                case "X":
                    if throws[0] == "A" {
                        self = trans["C"]
                    }
                    if throws[0] == "B" {
                        self = trans["A"]
                    }
                    if throws[0] == "C" {
                        self = trans["B"]
                    }
                case "Y": self = trans[throws[0]]
                case "Z":
                    if throws[0] == "A" {
                        self = trans["B"]
                    }
                    if throws[0] == "B" {
                        self = trans["C"]
                    }
                    if throws[0] == "C" {
                        self = trans["A"]
                    }
            }

            score += self.Compare(other) + self.Score
        }
    }

    return score
}
