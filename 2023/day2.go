package main

import (
	"fmt"
    "strings"
    "strconv"
)

func d2part1() {
    //input := readFile("inputs/day2_ex.txt")
    input := readFile("inputs/day2.txt")

    t_red, t_green, t_blue := 12,13,14

    valid_games := 0

    for i,line := range input {
        game_info := strings.Split(line,":")
        games := strings.Split(game_info[1],";")
        game_valid_takes := 0
    

        for _,game := range games {
            balls := strings.Split(game,",")
            red, green, blue := 0,0,0

            for _,color := range balls {
                ball_info := strings.Split(color," ")
                
                n,_ := strconv.Atoi(ball_info[1])

                switch ball_info[2] {
                    case "red":
                        red += n
                    case "green":
                        green += n
                    case "blue":
                        blue += n
                }
            }

            // if take is valid incremet valid game counter
            if red <= t_red && green <= t_green && blue <= t_blue {
                game_valid_takes += 1
            }
        }

        if game_valid_takes == len(games) {
            valid_games += i+1
        } 
    }

    fmt.Printf("Part 1: %v\n",valid_games)
}


func d2part2() {
    //input := readFile("inputs/day2_ex.txt")
    input := readFile("inputs/day2.txt")

    total_power := 0

    for _,line := range input {
        game_info := strings.Split(line,":")
        games := strings.Split(game_info[1],";")
        max_red, max_green, max_blue := 0,0,0
        
        for _,game := range games {
            balls := strings.Split(game,",")

            for _,color := range balls {
                ball_info := strings.Split(color," ")
                
                n,_ := strconv.Atoi(ball_info[1])

                switch ball_info[2] {
                    case "red":
                        if n > max_red {
                            max_red = n
                        }
                    case "green":
                        if n > max_green {
                            max_green = n
                        }
                    case "blue":
                        if n > max_blue {
                            max_blue = n
                        }
                }
            }
        }

        total_power += max_red*max_green*max_blue
    }

    fmt.Printf("Part 2: %v\n",total_power)
}


func day2() {
    d2part1()
    d2part2()
}
