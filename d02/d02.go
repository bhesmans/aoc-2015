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

type box struct {
	a, b, c int
}

type input struct {
	boxes []box
}

func (b box) paper() int {
	ab := b.a * b.b
	ac := b.a * b.c
	bc := b.b * b.c
	small := 0

	if ab <= ac && ab <= bc {
		small = ab
	} else if ac <= ab && ac <= bc {
		small = ac
	} else {
		small = bc
	}

	return 2*ab + 2*ac + 2*bc + small
}

func (b box) ribbon() int {
	p1 := 2*b.a + 2*b.b
	p2 := 2*b.a + 2*b.c
	p3 := 2*b.b + 2*b.c
	small := 0

	if p1 <= p2 && p1 <= p3 {
		small = p1
	} else if p2 <= p1 && p2 <= p3 {
		small = p2
	} else {
		small = p3
	}

	return small + b.a*b.b*b.c
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		t := strings.Split(l, "x")
		b := box{}
		b.a, _ = strconv.Atoi(t[0])
		b.b, _ = strconv.Atoi(t[1])
		b.c, _ = strconv.Atoi(t[2])
		i.boxes = append(i.boxes, b)
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func part1(i input) int {
	sum := 0
	for _, b := range i.boxes {
		sum += b.paper()
	}
	return sum
}

func part2(i input) int {
	sum := 0
	for _, b := range i.boxes {
		sum += b.ribbon()
	}
	return sum
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
