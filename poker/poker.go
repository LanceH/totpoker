package poker

import (
	"fmt"
	"strconv"
	"strings"
)

type Totp struct {
	Num     int
	S       string
	Digits  []int
	SDigits string
	Pairs   [3]string
	Trips   [2]string
}

func NewTotp(n int) (t Totp) {
	t = Totp{Num: n}
	t.S = fmt.Sprintf("%06d", n)

	// Count the number of each type of digit
	t.Digits = make([]int, 10)
	for _, c := range t.S {
		i, _ := strconv.Atoi(string(c))
		t.Digits[i] = t.Digits[i] + 1
	}

	// Turn the digit count into a string
	// useful for finding unsorted straights
	for _, v := range t.Digits {
		t.SDigits = t.SDigits + strconv.Itoa(v)
	}

	// initialize pairs
	t.Pairs[0] = t.S[0:2]
	t.Pairs[1] = t.S[2:4]
	t.Pairs[2] = t.S[4:6]

	// initialize trips
	t.Trips[0] = t.S[0:3]
	t.Trips[1] = t.S[3:6]
	return
}

func (t Totp) CheckSixOfAKind() bool {
	if strings.Index(t.SDigits, "6") > -1 {
		return true
	}
	return false
}

func (t Totp) CheckFiveOfAKind() bool {
	if strings.Index(t.SDigits, "5") > -1 {
		return true
	}
	return false
}

func (t Totp) CheckFourOfAKind() bool {
	if strings.Index(t.SDigits, "4") > -1 {
		return true
	}
	return false
}

func (t Totp) CheckDoubleTriple() bool {
	triples := 0
	for _, v := range t.Digits {
		if v == 3 {
			triples++
		}
	}
	if triples == 2 {
		return true
	}
	return false
}
