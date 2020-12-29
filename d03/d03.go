package main

import (
	"bufio"
	"fmt"
	"os"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type point struct {
	x, y int
}

var dir map[byte]point

type input struct {
	s     string
	homes map[point]int
}

func (p1 *point) add(p2 point) {
	p1.x += p2.x
	p1.y += p2.y
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

func part1(in *input) int {
	in.homes = make(map[point]int)
	current := point{0, 0}
	in.homes[current] += 1
	for i := 0; i < len(in.s); i++ {
		current.add(dir[in.s[i]])
		in.homes[current] += 1
	}
	return len(in.homes)
}

func part2(in *input) int {
	in.homes = make(map[point]int)
	santa := point{0, 0}
	rsanta := point{0, 0}
	in.homes[santa] += 1
	in.homes[rsanta] += 1
	santaTurn := true
	for i := 0; i < len(in.s); i++ {
		if santaTurn {
			santa.add(dir[in.s[i]])
			in.homes[santa] += 1
		} else {
			rsanta.add(dir[in.s[i]])
			in.homes[rsanta] += 1
		}
		santaTurn = !santaTurn
	}
	return len(in.homes)
}

func initDir() {
	dir = make(map[byte]point)
	dir['^'] = point{0, -1}
	dir['>'] = point{1, 0}
	dir['v'] = point{0, 1}
	dir['<'] = point{-1, 0}
}

func main() {
	fmt.Println("Hello")
	initDir()
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(&i))
	fmt.Printf("%v\n", part2(&i))
}
