package main

import (
	"log"
	"os"
	"strings"
)

const (
	FILE_PATH = "./02/input"
)

const (
	RoundWon  = 6
	RoundDraw = 3
	RoundLost = 0
)

const (
	PLAYER_ROCK     = "X"
	PLAYER_PAPER    = "Y"
	PLAYER_SCISSORS = "Z"

	ELF_ROCK     = "A"
	ELF_PAPER    = "B"
	ELF_SCISSORS = "C"
)

type choice string

type Round struct {
	playerChoice choice
	elfChoice    choice
	status       int
	scoredPoints int
}

func NewRound(elfChoice, playerChoice choice) *Round {
	return &Round{
		playerChoice: playerChoice,
		elfChoice:    elfChoice,
	}
}

func (r *Round) CalculateStatus() {
	if r.playerChoice == PLAYER_ROCK {
		if r.elfChoice == ELF_PAPER {
			r.status = RoundLost
		} else if r.elfChoice == ELF_SCISSORS {
			r.status = RoundWon
		} else {
			r.status = RoundDraw
		}
	} else if r.playerChoice == PLAYER_PAPER {
		if r.elfChoice == ELF_SCISSORS {
			r.status = RoundLost
		} else if r.elfChoice == ELF_ROCK {
			r.status = RoundWon
		} else {
			r.status = RoundDraw
		}
	} else if r.playerChoice == PLAYER_SCISSORS {
		if r.elfChoice == ELF_ROCK {
			r.status = RoundLost
		} else if r.elfChoice == ELF_PAPER {
			r.status = RoundWon
		} else {
			r.status = RoundDraw
		}
	}
}

func (r *Round) CalculatePlayerChoice() {
	if r.status == RoundWon {
		if r.elfChoice == ELF_ROCK {
			r.playerChoice = PLAYER_PAPER
		} else if r.elfChoice == ELF_PAPER {
			r.playerChoice = PLAYER_SCISSORS
		} else if r.elfChoice == ELF_SCISSORS {
			r.playerChoice = PLAYER_ROCK
		}
	} else if r.status == RoundDraw {
		if r.elfChoice == ELF_ROCK {
			r.playerChoice = PLAYER_ROCK
		} else if r.elfChoice == ELF_PAPER {
			r.playerChoice = PLAYER_PAPER
		} else if r.elfChoice == ELF_SCISSORS {
			r.playerChoice = PLAYER_SCISSORS
		}
	} else if r.status == RoundLost {
		if r.elfChoice == ELF_ROCK {
			r.playerChoice = PLAYER_SCISSORS
		} else if r.elfChoice == ELF_PAPER {
			r.playerChoice = PLAYER_ROCK
		} else if r.elfChoice == ELF_SCISSORS {
			r.playerChoice = PLAYER_PAPER
		}
	}
}

func (r *Round) CalculateScoredPoints() {
	var points int
	points += r.status

	if r.playerChoice == PLAYER_ROCK {
		points += 1
	} else if r.playerChoice == PLAYER_PAPER {
		points += 2
	} else if r.playerChoice == PLAYER_SCISSORS {
		points += 3
	}

	r.scoredPoints = points
}

func loadFileContent(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func parseFileContentIntoRounds(fileContent string) []*Round {
	var rounds []*Round
	parsedContent := strings.Split(fileContent, "\n")

	for _, r := range parsedContent {
		if len(r) > 0 {
			elfChoice := choice(r[0])
			playerChoice := choice(r[2])
			rounds = append(rounds, NewRound(elfChoice, playerChoice))
		}
	}

	return rounds
}

func choiceToRoundStatus(ch choice) int {
	var roundStatus int

	switch ch {
	case "X":
		roundStatus = RoundLost
		break
	case "Y":
		roundStatus = RoundDraw
		break
	case "Z":
		roundStatus = RoundWon
		break
	default:
		break
	}

	return roundStatus
}

func parseFileContentIntoExpectedRounds(fileContent string) []*Round {
	var rounds []*Round
	parsedContent := strings.Split(fileContent, "\n")

	for _, r := range parsedContent {
		if len(r) > 0 {
			elfChoice := choice(r[0])
			expectedStatus := choice(r[2])
			rounds = append(rounds, &Round{
				elfChoice: elfChoice,
				status:    choiceToRoundStatus(expectedStatus),
			})
		}
	}

	return rounds
}

func calculatePlayerChoiceForAllRounds(rounds []*Round) {
	for _, round := range rounds {
		round.CalculatePlayerChoice()
	}
}

func calculateStatusForAllRounds(rounds []*Round) {
	for _, round := range rounds {
		round.CalculateStatus()
	}
}

func calculateScoreForAllRounds(rounds []*Round) {
	for _, round := range rounds {
		round.CalculateScoredPoints()
	}
}

func getTotalPointsFromAllRounds(rounds []*Round) int {
	var totalScoredPoints int
	for _, round := range rounds {
		totalScoredPoints += round.scoredPoints
	}

	return totalScoredPoints
}

func solve() (int, error) {
	fileContent, err := loadFileContent(FILE_PATH)
	if err != nil {
		return 0, err
	}

	rounds := parseFileContentIntoRounds(fileContent)
	expectedRounds := parseFileContentIntoExpectedRounds(fileContent)

	calculateStatusForAllRounds(rounds)
	calculatePlayerChoiceForAllRounds(expectedRounds)

	calculateScoreForAllRounds(rounds)
	calculateScoreForAllRounds(expectedRounds)

	for i, r := range expectedRounds {
		log.Println(i+1, r)
	}

	totalScoredPoints := getTotalPointsFromAllRounds(rounds)
	totalScoredPointsFromExpected := getTotalPointsFromAllRounds(expectedRounds)

	log.Println(totalScoredPointsFromExpected)

	return totalScoredPoints, nil
}

func main() {
	solve()
}
