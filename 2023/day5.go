package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func day5() {
    d5p1()
    d5p2()
}

func parseSeeds(seeds string) []string {
    s := strings.Split(seeds," ")
    return s[1:]
}

type Interval struct {
    source int64
    dest int64
    size int64
}

type Table []Interval

func parseInterval(line string) Interval {
    s := strings.Split(line," ")
    source,_ := strconv.Atoi(s[1])
    dest,_ := strconv.Atoi(s[0])
    size,_ := strconv.Atoi(s[2])
    return Interval{int64(source),int64(dest),int64(size)}
}

func parseTables(lines []string) []Table {
    tables := []Table{}
    parseNewTable := false
    currTable := -1

    for i := 1; i < len(lines); i++ {
        if lines[i] == "" {
            parseNewTable = true
            i+=1
            currTable += 1
            continue
        } 

        if parseNewTable {
            tables = append(tables, Table{parseInterval(lines[i])}) 
            //fmt.Printf("New table: %v\n",lines[i])
            parseNewTable = false
        } else {
            tables[currTable] = append(tables[currTable], parseInterval(lines[i]))
        }
    }

    return tables
}

func lookupSeed(tables []Table, seed int64) int64 {
    curr_value := seed
    //fmt.Printf("Seed %v -> ",curr_value)
    for _,table := range tables {
        // if loop exits sucessfully, keep same value
        for _,interval := range table {
            if curr_value >= interval.source && curr_value <=interval.source + interval.size - 1 {
                curr_value = curr_value+interval.dest-interval.source
                break
            }
        }
        //fmt.Printf("%v -> ",curr_value)
    }

    return curr_value
}

func d5p1() {
    lines := readFile("inputs/day5.txt")
    seeds := parseSeeds(lines[0])
    tables := parseTables(lines)

    locations := []int64{}

    // calculate seed location
    for _,seed := range seeds {
        curr_value,_ := strconv.Atoi(seed)
        res := lookupSeed(tables,int64(curr_value))
        locations = append(locations, res)
    }

    fmt.Printf("Part 1: %v\n",slices.Min(locations))
}

type RangeInterval struct {
    start int64
    end int64
}

func parseSeedIntervals(line string) []RangeInterval {
    s := strings.Split(line," ")
    intervals := []RangeInterval{}

    for i:=1; i < len(s); i+=2 {
        start,_ := strconv.Atoi(s[i])
        size,_ := strconv.Atoi(s[i+1])
        intervals = append(intervals, RangeInterval{int64(start),int64(start+size-1)})
    }

    return intervals
}

type DomainMap []RangeInterval

func intervalIntersect(y RangeInterval, x Interval) ([2]RangeInterval, bool) {
    xEnd := x.dest + x.size-1
    // x overlaps y
    if x.dest <= y.start && xEnd >= y.end {
        intersectx := RangeInterval{x.source+y.start-x.dest,x.source+y.end-x.dest}
        intersecty := RangeInterval{y.start,y.end}
        //fmt.Println("outside")
        return [2]RangeInterval{intersectx,intersecty}, true
    } else if x.dest <= y.start && xEnd <= y.end && xEnd > y.start{
        intersectx := RangeInterval{x.source+y.start-x.dest,x.source+x.size-1}
        intersecty := RangeInterval{y.start,xEnd}
        //fmt.Println("left")
        return [2]RangeInterval{intersectx,intersecty}, true
    } else if x.dest >= y.start && xEnd >= y.end && x.dest < y.end{
        intersectx := RangeInterval{x.source,x.source+y.end-x.dest}
        intersecty := RangeInterval{x.dest,y.end}
        //fmt.Println("right")
        return [2]RangeInterval{intersectx,intersecty}, true
    } else if x.dest >= y.start && xEnd <= y.end {
        intersectx := RangeInterval{x.source,x.source+x.size-1}
        intersecty := RangeInterval{x.dest,xEnd}
        //fmt.Println("inside")
        return [2]RangeInterval{intersectx,intersecty}, true
    }
    //fmt.Println("false")
    return [2]RangeInterval{}, false
}

func searchDomain(layer int, tables []Table, seeds []RangeInterval, currentSearch RangeInterval) []RangeInterval {
    //serach seeds
    if layer == -1 {
        var ret []RangeInterval = nil
        for _,seed := range seeds {
            //fmt.Printf("Testing interval %v against seed %v\n",currentSearch,seed)
            if seed.start <= currentSearch.start && seed.end >= currentSearch.end {
                //fmt.Println("inside")
                ret = []RangeInterval{}
                ret = append(ret, RangeInterval{currentSearch.start,currentSearch.end})
            } else if seed.start >= currentSearch.start && seed.end <= currentSearch.end {
                //fmt.Println("outside")
                ret = []RangeInterval{}
                ret = append(ret, RangeInterval{seed.start,seed.end})
            } else if seed.start >= currentSearch.start && seed.end >= currentSearch.end && seed.start < currentSearch.end{
                //fmt.Println("left")
                ret = []RangeInterval{}
                ret = append(ret, RangeInterval{seed.start,currentSearch.end})
            } else if seed.start >= currentSearch.start && seed.end <= currentSearch.end && seed.end > currentSearch.start{
                //fmt.Println("right")
                ret = []RangeInterval{}
                ret = append(ret, RangeInterval{currentSearch.start,seed.end})
            }
        }
        return ret
    }
 
    searchTable := tables[layer]
    slices.SortFunc(searchTable, func(a, b Interval) int {
        return int(a.source-b.source)
    })

    //complete table with missing intervals
    if searchTable[0].source != 0 {
        searchTable = append(searchTable, Interval{0,0,searchTable[0].source})
    }
    last := searchTable[len(searchTable)-1].source+searchTable[len(searchTable)-1].size+1
    searchTable = append(searchTable, Interval{last,last,math.MaxInt32})

    valid_intersects := []RangeInterval{}

    for _,interval := range searchTable {
        //fmt.Printf("Intersecting %v -> %v\n",currentSearch,interval)
        intersection, intersects := intervalIntersect(currentSearch,interval)
        if intersects {
            //fmt.Printf("Found Intersection %v -> %v\n",intersection[1],intersection[0])
            valid_intersects = append(valid_intersects, intersection[0])
        }
    }

    slices.SortFunc(valid_intersects, func(a, b RangeInterval) int {
        return int(a.start-b.start)
    })
    
    solutions := []RangeInterval{}

    for _,intersection := range valid_intersects {
        seeds := searchDomain(layer-1,tables,seeds,intersection)
        if seeds != nil {
           solutions = append(solutions, seeds...) 
        }
    } 
    
    return solutions
}
   

func d5p2() {
    lines := readFile("inputs/day5.txt")
    seedIntervals := parseSeedIntervals(lines[0])
    tables := parseTables(lines)

    lastTable := tables[len(tables)-1]
    slices.SortFunc(lastTable, func(a, b Interval) int {
        return int(a.dest-b.dest)
    })

    //convert intervals to better format
    startingTable := []RangeInterval{}
    for _,interval := range lastTable {
        startingTable = append(startingTable, RangeInterval{interval.dest,interval.dest+interval.size-1})
    }

    //complete table with missing intervals
    if startingTable[0].start != 0 {
        startingTable = append(startingTable, RangeInterval{0,startingTable[0].start-1})
    }
    //startingTable = append(startingTable, RangeInterval{startingTable[len(startingTable)-2].end+1,math.MaxInt})

    // sort again
    slices.SortFunc(startingTable, func(a, b RangeInterval) int {
        return int(a.start-b.start)
    })

    //fmt.Printf("Seeds %v\n",seedIntervals)
    //fmt.Printf("Search intervals: %v\n",startingTable)

    solutions := []int64{}

    for _,interval := range startingTable {
        seeds := searchDomain(len(tables)-1,tables,seedIntervals,interval)
        if seeds != nil {
            for _,seed := range seeds {
                solutions = append(solutions, lookupSeed(tables,seed.start))
            }
        }
    }

    fmt.Printf("Part 2: %v\n",slices.Min(solutions))
}
