package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	fmt.Println(lines)

	possibleSum := 0
	powerSum := 0
	for i, line := range lines {
		r := regexp.MustCompile(`Game (\d+): (.*)`)
		matches := r.FindStringSubmatch(line)
		game := ParseGame(matches[2])

		isPossible := IsGamePossible(game)
		if isPossible {
			possibleSum += i + 1
		}

		power := CalculatePower(game)
		powerSum += power

		fmt.Println(i+1, isPossible, power, "->", game)
	}

	fmt.Println("Possible Sum", possibleSum)
	fmt.Println("Power Sum", powerSum)
}

type Game struct {
	rounds []Round
}

type Round struct {
	blue  int
	green int
	red   int
}

func CalculatePower(game Game) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, round := range game.rounds {
		if round.red >= maxRed {
			maxRed = round.red
		}
		if round.green >= maxGreen {
			maxGreen = round.green
		}
		if round.blue >= maxBlue {
			maxBlue = round.blue
		}
	}

	return maxRed * maxGreen * maxBlue
}

func IsGamePossible(game Game) bool {

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	isPossible := true
	for _, round := range game.rounds {
		if round.red > maxRed || round.green > maxGreen || round.blue > maxBlue {
			isPossible = false
		}
	}
	return isPossible
}

func ParseGame(gameStr string) Game {
	game := Game{}

	rounds := strings.Split(gameStr, ";")
	for _, roundStr := range rounds {
		game.rounds = append(game.rounds, ParseRound(roundStr))
	}

	return game
}

func ParseRound(roundStr string) Round {
	round := Round{}
	coloursStr := strings.Split(roundStr, ", ")
	for _, colourStr := range coloursStr {
		split := strings.Split(strings.TrimSpace(colourStr), " ")
		count, _ := strconv.Atoi(split[0])
		colour := split[1]
		switch colour {
		case "blue":
			round.blue = count
		case "green":
			round.green = count
		case "red":
			round.red = count
		}
	}
	return round
}
