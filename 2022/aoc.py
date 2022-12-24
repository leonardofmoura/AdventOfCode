import sys 

import day1.sol
import day2.sol
import day3.sol
import day4.sol
import day5.sol

days = [
    day1.sol,
    day2.sol,
    day3.sol,
    day4.sol,
    day5.sol,
]

if (len(sys.argv) < 2):
    print("usage: aoc.py <day_number>")
else:
    days[int(sys.argv[1])-1].solve()
