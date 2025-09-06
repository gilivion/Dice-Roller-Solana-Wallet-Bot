package models

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var validDices = map[string]int{
	"d4":    4,
	"d6":    6,
	"d8":    8,
	"d10":   10,
	"d12":   12,
	"d20":   20,
	"d100":  100,
	"d1000": 1000,
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RollMultipleDice(diceType string, numberOfDice int) (int, []int) {
	sides := validDices[diceType]
	total := 0
	rolls := make([]int, numberOfDice)

	for i := 0; i < numberOfDice; i++ {
		roll := rand.Intn(sides) + 1
		rolls[i] = roll
		total += roll
	}

	return total, rolls
}

func ParseRollCommand(input string) (int, string, int, error) {
	re := regexp.MustCompile(`(?i)(\d*)d(\d+)(?:\s*\+\s*(\d+))?`)
	matches := re.FindStringSubmatch(input)

	if len(matches) == 0 {
		return 0, "", 0, fmt.Errorf("Incorrect command")
	}
	numOfDice := 1
	if matches[1] != "" {
		numOfDice, _ = strconv.Atoi(matches[1])
	}
	diceType := "d" + matches[2]
	modifier := 0
	if matches[3] != "" {
		modifier, _ = strconv.Atoi(matches[3])
	}

	return numOfDice, diceType, modifier, nil
}
func IsValidDice(diceType string) bool {
	_, exists := validDices[diceType]
	return exists
}
func RollDice(diceType string) int {
	sides := validDices[diceType]
	return rand.Intn(sides) + 1
}

