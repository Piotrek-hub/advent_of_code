package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const FILE_PATH = "./input.txt"

type Elf struct {
	TotalCalories      int
	FoodsCaloriesArray []string
}

func NewElf() *Elf {
	return &Elf{}
}

func (e *Elf) GetTotalCalories() int {
	return e.TotalCalories
}

func (e *Elf) AddToTotalCalories(calories int) {
	e.TotalCalories += calories
}

func (e *Elf) calculateTotalCalories() error {
	for _, v := range e.FoodsCaloriesArray {
		caloriesValue, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		e.AddToTotalCalories(caloriesValue)
	}

	return nil
}

func loadFileContent(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func getCaloriesArrayFromFile(fileContent string) [][]string {
	var finalCaloriesArray [][]string
	calories := strings.Split(fileContent, "\n")

	tempArr := []string{}
	for i := 0; i < len(calories); i++ {
		if len(calories[i]) > 0 {
			tempArr = append(tempArr, calories[i])
		} else if len(calories[i]) == 0 {
			finalCaloriesArray = append(finalCaloriesArray, tempArr)
			tempArr = []string{}
		}

	}

	return finalCaloriesArray
}

func initElfs(caloriesArray [][]string) []*Elf {
	var elfs []*Elf

	for _, v := range caloriesArray {
		e := NewElf()
		e.FoodsCaloriesArray = v
		e.calculateTotalCalories()

		elfs = append(elfs, e)
	}

	return elfs
}

func GetElfWithMostCalories(elfes []*Elf) *Elf {
	var resultElf *Elf
	resultElf = elfes[0]
	for _, elf := range elfes {
		if elf.GetTotalCalories() > resultElf.GetTotalCalories() {
			resultElf = elf
		}
	}

	return resultElf
}

func solve() (int, error) {
	fileContent, err := loadFileContent(FILE_PATH)
	if err != nil {
		return 0, err
	}

	caloriesArray := getCaloriesArrayFromFile(fileContent)

	elfs := initElfs(caloriesArray)

	elfWithMostCalories := GetElfWithMostCalories(elfs)

	return elfWithMostCalories.GetTotalCalories(), nil
}

func main() {
	mostCalories, err := solve()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(mostCalories)
}
