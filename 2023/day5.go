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
        if lines[i] == "" {
            parseNewTable = true
            i+=1
            currTable += 1
            continue
        } 

        if parseNewTable {
            tables = append(tables, Table{parseInterval(lines[i])}) 
            fmt.Printf("New tale: %v\n",lines[i])
            parseNewTable = false
        } else {
            tables[currTable] = append(tables[currTable], parseInterval(lines[i]))
        }
    }

    locations := []string{}

    // calculate seed location
    for _,seed := range seeds {
        curr_value,_ := strconv.Atoi(seed)
        for _,table := range tables {
            for _,interval := range table {
                if curr_value >= interval.source && curr_value <=interval.source + interval.size - 1 {
                    curr_value = curr_value-interval.dest+interval.size-1
                    break
                }
            }
        } 
    }

    fmt.Printf("Seeds: %v\n",seeds)
    fmt.Printf("Tables %v",tables)
}
