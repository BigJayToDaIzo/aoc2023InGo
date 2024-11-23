package main

import (
	"bufio"
	"fmt"
	"os"

	g "example.com/game"
)

func main() {
	file, err := os.Open("../input.txt")
	g.ErrCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	games := []g.Game{}
	for scanner.Scan() {
		line := scanner.Text()
		game := g.ParseGame(line)
		games = append(games, game)
	}
	powerArr := []int{}
	for _, game := range games {
		maxR, maxG, maxB := 0, 0, 0
		// determine max values for each color
		for _, reveal := range game.Reveals() {
			if reveal.GetRed() > maxR {
				maxR = reveal.GetRed()
			}
			if reveal.GetBlue() > maxB {
				maxB = reveal.GetBlue()
			}
			if reveal.GetGreen() > maxG {
				maxG = reveal.GetGreen()
			}
		}
		maxPower := maxR * maxG * maxB
		powerArr = append(powerArr, maxPower)
	}
	accumulator := 0
	for _, power := range powerArr {
		accumulator += power
	}
	fmt.Println(accumulator)
}
