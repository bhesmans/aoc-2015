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
	ss []string
}

func dataLen(s string) int {
	ret := len(s)

	ret -= 2

	ret -= strings.Count(s, `\"`)
	s = strings.ReplaceAll(s, `\"`, "")
	ret -= strings.Count(s, `\\`)
	s = strings.ReplaceAll(s, `\\`, "")
	ret -= strings.Count(s, `\x`) * 3 // assumes it's well formed

	return ret
}

func encodeLen(s string) int {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)

	return len(s) + 2
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		i.ss = append(i.ss, l)
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func part1(i input) int {
	sum := 0
	for _, s := range i.ss {
		sum += len(s) - dataLen(s)
	}
	return sum
}

func part2(i input) int {
	sum := 0
	for _, s := range i.ss {
		sum += encodeLen(s) - len(s)
	}
	return sum
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
