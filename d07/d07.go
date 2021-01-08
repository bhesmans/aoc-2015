package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type gate interface {
	value(i input, cg cgate) int
}

type cgate struct {
	id      string
	defined bool
	val     int
	inputs  []string
	g       gate
}

func (cg *cgate) eval(i input) int {
	if cg.defined {
		return cg.val
	}

	for _, in := range cg.inputs {
		if _, ok := i.gates[in]; !ok {
			i.gates[in] = &cgate{g: raw{}, val: s2i(in)}
		}
		i.gates[in].eval(i)
	}

	cg.defined = true
	cg.val = cg.g.value(i, *cg)
	return cg.val
}

type raw struct{}

func (r raw) value(i input, cg cgate) int {
	return cg.val
}

type wire struct{}

func (w wire) value(i input, cg cgate) int {
	return i.gates[cg.inputs[0]].val
}

type and struct{}

func (a and) value(i input, cg cgate) int {
	return i.gates[cg.inputs[0]].val & i.gates[cg.inputs[1]].val
}

type or struct{}

func (o or) value(i input, cg cgate) int {
	return i.gates[cg.inputs[0]].val | i.gates[cg.inputs[1]].val
}

type lshift struct{}

func (l lshift) value(i input, cg cgate) int {
	return i.gates[cg.inputs[0]].val << i.gates[cg.inputs[1]].val
}

type rshift struct{}

func (r rshift) value(i input, cg cgate) int {
	return i.gates[cg.inputs[0]].val >> i.gates[cg.inputs[1]].val
}

type not struct{}

func (n not) value(i input, cg cgate) int {
	return 0xffff ^ i.gates[cg.inputs[0]].val
}

type input struct {
	gates map[string]*cgate
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func newGate(s string) (*cgate, string) {
	re := regexp.MustCompile(`^(.*) -> (.*)$`)
	t := re.FindStringSubmatch(s)
	left, out := t[1], t[2]

	re = regexp.MustCompile(`^([0-9a-z]*)$`)
	if re.MatchString(left) {
		return &cgate{g: wire{}, inputs: []string{t[1]}}, out
	}

	if strings.Contains(left, "NOT") {
		re = regexp.MustCompile(`^NOT (.*)$`)
		t = re.FindStringSubmatch(left)
		return &cgate{g: not{}, inputs: []string{t[1]}}, out
	}

	re = regexp.MustCompile(`^(.*) (.*) (.*)$`)
	t = re.FindStringSubmatch(left)
	i0, sgate, i1 := t[1], t[2], t[3]

	ret := cgate{inputs: []string{i0, i1}}
	if sgate == "AND" {
		ret.g = and{}
	} else if sgate == "OR" {
		ret.g = or{}
	} else if sgate == "LSHIFT" {
		ret.g = lshift{}
	} else if sgate == "RSHIFT" {
		ret.g = rshift{}
	} else {

		panic(fmt.Sprintf("aaaah %v %#v", s, t))
	}

	return &ret, out
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{gates: make(map[string]*cgate)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		g, out := newGate(l)
		i.gates[out] = g
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func part1(i input) int {
	return i.gates["a"].eval(i)
}

func part2(i input, p1 int) int {
	i.gates["b"].inputs[0] = fmt.Sprintf("%v", p1)
	return part1(i)
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	p1 := part1(i)
	fmt.Printf("%v\n", p1)
	i = readInput("input.txt")
	fmt.Printf("%v\n", part2(i, p1))
}
