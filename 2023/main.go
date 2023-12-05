package main

type Solution func()

var days = []Solution{day1,day2,day3}

func main() {
    days[2]()
}
