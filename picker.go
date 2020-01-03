// picker is a packaged that randomly
// chooses an element based on its
// weight.
package picker

import (
	"math/rand"
	"time"
	"errors"
)

// interval is used for 
// store a range element
type interval struct {
	// min, max is an interval
	min, max int
	// id identifier
	id string
}

type container struct {
	intervals []interval
	// Sum of intervals value
	sum int
}

func New() *container {
	return new(container)
}

func createInterval(id string, min, max int) interval {
	return interval{min, max, id}
}

// randNum Take a random number between 0 and maxNumber
func randNum(maxNumber int) int {
	randSource := rand.NewSource(time.Now().UnixNano())
	randEngine := rand.New(randSource)
	randNumber := randEngine.Intn(maxNumber)
	return randNumber
}

// Add adds an element that has
// a range in the choosen scope
func (c *container) Add(id string, value int) {
	itrvl := createInterval(id, c.sum, c.sum + value - 1)
	c.intervals = append(c.intervals, itrvl)
	c.sum += value
}

// RollDice picks an element 
// that satisfies some range
func (c container) RollDice() (string, error) {
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

// Reset resets picker
// variables
func (c *container) Reset() {
	c.intervals = nil
	c.sum = 0
}