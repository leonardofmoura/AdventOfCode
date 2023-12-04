package main

type Solution func()

var days = []Solution{day1,day2}

func main() {
    days[1]()
}
