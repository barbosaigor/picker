package picker

import (
	"testing"
)

func TestCreateInterval(t *testing.T) {
	i := createInterval("_", 10, 20)
	if i.min != 10 || i.max != 20 {
		t.Error("Fail to create interval")
	}
}

func TestAdd(t *testing.T) {
	c := New()
	c.Add("a", 10)
	c.Add("b", 20)
	c.Add("c", 30)
	c.Add("d", 40)
	if c.sum != 100 {
		t.Error("Fail to calculate sum")
	}
	if len(c.intervals) != 4 {
		t.Error("Fail to append intervals")
	}
}

func TestRollDice(t *testing.T) {
	c := New()
	c.Add("a", 10)
	c.Add("b", 20)
	c.Add("c", 30)
	c.Add("d", 40)
	for i := 0; i < 500; i++ {
		itvl, err := c.RollDice()
		if err != nil {
			t.Error("Something goes wrong" + err.Error())
		}
		if itvl == "" {
			t.Error("Something goes wrong" + err.Error())
		}
	}
}

func TestFree(t *testing.T) {
	c := New()
	for i := 0; i < 500; i++ {
		c.Add("a", 10)
		c.Add("b", 20)
		c.Add("c", 30)
		c.Add("d", 40)
		c.Reset()
		_, err := c.RollDice()
		if err == nil {
			t.Error("Picker should return an error" + err.Error())
		}
	}
}