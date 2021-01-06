package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type str string

type input struct {
	s []str
}

func (s str) rule1() bool {
	re := regexp.MustCompile(`[aeiou]{1}`)
	t := re.Split(string(s), 4)
	return len(t) == 4
}

func (s str) rule2() bool {
	ss := string(s)
	for i := 1; i < len(ss); i++ {
		if ss[i-1] == ss[i] {
			return true
		}
	}
	return false
}

func (s str) rule3() bool {
	re := regexp.MustCompile(`ab|cd|pq|xy`)
	return !re.MatchString(string(s))
}

func (s str) rule4() bool {
	ss := string(s)
	for i := 0; i < len(ss)-2; i++ {
		sub := string(ss[i : i+2])
		if strings.Index(ss[i+2:], sub) != -1 {
			return true
		}
	}
	return false
}

func (s str) rule5() bool {
	ss := string(s)
	for i := 2; i < len(ss); i++ {
		if ss[i-2] == ss[i] {
			return true
		}
	}
	return false
}

func (s str) nice() bool {
	if !s.rule1() {
		return false
	}

	if !s.rule2() {
		return false
	}

	if !s.rule3() {
		return false
	}

	return true
}

func (s str) nice2() bool {
	if !s.rule4() {
		return false
	}

	if !s.rule5() {
		return false
	}

	return true
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		i.s = append(i.s, str(l))
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func part1(i input) int {
	sum := 0
	for _, s := range i.s {
		if s.nice() {
			sum++
		}
	}
	return sum
}

func part2(i input) int {
	sum := 0
	for _, s := range i.s {
		if s.nice2() {
			sum++
		}
	}
	return sum
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
