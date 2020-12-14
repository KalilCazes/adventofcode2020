package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func checkPassword(password []string) int {

	first, err := strconv.Atoi(password[0])
	checkError(err)
	second, err := strconv.Atoi(password[1])
	checkError(err)

	letter := password[2]
	passwd := password[3]

	if passwd[first-1] != passwd[second-1] && (string(passwd[first-1]) == letter || string(passwd[second-1]) == letter) {
		return 1
	}
	return 0
}

func readFile(fileName string) []string {
	var lines []string

	f, err := os.Open(fileName)
	checkError(err)
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := scanner.Text()
		re := regexp.MustCompile("[\\:\\-\\s]+").Split(s, -1)
		first := re[0]
		second := re[1]
		letter := re[2]
		passwd := re[3]
		lines = append(lines, first, second, letter, passwd)
	}

	return lines
}

func main() {
	var step, result int
	const batch = 4

	password := readFile("../input")

	for i := batch; i < len(password)+1; i += batch {
		result += checkPassword(password[step:i])
		step = i
	}
	fmt.Println(result)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
