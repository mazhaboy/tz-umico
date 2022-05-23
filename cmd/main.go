package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"tz-umico/domain"
)

func main() {
	bug := Open(domain.BugPath)
	land := Open(domain.LandscapePath)
	res := BugCount(bug, land)
	fmt.Println(res)
}

func BugCount(bug [][]byte, land [][]byte) int {
	count := 0

	for i := range land {
		for j := range land[i] {
		loop:
			for y := range bug {
				for x := range bug[y] {
					if len(land) <= i+y {
						break loop
					}
					if len(land[i+y]) <= j+x {
						break loop
					}
					if land[i+y][j+x] == bug[y][x] {
						if len(bug)-1 == y && len(bug[y])-1 == x {
							count++
						}
					} else {
						break loop
					}
				}
			}
		}
	}

	return count
}

func Open(filePath string) [][]byte {
	arr := make([][]byte, 0)
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		arr = append(arr, []byte(scanner.Text()))
	}
	err = scanner.Err()

	if err != nil {
		log.Fatal()
	}

	return arr
}
