import sys 

import day1.sol
import day2.sol

days = [
    day1.sol,
    day2.sol
]

if (len(sys.argv) < 2):
    print("usage: aoc.py <day_number>")
else:
    days[int(sys.argv[1])-1].solve()
