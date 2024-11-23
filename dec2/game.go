package game

import (
	"strconv"
	"strings"
	"unicode"
)

type Reveal struct {
	red   int
	green int
	blue  int
}

func NewReveal(red, green, blue int) *Reveal {
	return &Reveal{
		red:   red,
		green: green,
		blue:  blue,
	}
}

type Game struct {
	id      int
	reveals []Reveal
}

func (g Game) Id() int {
	return g.id
}

func (g Game) Reveals() []Reveal {
	return g.reveals
}

func (r *Reveal) SetRed(red int) {
	r.red = red
}

func (r *Reveal) SetGreen(green int) {
	r.green = green
}

func (r *Reveal) SetBlue(blue int) {
	r.blue = blue
}

func (r Reveal) GetRed() int   { return r.red }
func (r Reveal) GetGreen() int { return r.green }
func (r Reveal) GetBlue() int  { return r.blue }

func (g Game) IsPossible(hypeBag *Reveal) bool {
	for _, reveal := range g.reveals {
		if reveal.red > hypeBag.red || reveal.green > hypeBag.green || reveal.blue > hypeBag.blue {
			return false
		}
	}
	return true
}

func ParseGame(line string) Game {
	l := strings.TrimPrefix(line, "Game ")
	id, err := strconv.Atoi(l[:strings.Index(l, ":")])
	ErrCheck(err)
	l = l[strings.Index(l, ":")+2:]
	reveals := strings.SplitAfterN(l, ";", -1)
	revealArray := []Reveal{}
	for _, r := range reveals {
		r = strings.TrimSuffix(r, ";")
		dice := strings.Split(r, ",")
		reveal := Reveal{}
		for _, die := range dice {
			die = strings.TrimPrefix(die, " ")
			numDice := die[:strings.Index(die, " ")]
			nd, err := strconv.Atoi(numDice)
			ErrCheck(err)
			trimFunc := func(r rune) bool {
				return !unicode.IsLetter(r)
			}
			switch strings.TrimLeftFunc(die, trimFunc) {
			case "red":
				reveal.red = nd
			case "green":
				reveal.green = nd
			case "blue":
				reveal.blue = nd
			}
		}
		revealArray = append(revealArray, reveal)
	}
	return Game{
		id:      id,
		reveals: revealArray,
	}
}

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}
