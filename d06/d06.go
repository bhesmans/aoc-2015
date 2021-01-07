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

type ins interface {
	execute(i input, p point)
	executeAll(i input, ins ins)
}

type point struct {
	x, y int
}

type cins struct {
	c1, c2 point
}

func (ci cins) executeAll(i input, ins ins) {
	for x := ci.c1.x; x <= ci.c2.x; x++ {
		for y := ci.c1.y; y <= ci.c2.y; y++ {
			ins.execute(i, point{x, y})
		}
	}
}

type toggle struct {
	cins
}

type turnOn struct {
	cins
}

type turnOff struct {
	cins
}

func (t toggle) execute(i input, p point) {
	if i.on[p] {
		delete(i.on, p)
	} else {
		i.on[p] = true
	}
	i.bright[p] += 2
}

func (t turnOn) execute(i input, p point) {
	i.on[p] = true
	i.bright[p]++
}

func (t turnOff) execute(i input, p point) {
	delete(i.on, p)
	i.bright[p]--
	if i.bright[p] < 0 {
		i.bright[p] = 0
	}
}

type input struct {
	inss   []ins
	on     map[point]bool
	bright map[point]int
}

func newCIns(s string) cins {
	return cins{}
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func newIns(s string) ins {
	re := regexp.MustCompile(`^(toggle|turn on|turn off) ([0-9]*),([0-9]*) through ([0-9]*),([0-9]*)$`)
	t := re.FindStringSubmatch(s)
	ci := cins{point{s2i(t[2]), s2i(t[3])}, point{s2i(t[4]), s2i(t[5])}}

	if t[1] == "toggle" {
		return toggle{ci}
	}

	if t[1] == "turn on" {
		return turnOn{ci}
	}

	if t[1] == "turn off" {
		return turnOff{ci}
	}

	panic("haaaaaaaaa")

}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{on: make(map[point]bool), bright: make(map[point]int)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		i.inss = append(i.inss, newIns(l))
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func part1(i input) int {
	for _, ins := range i.inss {
		ins.executeAll(i, ins)
	}
	return len(i.on)
}

func part2(i input) int {
	sum := 0
	for _, v := range i.bright {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
