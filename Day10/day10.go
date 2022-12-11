package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/golang-collections/collections/queue"
)

func process2(r io.Reader) (int, error) {

	scanner := bufio.NewScanner(r)

	var offsets queue.Queue

	for scanner.Scan() {

		line := scanner.Text()

		var command string

		fmt.Sscanf(line, "%s", &command)

		offsets.Enqueue(0)

		if command == "addx" {

			var registerOffset int
			fmt.Sscanf(line, "addx %d", &registerOffset)

			offsets.Enqueue(registerOffset)
		}

	}

	pixelf := func(sprite int, i int) byte {

		if math.Abs((float64)(sprite-i)) < 2 {
			return '#'
		}
		return '.'
	}
	x := 1
	position := 0
	var line [40]byte

	for offsets.Peek() != nil {

		line[position] = pixelf(x, position)
		x += offsets.Dequeue().(int)

		position++

		if position == 40 {

			fmt.Println(string(line[:]))
			position = 0
		}
	}

	return 0, nil

}

func process1(r io.Reader) (int, error) {

	scanner := bufio.NewScanner(r)

	var offsets queue.Queue

	for scanner.Scan() {

		line := scanner.Text()

		var command string

		fmt.Sscanf(line, "%s", &command)

		offsets.Enqueue(0)

		if command == "addx" {

			var registerOffset int
			fmt.Sscanf(line, "addx %d", &registerOffset)

			offsets.Enqueue(registerOffset)
		}

	}

	x := 1
	signal_i := 0
	sum := 0
	cycle := 1

	for offsets.Peek() != nil {

		if cycle == 40*signal_i+20 {

			sum += cycle * x
			signal_i++
		}
		x += offsets.Dequeue().(int)
		cycle++

	}

	return sum, nil

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result, err := process2(file)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}
