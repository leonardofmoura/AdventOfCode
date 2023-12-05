package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

func day4() {
    d4p1()
    d4p2()
}


func buildKeys(keys string) []string {
    ret := []string{}

    for i:=0; i <= len(keys)-2; i+=3 {
        ret = append(ret, keys[i:i+2]) 
    }

    return ret
}

func d4p1() {
    //cards := readFile("inputs/day4_ex.txt")
    cards := readFile("inputs/day4.txt")

    totalScore := 0

    for _,line := range cards {
        card := strings.Split(line,": ")
        nums := strings.Split(card[1]," | ")
        
        key := buildKeys(nums[0])
        draws := buildKeys(nums[1])
        winCount := 0

        for _,item := range key {
            //fmt.Printf("Testing %v\n",item)
            if slices.Contains(draws,item) {
                //fmt.Println("Match!")
                winCount += 1
            }
        }
        
        score := int(math.Pow(2,float64(winCount-1)))
        //fmt.Printf("Card %v Score: %v -> %v points\n",i+1,winCount,score)
         
        totalScore += score 
    }

    fmt.Printf("Part 1: %v\n",totalScore)
}

type Card struct {
    id int
    keys []string
    draws []string
    wins int
}

func parseCard(card string, id int) Card {
    nums := strings.Split(card,": ")
    keys := strings.Split(nums[1]," | ")
    
    key := buildKeys(keys[0])
    draws := buildKeys(keys[1])

    return Card{id,key,draws,-1}
}

func d4p2() {
    //cards := readFile("inputs/day4_ex.txt")
    cards := readFile("inputs/day4.txt")

    evalCards := []Card{}

    // parse all cards 
    for i,line := range cards {
        evalCards = append(evalCards, parseCard(line,i))
    } 

    for i := 0; i < len(evalCards); i++ {
        card := &evalCards[i]

        if card.wins == -1 {
            winCount := 0

            for _,item := range card.keys {
                //fmt.Printf("Testing %v\n",item)
                if slices.Contains(card.draws,item) {
                    //fmt.Println("Match!")
                    winCount += 1
                }
            }

            card.wins = winCount
        }


        if card.wins > 0 {
            next := card.id+1
            if next >= len(cards) {
                next = len(cards)
            }

            lenAppend := card.id+card.wins+1
            if lenAppend >= len(cards) {
                lenAppend = len(cards)
            }
            
            evalCards = append(evalCards, evalCards[next:lenAppend]...)
            //fmt.Printf("Card %v: adding cards [%v:%v]\n",card.id,next,lenAppend)

        }
    }


    fmt.Printf("Part 2: %v\n",len(evalCards))
}
