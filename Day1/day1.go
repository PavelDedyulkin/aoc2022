package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	current := 0

	var top [4]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			top[3] = current
			slice := top[:]
			sort.Sort(sort.Reverse(sort.IntSlice(slice)))
			current = 0
		} else {
			res, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			current += res
		}
	}

	top[3] = current
	slice := top[:]
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	print(top[0] + top[1] + top[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
