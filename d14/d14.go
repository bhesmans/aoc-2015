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

type reindeer struct {
	name             string
	speed, fly, rest int
	currentPos       int
	points           int
}

type input struct {
	reindeers  []reindeer
	currentSec int
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func (i *input) addLine(s string) {
	re := regexp.MustCompile(`^(.*) can fly ([0-9]*) km/s for ([0-9]*) seconds, but then must rest for ([0-9]*) seconds\.$`)
	t := re.FindStringSubmatch(s)

	name, speed, fly, rest := t[1], t[2], t[3], t[4]
	i.reindeers = append(i.reindeers,
		reindeer{
			name:  name,
			speed: s2i(speed),
			fly:   s2i(fly),
			rest:  s2i(rest)})
}

func (r reindeer) cycle() int {
	return r.fly + r.rest
}

func (r *reindeer) after(sec int) int {
	km := (sec / r.cycle()) * (r.speed * r.fly)

	rest := (sec % r.cycle())

	if rest >= r.fly {
		rest = r.fly
	}

	tot := km + (r.speed * rest)

	r.currentPos = tot

	return tot
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

func (in *input) step() {
	in.currentSec++
	max := 0
	for i, _ := range in.reindeers {
		r := &(in.reindeers[i])
		dist := r.after(in.currentSec)
		if dist > max {
			max = dist
		}
	}

	for i, _ := range in.reindeers {
		r := &(in.reindeers[i])
		if r.currentPos == max {
			r.points++
		}
	}
}

func part1(i input) int {
	max := 0
	for _, r := range i.reindeers {
		dist := r.after(2503)
		if dist > max {
			max = dist
		}
	}
	return max
}

func part2(in input) int {
	for i := 0; i < 2503; i++ {
		in.step()
	}

	max := 0
	for _, r := range in.reindeers {
		if r.points > max {
			max = r.points
		}
	}

	return max
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
