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

type pwd []byte

type input struct {
	pwd pwd
}

var excluded map[byte]bool

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
	i.pwd = []byte(l)

	err = scanner.Err()
	panicOnError(err)

	return i
}

func (p *pwd) next() {
	for i := len(*p) - 1; i >= 0; i-- {
		(*p)[i] += 1
		if (*p)[i] > 'z' {
			(*p)[i] = 'a'
		} else {
			// input do no contains those
			// not the last letter
			if excluded[(*p)[i]] {
				(*p)[i] += 1
			}
			return
		}
	}

	panic("wrap around !")
}

func (p *pwd) rule1() bool {
	for i := 2; i < len(*p); i++ {
		a, b, c := (*p)[i-2], (*p)[i-1], (*p)[i]

		if a == b-1 && a == c-2 {
			return true
		}
	}

	return false
}

func (p *pwd) pair(from int, not byte) (bool, int) {
	for i := from + 1; i < len(*p); i++ {
		a, b := (*p)[i-1], (*p)[i]
		if a == b && a != not {
			return true, i
		}
	}

	return false, -1
}

func (p *pwd) rule3() bool {
	if ok, i := p.pair(0, ' '); ok {
		ok, _ = p.pair(i+1, (*p)[i])
		return ok
	}
	return false
}

func (p *pwd) valid() bool {
	return p.rule1() && p.rule3()
}

func part1(in input) string {
	for !in.pwd.valid() {
		in.pwd.next()
	}

	return string(in.pwd)
}

func initGlobal() {
	excluded = make(map[byte]bool)
	excluded['i'] = true
	excluded['o'] = true
	excluded['l'] = true
}

func main() {
	fmt.Println("Hello")
	initGlobal()
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	i.pwd.next()
	fmt.Printf("%v\n", part1(i))
}
