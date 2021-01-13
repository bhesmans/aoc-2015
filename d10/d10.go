package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type input struct {
	number string
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	l := scanner.Text()
	i.number = l

	err = scanner.Err()
	panicOnError(err)

	return i
}

func _say(number *string) string {
	i, c := (*number)[0], 0

	for c < len(*number) && (*number)[c] == i {
		c++
	}

	*number = (*number)[c:]
	return fmt.Sprintf("%v%c", c, i)
}

func say(number *string) string {
	// Note: need string builder here, otherwise it's slow as f***
	var ret strings.Builder
	for len(*number) != 0 {
		ret.WriteString(_say(number))
	}

	return ret.String()
}

func part1(in input, iter int) int {
	for i := 0; i < iter; i++ {
		in.number = say(&in.number)
	}
	return len(in.number)
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i, 40))
	fmt.Printf("%v\n", part1(i, 50))
}
