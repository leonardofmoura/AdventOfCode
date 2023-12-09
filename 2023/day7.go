package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func day7() {
    d7p1()
    d7p2()
}

func matchHand(high, second int) int {
    if high == 5 {
        return 0
    } else if high == 4 {
        return 1
    } else if high == 3 && second == 2 {
        return 2
    } else if high == 3 && second == 1 {
        return 3
    } else if high == 2 && second == 2 {
        return 4
    } else if high == 2 && second == 1 {
        return 5
    } else {
        return 6
    }
}


func loadHands(lines []string) map[string]int {
    game := map[string]int{}
    //store all hands in a map
    for _,line := range lines {
        l := strings.Split(line," ")
        n,_ := strconv.Atoi(l[1])
        game[l[0]] = n
    }
    return game
}


func extractHand(hand, ordering string) []int {
    hand_cards := make([]int, len(ordering))
    for _,card := range hand {
        index := strings.Index(ordering,string(card))
        hand_cards[index] += 1
    }
    return hand_cards
}


func orderHandTypes(hands *[][]string, ordering string) {
    // sort the arrays in each hand
    for _,cards := range *hands {
        slices.SortFunc(cards,func(a, b string) int {
            for i := 0; i < len(a); i++ {
                if a[i] == b[i] {
                    continue
                }
                return strings.Index(ordering,string(b[i])) - strings.Index(ordering,string(a[i])) 
            }
            return 0
        })
    } 
}


func calculateWins(hands [][]string, game map[string]int) int64 {
    var total_wins int64 = 0
    current_rank := len(game)
    for _,hand := range hands {
        if len(hand) > 0 {
            for _,cards := range hand {
                total_wins += int64(current_rank*game[cards])
                current_rank -= 1
            }
        }
    }
    return total_wins
}

func d7p1() {
    lines := readFile("inputs/day7.txt")
    game := loadHands(lines)


    // parse all hands into its type
    hands := make([][]string, 7)
    
    ordering := "23456789TJQKA"

    for hand := range game {
        hand_cards := extractHand(hand,ordering)

        // check highest card
        high, second := 0,0
        for _,card := range hand_cards {
            if card > high {
                second = high
                high = card
            } else if card > second {
                second = card
            }
        }
        //fmt.Printf("%v: -> %v (%v,%v)\n",hand,hand_cards,high,second)

        // match with the correct hand
        score := matchHand(high,second)
        hands[score] = append(hands[score], hand)
    }

    //order hands
    orderHandTypes(&hands,ordering)

    // calculate winnings
    total_wins := calculateWins(hands,game)
    
    fmt.Printf("Part 1: %v\n",total_wins)
}

func d7p2() {
    lines := readFile("inputs/day7.txt")
    game := loadHands(lines)


    // parse all hands into its type
    hands := make([][]string, 7)
    
    ordering := "J23456789TQKA"

    for hand := range game {
        hand_cards := extractHand(hand,ordering)

        // check highest card
        high, second := 0,0
        for i:=len(hand_cards)-1 ; i > 0; i-- {
            card := hand_cards[i]
            if card > high {
                second = high
                high = card
            } else if card > second {
                second = card
            }
        }
        // deal with jokers
        high += hand_cards[0]

        //fmt.Printf("%v: -> %v (%v,%v)\n",hand,hand_cards,high,second)

        // match with the correct hand
        score := matchHand(high,second)
        hands[score] = append(hands[score], hand)
    }

    //order hands
    orderHandTypes(&hands,ordering)

    // calculate winnings
    total_wins := calculateWins(hands,game)
    
    fmt.Printf("Part 2: %v\n",total_wins)
}
