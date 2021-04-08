
package FuzzyBool

import "fmt"
type FuzzyBool float32
//func New(value interface{}) (*FuzzyBool, error) {
//	amount, err := float32ForValue(value)
//	return &FuzzyBool{amount}, err
//}
func master() {


	a, _ := New(0)
	b, _ := New(.25)
	c, _ := New(.75)
	d := c.Copy()
	d, _ = New(1)
	process(a, b, c, d)
	s := []FuzzyBool{a, b, c, d}
	fmt.Println(s)
}

func process(a, b, c, d FuzzyBool) {
	fmt.Println("Original:", a, b, c, d)
	fmt.Println("Not:     ", a.Not(), b.Not(), c.Not(), d.Not())
	fmt.Println("Not Not: ", a.Not().Not(), b.Not().Not(), c.Not().Not(),
		d.Not().Not())
	fmt.Print("0.And(.25)→", a.And(b), "  .25.And(.75)→", b.And(c),
		"  .75.And(1)→", c.And(d), "  0.And(.25,.75,1)→", a.And(b, c, d),
		"\n")
	fmt.Print("0.Or(.25)→", a.Or(b), "  .25.Or(.75)→", b.Or(c),
		"  .75.Or(1)→", c.Or(d), "  0.Or(.25,.75,1)→", a.Or(b, c, d), "\n")
	fmt.Println("a < c, a == c, a > c:", a.Less(c), a.Equal(c), c.Less(a))
	fmt.Println("Bool:    ", a.Bool(), b.Bool(), c.Bool(), d.Bool())
	fmt.Println("Float:   ", a.Float(), b.Float(), c.Float(), d.Float())
}

func New(value interface{}) (FuzzyBool, error) {
	var fuzzy float32
	switch value := value.(type) { // shadow variable
	case float32:
		fuzzy = value
	case float64:
		fuzzy = float32(value)
	case int:
		fuzzy = float32(value)
	case bool:
		fuzzy = 0
		if value {
			fuzzy = 1
		}
	default:
		return FuzzyBool(0), fmt.Errorf("fuzzybool.New(): %v is not a " +
			"number or boolean\n", value)
	}
	if fuzzy < 0 {
		fuzzy = 0
	} else if fuzzy > 1 {
		fuzzy = 1
	}
	return FuzzyBool(fuzzy), nil
}

func (fuzzy FuzzyBool) Copy() FuzzyBool {
	return FuzzyBool(fuzzy)
}

func (fuzzy FuzzyBool) String() string {
	return fmt.Sprintf("%.0f%%", 100*float32(fuzzy))
}

func (fuzzy FuzzyBool) Not() FuzzyBool {
	return FuzzyBool(1 - float32(fuzzy))
}

func (fuzzy FuzzyBool) And(first FuzzyBool,
	rest ...FuzzyBool) FuzzyBool {
	minimum := fuzzy
	rest = append(rest, first)
	for _, other := range rest {
		if minimum > other {
			minimum = other
		}
	}
	return FuzzyBool(minimum)
}

func (fuzzy FuzzyBool) Or(first FuzzyBool,
	rest ...FuzzyBool) FuzzyBool {
	maximum := fuzzy
	rest = append(rest, first)
	for _, other := range rest {
		if maximum < other {
			maximum = other
		}
	}
	return FuzzyBool(maximum)
}

func (fuzzy FuzzyBool) Less(other FuzzyBool) bool {
	return fuzzy < other
}

func (fuzzy FuzzyBool) Equal(other FuzzyBool) bool {
	return fuzzy == other
}

func (fuzzy FuzzyBool) Bool() bool {
	return float32(fuzzy) >= .5
}

func (fuzzy FuzzyBool) Float() float64 {
	return float64(fuzzy)
}
