package main

type Solution func()

var days = []Solution{day1,day2,day3,day4,day5,day6,day7}

func main() {
    days[6]()
}
