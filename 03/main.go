package main

import (
	"log"
	"os"
	"strings"
)

const FILE_PATH = "./03/input"
const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Backpack struct {
	content           string
	firstCompartment  string
	secondCompartment string
	sharedItem        string
}

func NewBackpack(content string) *Backpack {
	return &Backpack{content: content}
}

func (b *Backpack) DivideItemsIntoCompartments() {
	b.firstCompartment = b.content[0 : len(b.content)/2]
	b.secondCompartment = b.content[len(b.content)/2:]
}

func (b *Backpack) FindSharedItem() {
	for _, fl := range b.firstCompartment {
		for _, sl := range b.secondCompartment {
			if fl == sl {
				b.sharedItem = string(fl)
				return
			}
		}
	}
}

func (b *Backpack) GetSharedItemValue() int {
	return GetItemPriority(b.sharedItem)
}

func FindSharedItemForAllBackpacks(backpacks []*Backpack) {
	for _, b := range backpacks {
		b.FindSharedItem()
	}
}

func getBackpacksFromFileContent(fileContent string) []*Backpack {
	var backpacks []*Backpack
	contentArray := strings.Split(fileContent, "\n")

	for _, backpackContent := range contentArray {
		b := NewBackpack(backpackContent)
		b.DivideItemsIntoCompartments()
		backpacks = append(backpacks, b)
	}

	return backpacks
}

func SumBackpacksSharedItemsValue(backpacks []*Backpack) int {
	var sum int
	for _, b := range backpacks {
		sum += b.GetSharedItemValue()
	}
	return sum
}

type Group struct {
	backpacks    []*Backpack
	mainItem     string
	itemPriority int
}

func (g *Group) FindMainItem() {
	log.Println(g.backpacks[0].content)
	log.Println(g.backpacks[1].content)
	log.Println(g.backpacks[2].content)
	log.Println("|----------------|")
	for _, first := range g.backpacks[0].content {
		for _, second := range g.backpacks[1].content {
			for _, third := range g.backpacks[2].content {
				if first == second && first == third {
					g.mainItem = string(first)
				}
			}
		}
	}
}

func (g *Group) AddBackpack(backpack *Backpack) bool {
	if len(g.backpacks) == 3 {
		return false
	}
	g.backpacks = append(g.backpacks, backpack)
	return true
}

func (g *Group) SetMainItemPriority() {
	g.itemPriority = GetItemPriority(g.mainItem)
}

func (g *Group) GetMainItemPriority() int {
	return g.itemPriority
}

func GetItemPriority(item string) int {
	return strings.Index(ALPHABET, item) + 1
}

func LoadFileContent(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func getGroupsFromFileContent(fileContent string) []*Group {
	var groups []*Group
	backpacks := getBackpacksFromFileContent(fileContent)

	tempGroup := &Group{}
	for _, backpack := range backpacks {
		if success := tempGroup.AddBackpack(backpack); success == false {
			groups = append(groups, tempGroup)
			tempGroup = &Group{}
			tempGroup.AddBackpack(backpack)
		}
	}

	groups = append(groups, tempGroup)

	return groups
}

func findAllGroupsMainItemAndSetPriority(groups []*Group) {
	for _, group := range groups {
		group.FindMainItem()
		group.SetMainItemPriority()
	}
}

func sumAllGroupsMainItemPriority(groups []*Group) int {
	var sum int
	for _, group := range groups {
		sum += group.GetMainItemPriority()
	}

	return sum
}

func solve() error {
	fileContent, err := LoadFileContent(FILE_PATH)
	if err != nil {
		return err
	}

	//backpacks := getBackpacksFromFileContent(fileContent)
	//
	//FindSharedItemForAllBackpacks(backpacks)
	//
	//log.Println(SumBackpacksSharedItemsValue(backpacks))

	groups := getGroupsFromFileContent(fileContent)

	findAllGroupsMainItemAndSetPriority(groups)

	log.Println(sumAllGroupsMainItemPriority(groups))

	return nil
}

func main() {
	solve()
}
