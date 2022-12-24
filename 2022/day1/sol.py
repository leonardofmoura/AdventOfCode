from lib import *

def solve():
    lines = readfile('day1')

    gt = []
    s = 0

    for l in lines:
        if (l == '\n'):
            gt.append(s) 

            s = 0
            continue

        s += int(l[:-1])

        gt.sort(reverse=True)

    print(f"Solution 1: {gt[0]}")
    print(f"Solution 2: {gt[0] +  gt[1] + gt[2]}")
