package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "sort"
    "strings"
    "strconv"
    
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

    var elfArr [][]int
    var out []int

    for _, num := range textArr {
        if num != "" {
            n, _ := strconv.Atoi(num)
            out = append(out, n)
        }

        if num == "" {
            elfArr = append(elfArr, out)
            out = []int{}
        }
    }

    sorted := sortHighToLow[int](caloriesEach(elfArr))
    fmt.Println(sorted[0])
    fmt.Println(sum[int](sorted[:3]))
}

func maxCalories(elves [][]int) int {
    return max[int](caloriesEach(elves))
}

func caloriesEach(elves [][]int) []int {
    var calories []int
    for _, elf := range elves {
        calories = append(calories, sum[int](elf))
    }
    return calories
}

type number interface {
    int|int8|int16|int32|int64|uint|uint8|uint16|uint32|uint64|float32|float64
}

func sum[T number](nums []T) T {
    var out T
    for _, num := range nums {
        out += num
    }
    return out
}

func max[T number](nums []T) T {
    var m T
    for _, num := range nums {
        if m < num {
            m = num
        }
    }
    return m
}

func sortHighToLow[T number](nums []T) []T {
    sort.SliceStable(nums, func (i, j int) bool { return nums[i] > nums[j] })
    return nums
}
