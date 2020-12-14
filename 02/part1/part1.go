package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkPassword(password []string) int {

	min, err := strconv.Atoi(password[0])
	checkError(err)
	max, err := strconv.Atoi(password[1])
	checkError(err)

	letter := password[2]
	passwd := password[3]

	result := strings.Count(passwd, letter)

	if result >= min && result <= max {
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
		min := re[0]
		max := re[1]
		letter := re[2]
		passwd := re[3]
		lines = append(lines, min, max, letter, passwd)
	}

	return lines
}

func main() {
	var step, result int
	const batch = 4

	password := readFile("../input")

	for i := batch; i < len(password); i += batch {
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
