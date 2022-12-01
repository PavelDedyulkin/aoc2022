package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func process(r io.Reader) (int, error) {

	scanner := bufio.NewScanner(r)
	current := 0

	var top [4]int

	for {

		eof := !scanner.Scan()

		currentLoaded := scanner.Text() == ""

		if currentLoaded {

			top[len(top)-1] = current
			slice := top[:]
			sort.Sort(sort.Reverse(sort.IntSlice(slice)))
			current = 0

			if eof {
				break
			}

			continue
		}

		cnt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		current += cnt

	}

	return top[0] + top[1] + top[2], scanner.Err()

}

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result, err := process(file)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}
