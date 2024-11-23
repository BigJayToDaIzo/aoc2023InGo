package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("../input.txt")
	errCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	codes := []int{}
	var code int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		trimFunc := func(r rune) bool {
			return !unicode.IsDigit(r)
		}
		line = strings.TrimFunc(line, trimFunc)
		if len(line) == 1 {
			line += line
			code, err = strconv.Atoi(line)
			errCheck(err)
		} else {
			line = string(line[0]) + string(line[len(line)-1])
			code, err = strconv.Atoi(string(line))
			errCheck(err)
		}
		codes = append(codes, code)
	}
	accumulator := 0
	for _, code := range codes {
		accumulator += code
	}
	fmt.Println(accumulator)
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
