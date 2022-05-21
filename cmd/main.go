package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	bagPath       = "resources/bug/bug.txt"
	landscapePath = "resources/test-keys/landscape.txt"
)

func main() {
	bug, bugColumns, bugRows := Open(bagPath)
	landscape, landscapeColumns, landscapeRows := Open(landscapePath)
	fmt.Println(bug, bugColumns, bugRows)
	fmt.Println(landscape, landscapeColumns, landscapeRows)
}

func Open(filePath string) ([][]byte, int, int) {
	bug := make([][]byte, 0)
	columns := 0
	rows := 0
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if columns < len(scanner.Text()) {
			columns = len(scanner.Text())
		}
		bug = append(bug, []byte(scanner.Text()))
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal()
	}
	rows = len(bug)
	return bug, rows, columns
}
