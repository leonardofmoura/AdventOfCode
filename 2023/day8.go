package main

import (
	"fmt"
	"regexp"
	"strings"
    "math"
)


func day8() {
    d8p1()
    d8p2()
}

// from: https://github.com/TheAlgorithms/Go/blob/master/math/gcd/gcditerative.go
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// from: https://github.com/TheAlgorithms/Go/blob/master/math/lcm/lcm.go
func lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func d8p1() {
    lines := readFile("inputs/day8.txt")

    maps := make(map[string][]string)
    re := regexp.MustCompile(`\(|\)`)

    // parse the inputs
    for i := 2; i < len(lines); i++ {
        clean_line := re.ReplaceAllString(lines[i],"")
        src := strings.Split(clean_line," = ")
        dest := strings.Split(src[1],", ")
        maps[src[0]] = dest
    }

    path := lines[0]
    i := 0
    current_node := "AAA"
    steps := 0

    for ; current_node != "ZZZ"; {
        if path[i] == 'R' {
            current_node = maps[current_node][1]
        } else if path[i] == 'L' {
            current_node = maps[current_node][0]
        }
        i = (i + 1) %  len(path) 
        steps += 1
    }

    fmt.Printf("Part 1: %v\n",steps)
}


func d8p2() {
    lines := readFile("inputs/day8.txt")

    maps := make(map[string][]string)
    re := regexp.MustCompile(`\(|\)`)

    // parse the inputs
    for i := 2; i < len(lines); i++ {
        clean_line := re.ReplaceAllString(lines[i],"")
        src := strings.Split(clean_line," = ")
        dest := strings.Split(src[1],", ")
        maps[src[0]] = dest
    }


    // discover staring nodes
    starting_nodes := []string{}
    for key := range maps {
        if key[2] == 'A' {
            starting_nodes = append(starting_nodes, key)
        }
    }

    path := lines[0]
    loop_len := []int64{}
    for _,node := range starting_nodes {
        current_node := node

        i := 0
        var steps int64 = 0

        for true {
            if current_node[2] == 'Z' {
                loop_len = append(loop_len, steps)
                break
            }

            if path[i] == 'R' {
                current_node = maps[current_node][1]
            } else if path[i] == 'L' {
                current_node = maps[current_node][0]
            }

            i = (i + 1) %  len(path) 
            steps += 1
        }
    }

    fmt.Println(loop_len)

    lastlcm := int64(loop_len[0])
    for i := 1; i < len(loop_len); i++ {
        lastlcm = lcm(int64(lastlcm),int64(loop_len[i]))
    }
        
    fmt.Printf("Part 2: %v\n",lastlcm)
}
