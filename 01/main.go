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

func getElfWithMostCalories(elfs []*Elf) *Elf {
	var resultElf *Elf
	resultElf = elfs[0]
	for _, elf := range elfs {
		if elf.GetTotalCalories() > resultElf.GetTotalCalories() {
			resultElf = elf
		}
	}

	return resultElf
}

func getThreeElfsWithTopCalories(elfs []*Elf) []*Elf {
	top1, top2, top3 := elfs[0], elfs[1], elfs[2]
	for _, elf := range elfs {
		if elf.GetTotalCalories() > top3.GetTotalCalories() {
			if elf.GetTotalCalories() > top2.GetTotalCalories() {
				if elf.GetTotalCalories() > top1.GetTotalCalories() {
					top2 = top1
					top3 = top2
					top1 = elf
					continue
				}
				top3 = top2
				top2 = elf
				continue
			}
			top3 = elf
			continue
		}
	}

	return []*Elf{top1, top2, top3}
}

func sumElfsCalories(elfs []*Elf) int {
	var sum int
	for _, elf := range elfs {
		sum += elf.GetTotalCalories()
	}

	return sum
}

func solve() (int, int, error) {
	fileContent, err := loadFileContent(FILE_PATH)
	if err != nil {
		return 0, 0, err
	}

	caloriesArray := getCaloriesArrayFromFile(fileContent)

	elfs := initElfs(caloriesArray)

	elfWithMostCalories := getElfWithMostCalories(elfs)

	topThreeElfs := getThreeElfsWithTopCalories(elfs)

	topThreeSum := sumElfsCalories(topThreeElfs)

	return elfWithMostCalories.GetTotalCalories(), topThreeSum, nil
}

func main() {
	mostCalories, topThreeCaloriesSum, err := solve()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(mostCalories, " \n ", topThreeCaloriesSum)
}
