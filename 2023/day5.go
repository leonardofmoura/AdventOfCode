package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day5() {
    d5p1()
}

func parseSeeds(seeds string) []string {
    s := strings.Split(seeds," ")
    return s[1:]
}

type Interval struct {
    source int
    dest int
    size int
}

type Table []Interval

func parseInterval(line string) Interval {
    s := strings.Split(line," ")
    source,_ := strconv.Atoi(s[1])
    dest,_ := strconv.Atoi(s[0])
    size,_ := strconv.Atoi(s[2])
    return Interval{source,dest,size}
}

func d5p1() {
    lines := readFile("inputs/day5_ex.txt")
    seeds := parseSeeds(lines[0])

    tables := []Table{}

    parseNewTable := false
    currTable := -1
    for i := 1; i < len(lines); i++ {
        if lines[i] == "\n" {
            parseNewTable = true
            i+=2
            currTable += 1
            continue
        } 

        if parseNewTable {
            tables = append(tables, Table{parseInterval(lines[i])}) 
        } else {
            tables[currTable] = append(tables[currTable], parseInterval(lines[i]))
        }
    }

    fmt.Printf("Seeds: %v\n",seeds)
    fmt.Printf("Tables %v",len(tables))
}
