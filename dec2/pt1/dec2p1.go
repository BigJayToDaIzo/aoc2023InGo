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
	hypotheticalBag := g.NewReveal(12, 13, 14)
	possibleGameIds := []int{}
	for _, game := range games {
		if game.IsPossible(hypotheticalBag) {
			possibleGameIds = append(possibleGameIds, game.Id())
		}
	}
	solution := 0
	for _, id := range possibleGameIds {
		solution += id
	}
	fmt.Println(solution)
}
