from lib import *
import copy

def solve():
    lines = readfile('day5')
    solve1(lines)

def solve1(lines):
    nstacks = 0
    moveStart = 0

    # find number of stacks
    for i in range(len(lines)):
        if lines[i].find(' 1 ') == 0:
            nstacks = lines[i][-2]
            moveStart = i + 2
            break

    stacks = [[] for _ in range(int(nstacks))]

    # load the stacks 
    for l in lines:
        if (l[1] == '1'):
            break
        
        col = 0
        for i in range(1,len(l),4):
            if l[i] != ' ':
                stacks[col].append(l[i])
            col += 1

    # reverse stacks
    for s in stacks:
        s.reverse()

    stacks2 = copy.deepcopy(stacks)

    # apply movements for first crane
    for i in range(moveStart,len(lines)):
        move = [int(i) for i in lines[i].split() if i.isdigit()]
        n = move[0]
        src = move[1] - 1
        dest = move[2] - 1

        for _ in range(n):
            stacks[dest].append(stacks[src][-1])
            stacks[src].pop()

    res = ''

    # generate string
    for s in stacks:
        res += s[-1]

    print(f"Solution 1: {res}")

    # apply movements for second crane
    for i in range(moveStart, len(lines)):
        move = [int(i) for i in lines[i].split() if i.isdigit()]
        n = move[0]
        src = move[1] - 1
        dest = move[2] - 1

        stacks2[dest] += stacks2[src][-n:]
        del stacks2[src][-n:]
    

    res = ''

    # generate string
    for s in stacks2:
        res += s[-1]

    print(f"Solution 2: {res}")
