package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readFile(filename string) []string {
    var ret []string

    f, err := os.Open(filename)

    if err != nil {
        fmt.Println(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        ret = append(ret, scanner.Text())
    }

    if err:= scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return ret
}

 	
func readInts(filename string) [][]int {
    lines := readFile(filename)
    find_digits := regexp.MustCompile(`\d+`)
    
    ret := [][]int{}

    for _,line := range lines {
        nums := find_digits.FindAllString(line,-1)
        num_ints := []int{}
        for _,s := range nums {
            n,_ := strconv.Atoi(s)
            num_ints = append(num_ints, n)
        }
        ret = append(ret, num_ints)
    }
    
    return ret
}
