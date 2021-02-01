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

type input struct {
	happiness map[string]map[string]int
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func (i input) _addLine(from, to string, d int) {
	if i.happiness[from] == nil {
		i.happiness[from] = make(map[string]int)
	}

	i.happiness[from][to] = d
}

func (i input) addLine(s string) {
	re := regexp.MustCompile(`^(.*) would (lose|gain) ([0-9]*) happiness units by sitting next to (.*)\.$`)
	t := re.FindStringSubmatch(s)

	from, to := t[1], t[4]
	gain := s2i(t[3])

	if t[2] == "lose" {
		gain *= (-1)
	}

	i._addLine(from, to, gain)
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{happiness: make(map[string]map[string]int)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		i.addLine(l)
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func onTable(table []string, guest string) bool {
	for _, g := range table {
		if g == guest {
			return true
		}
	}
	return false
}

func prec(i, max int) int {
	i -= 1
	if i < 0 {
		return max - 1
	}
	return i
}

func next(i, max int) int {
	return (i + 1) % max
}

func (in input) tableHappiness(table []string) int {
	sum := 0
	n := len(in.happiness)

	for i, g := range table {
		sum += in.happiness[g][table[prec(i, n)]]
		sum += in.happiness[g][table[next(i, n)]]
	}

	return sum
}

func tableHappinessMinMax(i input, table []string, min, max *int) {
	if len(table) == len(i.happiness) {
		h := i.tableHappiness(table)
		if h < *min {
			*min = h
		}
		if h > *max {
			*max = h
		}
		return
	}

	for g, _ := range i.happiness {
		if !onTable(table, g) {
			tableHappinessMinMax(i, append(table, g), min, max)
		}
	}
}

func part1(i input) int {
	min, max := -9999999, 0
	tableHappinessMinMax(i, []string{}, &min, &max)
	return max
}

func part2(i input) int {
	min, max := -9999999, 0

	i.happiness["me"] = make(map[string]int)
	for g, _ := range i.happiness {
		i.happiness[g]["me"] = 0
		i.happiness["me"][g] = 0
	}

	tableHappinessMinMax(i, []string{}, &min, &max)
	return max
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
