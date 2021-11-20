package wanderlust_test
import (
	"github.com/cheekybits/is"
	wanderlust "github.com/mainak90/wanderlust"
	"testing"
)
func TestBudgetValues(t *testing.T) {
	is := is.New(t)
	is.Equal(int(wanderlust.Budget1), 1)
	is.Equal(int(wanderlust.Budget2), 2)
	is.Equal(int(wanderlust.Budget3), 3)
	is.Equal(int(wanderlust.Budget4), 4)
	is.Equal(int(wanderlust.Budget5), 5)
}

func TestBudgetString(t *testing.T) {
	is := is.New(t)
	is.Equal(wanderlust.Budget1.String(), "$")
	is.Equal(wanderlust.Budget2.String(), "$$")
	is.Equal(wanderlust.Budget3.String(), "$$$")
	is.Equal(wanderlust.Budget4.String(), "$$$$")
	is.Equal(wanderlust.Budget5.String(), "$$$$$")
}

func TestParseBudget(t *testing.T) {
	is := is.New(t)
	is.Equal(wanderlust.ParseBudget("$"), wanderlust.Budget1)
	is.Equal(wanderlust.ParseBudget("$$"), wanderlust.Budget2)
	is.Equal(wanderlust.ParseBudget("$$$"), wanderlust.Budget3)
	is.Equal(wanderlust.ParseBudget("$$$$"), wanderlust.Budget4)
	is.Equal(wanderlust.ParseBudget("$$$$$"), wanderlust.Budget5)
}

func TestParseBudgetRange(t *testing.T) {
	is := is.New(t)
	var l wanderlust.BudgetRange
	l, _ = wanderlust.ParseBudgetRange("$$...$$$")
	is.Equal(l.From, wanderlust.Budget2)
	is.Equal(l.To, wanderlust.Budget3)
}
func TestCostRangeString(t *testing.T) {
	is := is.New(t)
	is.Equal("$$...$$$$", (&wanderlust.BudgetRange{
		From: wanderlust.Budget2,
		To: wanderlust.Budget4,
	}).String())
}