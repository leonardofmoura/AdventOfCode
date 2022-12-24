from lib import *

def solve():
    lines = readfile('day4')
    solve1(lines)

def solve1(lines):
    count1 = 0
    count2 = 0

    for l in lines:
        intervals = l.split(',')
        i1 = intervals[0].split('-')
        i2 = intervals[1].split('-')

        if ((int(i1[0]) <= int(i2[0]) and int(i1[1]) >= int(i2[1])) or (int(i2[0]) <= int(i1[0]) and int(i2[1]) >= int(i1[1]))):
            count1 += 1

        if (int(i2[0]) <= int(i1[1]) and int(i2[1]) >= int(i1[0])) or (int(i1[0]) <= int(i2[1] and int(i1[1]) >= int(i2[0]))):
            count2 += 1

    print(f"Solution 1: {count1}")
    print(f"Solution 2: {count2}")
