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

type ingredient struct {
	props []int
}

type input struct {
	ingredients []ingredient
}

func s2i(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func (in *input) addLine(s string) {
	re := regexp.MustCompile(`^(.*):(.*)$`)
	t := re.FindStringSubmatch(s)
	s = t[2]

	i := ingredient{}
	re = regexp.MustCompile(`^ ([a-z]*) (-?[0-9]*),?(.*)$`)
	for t = re.FindStringSubmatch(s); len(s) != 0; t = re.FindStringSubmatch(s) {
		i.props = append(i.props, s2i(t[2]))
		s = t[3]
	}

	in.ingredients = append(in.ingredients, i)
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

func (in *input) calories(qties []int) int {
	cal := 0
	cali := len(in.ingredients[0].props) - 1

	for qtyi, qty := range qties {
		cal += (qty * in.ingredients[qtyi].props[cali])
	}

	return cal
}

func (in *input) score(qties []int) int {
	score := 1
	for pi := range in.ingredients[0].props {
		if pi == len(in.ingredients[0].props)-1 {
			// Ignore calories for the score
			continue
		}

		prop := 0
		for qtyi, qty := range qties {
			prop += (qty * in.ingredients[qtyi].props[pi])
		}

		if prop <= 0 {
			return 0
		}

		score *= prop
	}

	return score
}

func (in *input) maxScore(qties []int, left int, max, max500 *int) {
	if len(qties) == len(in.ingredients)-1 {
		// last ingredient takes the rest
		qties = append(qties, left)
		score := in.score(qties)
		cal := in.calories(qties)

		if score > *max {
			*max = score
		}

		if cal == 500 && score > *max500 {
			*max500 = score
		}

		return
	}

	for i := 0; i <= left; i++ {
		in.maxScore(append(qties, i), left-i, max, max500)
	}
}

func part12(in input) (int, int) {
	max, max500 := 0, 0
	in.maxScore([]int{}, 100, &max, &max500)
	return max, max500
}

func part2(in input) int {
	return 0
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	p1, p2 := part12(i)
	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)
}
