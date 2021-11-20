package wanderlust

import (
	"errors"
	"strings"
)

type Budget int8
const (
	_ Budget = iota
	Budget1
	Budget2
	Budget3
	Budget4
	Budget5
)

var BudgetStrings = map[string]Budget{
	"$":     Budget1,
	"$$":    Budget2,
	"$$$":   Budget3,
	"$$$$":  Budget4,
	"$$$$$": Budget5,
}

func (l Budget) String() string {
	for s, v := range BudgetStrings {
		if l == v {
			return s
		}
	}
	return "invalid"
}

//BudgetString to Budget
func ParseBudget(s string) Budget {
	return BudgetStrings[s]
}

type BudgetRange struct {
	From Budget
	To   Budget
}

//BudgetRange to String
func (r BudgetRange) String() string {
	return r.From.String() + "..." + r.To.String()
}

//BudgetRangeString to Budgetrange
func ParseBudgetRange(s string) (BudgetRange, error) {
	var r BudgetRange
	segs := strings.Split(s, "...")
	if len(segs) != 2 {
		return r, errors.New("invalid Budget range")
	}
	r.From = ParseBudget(segs[0])
	r.To = ParseBudget(segs[1])
	return r, nil
}