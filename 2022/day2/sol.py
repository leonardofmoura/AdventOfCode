from lib import *

map = {
    'A' : 0,
    'B' : 1,
    'C' : 2,
    'X' : 0,
    'Y' : 1,
    'Z' : 2,
}


def solve():
    solve1()
    solve2()

def solve1():
    lines = readfile('day2')

    score = 0

    for l in lines:
        game = l.split(' ')
        if (map[game[1]] == (map[game[0]] + 1) % 3): # win
            score += 6
        
        elif (map[game[0]] == map[game[1]]): # draw
            score += 3    

        score += map[game[1]] + 1 # add play

    print(f"Solution 1: {score}")


def solve2():
    lines = readfile('day2')

    score = 0

    for l in lines:
        game = l.split()

        if (l[1] == 'Y'): # draw 
           score += map[game[0]] + 1 + 3 

        elif (l[1] == 'Z'): # win 
            score += (map[game[0]] + 1) % 3 + 1 + 6

        else: # lose
            score += (map[game[0]] - 1) % 3 + 1

    print(f"Solution 2: {score}")
