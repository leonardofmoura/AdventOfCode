package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) []string {
    var ret []string

    f, err := os.Open(filename)

    if err != nil {
        fmt.Println(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        ret = append(ret, scanner.Text())
    }

    if err:= scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return ret
}

