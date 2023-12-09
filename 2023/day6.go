package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day6() {
    d6p1()
    d6p2()
}

func d6p1() {
    data := readInts("inputs/day6.txt")

    wins := 1

    for i,maxTime := range data[0] {
        // brute force possible distances for this race
        round_wins := 0
        for hold := 1; hold < maxTime; hold++ {
            travel_distance := hold*(maxTime-hold) 
            if travel_distance > data[1][i] {
                round_wins += 1
            }
        }
        wins *= round_wins
    }

    fmt.Printf("Part 1: %v\n",wins)
}


func d6p2() {
    data := readFile("inputs/day6.txt")
    re := regexp.MustCompile(`\d+`)
    times := strings.Join(re.FindAllString(data[0],-1)[:],"")
    distances := strings.Join(re.FindAllString(data[1],-1),"")

    time,_ := strconv.Atoi(times)
    dist,_ := strconv.Atoi(distances)


    round_wins := 0
    for hold := 1; hold < time; hold++ {
        travel_distance := hold*(time-hold) 
        if travel_distance > dist {
            round_wins += 1
        }
    }

    fmt.Printf("Part 2: %v\n",round_wins)
}
