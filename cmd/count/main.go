package main

import (
	"fmt"
	"sort"

	"github.com/LanceH/totpoker/poker"
)

var count map[string]int

const SIXOFAKIND = "Six of a Kind"
const FIVEOFAKIND = "Five of a Kind"
const FOUROFAKIND = "Four of a Kind"
const DOUBLETRIPLE = "Double Triple"

func main() {
	for i := 0; i < 1000000; i++ {
		t := poker.NewTotp(i)
		check(t)
	}
	keys := []string{}
	for k, _ := range count {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] < count[keys[j]]
	})
	for _, v := range keys {
		fmt.Printf("%20v: %6d\n", v, count[v])
	}
}

func check(t poker.Totp) {
	if t.CheckSixOfAKind() {
		count[SIXOFAKIND] += 1
	}
	if t.CheckFiveOfAKind() {
		count[FIVEOFAKIND] += 1
	}
	if t.CheckFourOfAKind() {
		count[FOUROFAKIND] += 1
	}
	if t.CheckDoubleTriple() {
		count[DOUBLETRIPLE] += 1
	}
}

func init() {
	count = make(map[string]int)
}
