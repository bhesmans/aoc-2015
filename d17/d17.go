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

type input struct {
	boxes    []int
	boxesFit map[int]int
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func (in *input) addLine(s string) {
	in.boxes = append(in.boxes, s2i(s))
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		i.addLine(l)
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func (in *input) combiFit(combi, total int) bool {
	ctotal := 0
	bcount := 0

	for i := 0; i < len(in.boxes); i++ {
		if combi&1 == 1 {
			bcount++
			ctotal += in.boxes[i]
		}
		combi >>= 1
	}

	if ctotal == total {
		in.boxesFit[bcount] += 1
	}
	return ctotal == total
}

func (in *input) combi(total int) int {
	in.boxesFit = map[int]int{}
	count := 0
	maxCombi := 1 << len(in.boxes)

	for i := 0; i < maxCombi; i++ {
		if in.combiFit(i, total) {
			count++
		}
	}

	return count
}

func part1(in *input) int {
	return in.combi(150)
}

func part2(in *input) int {
	minCount := 999
	nCombi := 0
	for count, n := range in.boxesFit {
		if count < minCount {
			minCount = count
			nCombi = n
		}
	}
	return nCombi
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(&i))
	fmt.Printf("%v\n", part2(&i))
}
