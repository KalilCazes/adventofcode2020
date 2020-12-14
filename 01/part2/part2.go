package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sliceToMap(numbers []int) map[int]int {
	m := make(map[int]int)
	for _, v := range numbers {
		m[v]++
	}
	return m
}

func readFile(fileName string) []int {
	var intInput []int

	f, err := os.Open(fileName)
	checkError(err)
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		strings := scanner.Text()
		tmp, err := strconv.Atoi(strings)
		checkError(err)
		intInput = append(intInput, tmp)
	}

	return intInput
}

func main() {
	var result, solution int
	numbers := readFile("../input")
	m := sliceToMap(numbers)

	for _, i := range numbers {
		for _, j := range numbers {
			result = 2020 - (i + j)
			if _, ok := m[result]; ok {
				solution = result * j * i
				break
			}
		}
	}

	fmt.Println(solution)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
