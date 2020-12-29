package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
)

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

type input struct {
	s string
}

func readInput(inputName string) input {
	file, err := os.Open(inputName)
	panicOnError(err)
	defer file.Close()

	i := input{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	l := scanner.Text()
	i.s = l

	err = scanner.Err()
	panicOnError(err)

	return i
}

func checkVal2(in input, i int) bool {
	data := []byte(fmt.Sprintf("%v%v", in.s, i))
	sum := md5.Sum(data)

	for j := 0; j < 3; j++ {
		if sum[j] != 0 {
			return false
		}
	}

	return true
}

func checkVal(in input, i int) bool {
	data := []byte(fmt.Sprintf("%v%v", in.s, i))
	sum := md5.Sum(data)

	for j := 0; j < 2; j++ {
		if sum[j] != 0 {
			return false
		}
	}

	return (sum[2] & 0xf0) == 0
}

func part1(i input) int {
	ret := 0
	for !checkVal(i, ret) {
		ret++
	}
	return ret
}

func part2(i input) int {
	ret := 0
	for !checkVal2(i, ret) {
		ret++
	}
	return ret
}

func main() {
	fmt.Println("Hello")
	i := readInput("input.txt")
	fmt.Printf("%v\n", part1(i))
	fmt.Printf("%v\n", part2(i))
}
