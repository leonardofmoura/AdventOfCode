package main

import (
    "fmt"
    "regexp"
    "strings"
    "slices"
    "unicode"
)

func d1part1() {
    lines := readFile("inputs/day1.txt")

    numbers := regexp.MustCompile(`\D`)

    // iterate all lines
    for i,line := range lines {
        lines[i] = numbers.ReplaceAllString(line,"")
    }

    sum := 0 

    for _,l := range lines {
        if len(l) > 0 {
            sum += int(l[0] - '0') * 10
            sum += int(l[len(l)-1] - '0')
        }
    }

    fmt.Printf("Part1: %v\n",sum)
}

func d1part2() {
    nums := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    lines := readFile("inputs/day1.txt")

    sum := 0

    // iterate all lines
    for _,l := range lines {
        var values []int
        var indexes []int

        // test every keyword
        for key := range nums {
            j := 0

            // find all instances!
            for {
                index := strings.Index(l[j:],key)

                if index == -1 {
                    break
                }
                
               indexes = append(indexes, index+j) 
               values = append(values, nums[key])
               //fmt.Printf("Found %v in line %v\n", key, i)
               j += index+1
            }
        }

        //also find every digit 
        for i,c := range l {
            if unicode.IsDigit(rune(c)) {
                indexes = append(indexes, i)
                values = append(values, int(c)-'0')
            }
        }

        // min index is first digit
        // the digit value is the corresponding elemnt in values
        m := values[slices.Index(indexes, slices.Min(indexes))] * 10
        //fmt.Printf("First: %v\n",values[slices.Index(indexes, slices.Min(indexes))])

        // do the inverse for last digit
        n := values[slices.Index(indexes, slices.Max(indexes))] 
        //fmt.Printf("Second: %v\n",values[slices.Index(indexes, slices.Max(indexes))])

        //fmt.Printf("line %v: %v\n",i,m+n)

        sum += m+n
    }

    fmt.Printf("Part 2: %v\n",sum)
}

func day1() {
    d1part1()
    d1part2()
}
