package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func day3() {
    d3p1()
    d3p2()
}

func getNumberLen(schem []string, line, startCol int) int {
    i := startCol
    for ; i < len(schem[line]); i++ {
        if !unicode.IsDigit(rune(schem[line][i])) {
            return i-startCol
        }
    }
    return i-startCol
}

func indentifyPartNumber(schem []string, line,startCol,nLen int) bool {
    lineStart, colStart, lineEnd, colEnd := line-1,startCol-1,line+1,startCol+nLen

    // check schematic boundaries
    if lineStart < 0 {
        lineStart = 0
    }
    if colStart < 0 {
        colStart = 0
    }
    if lineEnd >= len(schem) {
        lineEnd = len(schem)-1
    }
    if colEnd >= len(schem[0]) {
        colEnd = len(schem[0])-1
    }

    // can be optimized by skiping the number itself
    for i := lineStart; i <= lineEnd; i++ {
        for j := colStart; j <= colEnd; j++ {
            c := schem[i][j]
            //fmt.Printf("%c ",rune(c))
            if !unicode.IsDigit(rune(c)) && c != '.' {
                //fmt.Printf("\nPartNumber!\n")
                return true
            }
        }
    }

    //fmt.Printf("\nNotPartNumber\n")
    
    return false
}


func d3p1() {
    //schem := readFile("inputs/day3_ex.txt")
    schem := readFile("inputs/day3.txt")

    partNumberSum := 0

    for i,line := range schem {
        for j := 0; j < len(line); j++ {
            
            // identify numbers
            if unicode.IsDigit(rune(line[j])) {
                nLen:= getNumberLen(schem,i,j)
                //fmt.Printf("Found number %v\n",line[j:j+nLen]) 

                // test surrounding area
                if indentifyPartNumber(schem,i,j,nLen) {
                    num,_ := strconv.Atoi(line[j:j+nLen])
                    partNumberSum += num
                }

                j += nLen
            }
        }
    }

    fmt.Printf("Part 1: %v\n",partNumberSum)
}


func indentifyGearPartNumber(schem []string, line,startCol,nLen int) ([2]int, bool) {
    lineStart, colStart, lineEnd, colEnd := line-1,startCol-1,line+1,startCol+nLen

    // check schematic boundaries
    if lineStart < 0 {
        lineStart = 0
    }
    if colStart < 0 {
        colStart = 0
    }
    if lineEnd >= len(schem) {
        lineEnd = len(schem)-1
    }
    if colEnd >= len(schem[0]) {
        colEnd = len(schem[0])-1
    }

    // can be optimized by skiping the number itself
    for i := lineStart; i <= lineEnd; i++ {
        for j := colStart; j <= colEnd; j++ {
            c := schem[i][j]
            if c == '*' {
                //fmt.Printf("Found Gear: %v\n",[2]int{i,j})
                return [2]int{i,j}, true
            }
        }
    }
    return [2]int{0,0}, false
}


func d3p2() {
    schem := readFile("inputs/day3.txt")

    gearMap := make(map[[2]int][]int)

    for i,line := range schem {
        for j := 0; j < len(line); j++ {
            
            // identify numbers
            if unicode.IsDigit(rune(line[j])) {
                nLen:= getNumberLen(schem,i,j)
                //fmt.Printf("Found number %v\n",line[j:j+nLen])
                num,_ := strconv.Atoi(line[j:j+nLen])

                // populate map with all adgecent gears
                coords, gear := indentifyGearPartNumber(schem,i,j,nLen)
                if gear {
                    gearCoordList, ok := gearMap[coords]
                    if ok {
                        gearMap[coords] = append(gearCoordList, num)
                    } else {
                        gearMap[coords] = []int{num}
                    }
                }

                j += nLen
            }
        }
    }

    gearRatioSum := 0

    for _,partNumbers := range gearMap {
        //fmt.Printf("%v\n",partNumbers)
        if len(partNumbers) == 2 {
            gearRatioSum += partNumbers[0] * partNumbers[1]
        }
    }

    fmt.Printf("Part 2: %v\n",gearRatioSum)
}
