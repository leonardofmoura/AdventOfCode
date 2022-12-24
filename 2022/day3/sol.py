from lib import *

def letterValue(letter):
    if (ord(letter) <= 90): # uppercase
        return ord(letter) - ord('A') + 26
    else:
        return ord(letter) - ord('a')


def solve():
    lines = readfile('day3')
    solve1(lines)
    solve2(lines)

def solve1(lines):
    sum = 0

    for l in lines:
        map1 = [0] * 52
        map2 = [0] * 52

        half = int(len(l)/2)
        c1 = l[:half]
        c2 = l[half:]

        for c in c1:
            map1[letterValue(c)] += 1

        for c in c2:
            map2[letterValue(c)] += 1

        for i in range(52):
            if (map1[i] > 0 and map2[i] > 0):
                sum += i+1

    print(f"Solution 1: {sum}")


def solve2(lines):
    map = []

    sum = 0

    for l in lines:
        lineMap = [0] * 52

        for c in l:
            lineMap[letterValue(c)] += 1
        
        map.append(lineMap)

    for letter in range(52):
        for i in range(0,len(map),3):
            if (map[i][letter] > 0 and map[i+1][letter] > 0 and map[i+2][letter] > 0):
                sum += letter + 1

    print(f"Solution 2: {sum}")


