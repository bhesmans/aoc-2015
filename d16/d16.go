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

type sue struct {
	prop map[string]int
}

type input struct {
	sues map[int]sue
}

var target sue

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func (in *input) addLine(s string) {
	re := regexp.MustCompile(`^Sue ([0-9]*):(.*)$`)
	t := re.FindStringSubmatch(s)

	sue := sue{prop: map[string]int{}}
	i := s2i(t[1])

	s = t[2]

	re = regexp.MustCompile(`^ ([a-z]*): (-?[0-9]*),?(.*)$`)
	for t = re.FindStringSubmatch(s); len(s) != 0; t = re.FindStringSubmatch(s) {
		sue.prop[t[1]] = s2i(t[2])
		s = t[3]
	}

	in.sues[i] = sue
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{sues: map[int]sue{}}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		i.addLine(l)
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func (in *input) filterOut(prop string, n int) {
	for k, v := range in.sues {
		nn, ok := v.prop[prop]

		if !ok {
			continue
		}

		if prop == "cats" || prop == "trees" {
			if nn <= n {
				delete(in.sues, k)
			}
			continue
		} else if prop == "pomeranians" || prop == "goldfish" {
			if nn >= n {
				delete(in.sues, k)
			}
			continue
		} else if nn != n {
			delete(in.sues, k)
		}
	}
}

func (in *input) filterOut2(prop string, n int) {
	for k, v := range in.sues {
		if nn, ok := v.prop[prop]; ok && nn != n {
			delete(in.sues, k)
		}
	}
}

func (in *input) findTarget(target sue, part2 bool) int {
	for k, v := range target.prop {
		if part2 {
			in.filterOut(k, v)
		} else {
			in.filterOut2(k, v)
		}
	}

	if len(in.sues) != 1 {
		panic("aaaaaaaaaah")
	}

	for k, _ := range in.sues {
		return k
	}

	return 0
}

func part1(in input) int {
	return in.findTarget(target, false)
}

func part2(in input) int {
	return in.findTarget(target, true)
}

func initTarget() {
	target = sue{prop: map[string]int{}}
	target.prop["children"] = 3
	target.prop["cats"] = 7
	target.prop["samoyeds"] = 2
	target.prop["pomeranians"] = 3
	target.prop["akitas"] = 0
	target.prop["vizslas"] = 0
	target.prop["goldfish"] = 5
	target.prop["trees"] = 3
	target.prop["cars"] = 2
	target.prop["perfumes"] = 1
}

func main() {
	fmt.Println("Hello")
	initTarget()
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	i = readInput("input.txt")
	fmt.Printf("%v\n", part2(i))
}
