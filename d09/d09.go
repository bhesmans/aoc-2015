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
	dist map[string]map[string]int
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func (i input) _addDest(from, to string, d int) {
	if i.dist[from] == nil {
		i.dist[from] = make(map[string]int)
	}

	i.dist[from][to] = d
}

func (i input) addDest(s string) {
	re := regexp.MustCompile(`^(.*) to (.*) = ([0-9]*)$`)
	t := re.FindStringSubmatch(s)

	from, to := t[1], t[2]
	d := s2i(t[3])

	i._addDest(from, to, d)
	i._addDest(to, from, d)
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{dist: make(map[string]map[string]int)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		i.addDest(l)
	}

	err = scanner.Err()
	panicOnError(err)

	return i
}

func inRoute(route []string, dst string) bool {
	for _, d := range route {
		if d == dst {
			return true
		}
	}
	return false
}

func (in input) routeLen(route []string) int {
	sum := 0
	for i := 1; i < len(route); i++ {
		sum += in.dist[route[i-1]][route[i]]
	}
	return sum
}

func shortestRoute(i input, route []string, minDist, maxDist *int) {
	if len(route) == len(i.dist) {
		rLen := i.routeLen(route)
		if rLen < *minDist {
			*minDist = rLen
		}
		if rLen > *maxDist {
			*maxDist = rLen
		}
		return
	}

	for d, _ := range i.dist {
		if !inRoute(route, d) {
			shortestRoute(i, append(route, d), minDist, maxDist)
		}
	}
}

func part1(i input) (int, int) {
	shortest := 9999999999 // Should be enough ? :/
	longest := 0
	shortestRoute(i, []string{}, &shortest, &longest)
	return shortest, longest
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	p1, p2 := part1(i)
	fmt.Printf("%v\n%v\n", p1, p2)
}
