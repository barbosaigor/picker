package picker

import (
	"errors"
	"math/rand"
	"time"
)

// interval is used for
// store a range element
type interval struct {
	// min, max is an interval
	min, max uint32
	// id identifier
	id string
}

// Picker manages elements and implements the selection algorithm
type Picker struct {
	intervals []interval
	// Sum of intervals value
	sum uint32
}

// New creates a picker
func New() *Picker {
	return new(Picker)
}

func createInterval(id string, min, max uint32) interval {
	return interval{min, max, id}
}

// randNum Take a random number between 0 and maxNumber
func randNum(maxNumber uint32) uint32 {
	randSource := rand.NewSource(time.Now().UnixNano())
	randEngine := rand.New(randSource)
	randNumber := randEngine.Intn(int(maxNumber))
	return uint32(randNumber)
}

// Add attaches an element
func (c *Picker) Add(id string, value uint32) {
	itrvl := createInterval(id, c.sum, c.sum+value-1)
	c.intervals = append(c.intervals, itrvl)
	c.sum += value
}

// RollDice selects a random element taking into account each element probability
func (c Picker) RollDice() (string, error) {
	if c.sum == 0 {
		return "", errors.New("No interval attached")
	}
	randNumber := randNum(c.sum)
	for _, itrvl := range c.intervals {
		if itrvl.min <= randNumber && itrvl.max >= randNumber {
			return itrvl.id, nil
		}
	}
	return "", errors.New("No interval found")
}

// Reset cleans picker variables
func (c *Picker) Reset() {
	c.intervals = nil
	c.sum = 0
}
