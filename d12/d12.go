package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type input struct {
	s string
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
	i.s = l

	err = scanner.Err()
	panicOnError(err)

	return i
}

func readString(s string, from *int) int {
	*from += 1
	for s[*from] != '"' {
		*from++
	}
	*from += 1
	return 0
}

func readInt(s string, from *int) int {
	v := 0
	mul := 1
	if s[*from] == '-' {
		mul *= -1
		*from++
	}

	for s[*from] >= '0' && s[*from] <= '9' {
		c := s[*from]
		v *= 10
		v += int(c - '0')
		*from++
	}

	return (v * mul)
}

func readList(s string, from *int) int {
	sum := 0
	for c := s[*from]; c != ']'; c = s[*from] {
		*from += 1
		sum += readItem(s, from)
	}
	*from += 1
	return sum
}

func checkString(s string, from int, check string) bool {
	i := 0
	for i = 0; i < len(check) && from < len(s); i, from = i+1, from+1 {
		if s[from] != check[i] {
			return false
		}
	}
	return i == len(check)
}

func readDict(s string, from *int) int {
	sum := 0
	mult := 1
	for c := s[*from]; c != '}'; c = s[*from] {
		*from += 1
		readString(s, from)
		*from += 1 // :
		if s[*from] == '"' && checkString(s, *from+1, "red") {
			mult = 0
		}
		sum += readItem(s, from)
	}
	*from += 1
	return sum * mult
}

func readItem(s string, from *int) int {
	c := s[*from]
	if c == '{' {
		return readDict(s, from)
	} else if c == '[' {
		return readList(s, from)
	} else if c == '"' {
		return readString(s, from)
	} else {
		return readInt(s, from)
	}
}

func (i input) sum() int {
	sum := 0
	re := regexp.MustCompile(`-?[0-9]*`)
	for _, sval := range re.FindAllString(i.s, -1) {
		sum += s2i(sval)
	}
	return sum

}

func part1(in input) int {
	return in.sum()
}

func part2(in input) int {
	from := 0
	return readItem(in.s, &from)
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
