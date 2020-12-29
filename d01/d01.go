package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type input struct {
	s string
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	l := scanner.Text()
	i.s = l

	err = scanner.Err()
	panicOnError(err)

	return i
}

func part1(i input) int {
	up := strings.Count(i.s, "(")
	down := strings.Count(i.s, ")")
	return up - down
}

func part2(i input) int {
	current := 0
	step := 0
	for current != -1 {
		if i.s[0] == '(' {
			current++
		} else {
			current--
		}
		i.s = i.s[1:]
		step++
	}
	return step
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
