package main

import (
	"fmt"
	"slices"
	"strings"
)

func day9() {
    d9p1()
    d9p2()
}

func testZero(arr []int) bool {
    for _,n := range arr {
        if n != 0 {
            return false
        }
    }
    return true
}

func visualize(levels [][]int) {
    for i,line := range levels {
        spaces := strings.Repeat(" ",i)
        fmt.Print(spaces)
        for _,c := range line {
            fmt.Printf("%v ",c)
        }
        fmt.Println()
    }
}

func predictNext(line []int) int {
    levels := [][]int{line}

    i := 0
    for ; !testZero(levels[i]); {
        next_level := []int{}
        for j := 0; j < len(levels[i])-1; j++ {
            next_level = append(next_level,  levels[i][j+1] - levels[i][j])
        }
        levels = append(levels, next_level)
        i++
    }

    //visualize(levels)

    for i := len(levels)-2; i >= 0 ; i-- {
        next := levels[i][len(levels[i])-1] + levels[i+1][len(levels[i+1])-1]
        levels[i] = append(levels[i], next)
    }

    return levels[0][len(levels[0])-1]
}

func d9p1() {
    lines := readInts("inputs/day9.txt") 

    results := []int{}
    for _,line := range lines {
        results = append(results, predictNext(line))
    }

    fmt.Printf("Part 1: %v\n",calcSum(results))
}


func d9p2() {
    lines := readInts("inputs/day9.txt") 

    results := []int{}
    for _,line := range lines {
        slices.Reverse(line)
        results = append(results, predictNext(line))
    }

    fmt.Printf("Part 2: %v\n",calcSum(results))
}
