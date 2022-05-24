package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"tz-umico/domain"
)

//Entry point of program.
func main() {
	bug := Open(domain.BugPath)
	landscape := Open(domain.LandscapePath)
	res := BugCount(bug, landscape)
	// Printing result.
	fmt.Println(res)
}

// BugCount function counts number of bugs in the landscape.
func BugCount(bug [][]byte, landscape [][]byte) int {
	count := 0
	// Iterate through y-coordinates of the landscape.
	for i := range landscape {
		// Iterate through x-coordinates of the landscape.
		for j := range landscape[i] {
		loop:
			// Iterate through y-coordinates of the bug.
			for y := range bug {
				// Iterate through x-coordinates of the bug.
				for x := range bug[y] {
					// Checking to out of range of y-coordinates.
					if len(landscape) <= i+y {
						break loop
					}
					// Checking to out of range of x-coordinates.
					if len(landscape[i+y]) <= j+x {
						break loop
					}
					// Comparing bytes of landscape and bug.
					if landscape[i+y][j+x] == bug[y][x] {
						// Checking if we are in the last byte of the bug.
						if len(bug)-1 == y && len(bug[y])-1 == x {
							// If yes -> count +1 and continue iteration.
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

// Open function reads file and sterilize content of the file to [][]byte format.
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
