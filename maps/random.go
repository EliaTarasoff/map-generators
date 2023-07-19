package maps

import (
	"math"
	"math/rand"
	"time"
)

func NewSaneRandomGeneratorNow() *SaneRandomGenerator {
	random := rand.NewSource(time.Now().Unix())
	return NewSaneRandomGenerator(random)
}

func NewSaneRandomGenerator(random rand.Source) *SaneRandomGenerator {
	return &SaneRandomGenerator{
		random: random,
	}
}

type SaneRandomGenerator struct {
	random rand.Source
}

func (random *SaneRandomGenerator) Int(min, max int) int {
	raw := random.random.Int63()
	floated := float64(raw) / float64(math.MaxInt64)
	diff := float64(max - min)
	return int(floated*diff) + min
}
