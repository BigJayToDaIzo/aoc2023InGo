package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	errCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	codes := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		codes = append(codes, parseLine(line))
	}
	fmt.Println(sumCodes(codes))
}

func sumCodes(codes []int) int {
	acc := 0
	for _, code := range codes {
		acc += code
	}
	return acc
}

func parseLine(line string) int {
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	lowIndex := math.MaxInt32
	lowIdxDigit := 0
	highIndex := -1
	highIdxDigit := 0
	for word, digit := range digits {
		if i := strings.Index(line, word); i != -1 {
			if i < lowIndex {
				lowIndex = i
				lowIdxDigit = digit
			}
		}
		if i := strings.LastIndex(line, word); i != -1 {
			if i > highIndex {
				highIndex = i
				highIdxDigit = digit
			}
		}
		if i := strings.Index(line, strconv.Itoa(digit)); i != -1 {
			if i < lowIndex {
				lowIndex = i
				lowIdxDigit = digit
			}
		}
		if i := strings.LastIndex(line, strconv.Itoa(digit)); i != -1 {
			if i > highIndex {
				highIndex = i
				highIdxDigit = digit
			}
		}
	}
	return lowIdxDigit*10 + highIdxDigit
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
