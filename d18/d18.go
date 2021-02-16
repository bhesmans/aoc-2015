package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type couple struct {
	x, y int
}

type input struct {
	lights  map[couple]bool
	nlights map[couple]bool
	size    int
	part2   bool
}

type fmap func(*input, couple)

func (in *input) foreach(f fmap) {
	for x := 0; x < in.size; x++ {
		for y := 0; y < in.size; y++ {
			f(in, couple{x, y})
		}
	}
}

func (in *input) on(c couple) bool {
	return in.lights[c]
}

func (c *couple) add(c2 couple) {
	c.x += c2.x
	c.y += c2.y
}

func (c *couple) equals(c2 couple) bool {
	return c.x == c2.x && c.y == c2.y
}

func stepFor(in *input, c couple) {
	n := in.around(c)
	on := in.on(c)

	if on && (n == 2 || n == 3) {
		in.nlights[c] = true
	}

	if !on && n == 3 {
		in.nlights[c] = true
	}

	if in.part2 && (c.equals(couple{0, 0}) ||
		c.equals(couple{0, 99}) ||
		c.equals(couple{99, 0}) ||
		c.equals(couple{99, 99})) {
		in.nlights[c] = true
	}

}

func (in *input) step() {
	in.nlights = map[couple]bool{}
	in.foreach(stepFor)
	in.lights = in.nlights
}

func (in *input) around(c couple) int {
	n := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			tc := couple{i, j}
			tc.add(c)

			if in.on(tc) {
				n++
			}
		}
	}
	return n
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func (in *input) addLine(y int, s string) {
	for x, c := range s {
		if c == '#' {
			in.lights[couple{x, y}] = true
		}
	}
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	in := input{lights: map[couple]bool{}, size: 100}
	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		in.addLine(y, l)
		y++
	}

	err = scanner.Err()
	panicOnError(err)

	return in
}

func part1(in *input) int {
	for i := 0; i < 100; i++ {
		in.step()
	}
	return len(in.lights)
}

func part2(in *input) int {
	in.part2 = true
	return part1(in)
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(&i))
	i = readInput("input.txt")
	fmt.Printf("%v\n", part2(&i))
}
