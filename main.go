package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type record struct {
	rawValue int
	time     int
	coord    int
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		panic("insufficient arguments; usage: `command size_of_grid input`")
	}

	filename := args[2]
	length, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}
	gridPerTime := length * length

	rawContent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	content := string(rawContent)

	contents := strings.Split(content, " ")
	inputSlice := []record{}
	for i, v := range contents {
		result, _ := strconv.Atoi(v)
		if result <= 0 {
			continue
		}

		time := i / gridPerTime
		coord := i % gridPerTime + 1
		inputSlice = append(inputSlice, record{
			rawValue: result,
			time: time,
			coord: coord,
		})
	}

	printGrid(length)

	fmt.Println("\nhistory:")
	for _, v := range inputSlice {
		fmt.Printf("%+v\n", v)
	}
}

func printGrid(length int) {
	current := 1
	fmt.Println("grid:")
	for inner := 0; inner < length; inner++ {
		for outer := 0; outer < length; outer++ {
			fmt.Printf("%3v ", current)
			current += 1
		}
		fmt.Println()
	}
}
