package geometry

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetHighestValueSegments(t *testing.T) {
	failMatch := func(got, expected []Segment, t *testing.T) {
		t.Error(spew.Sprintf("got high-segments %v, expected %v", got, expected))
	}
	segmentsMatch := func(ins, expecteds []Segment) bool {
		if len(ins) != len(expecteds) {
			return false
		}
		for i, expected := range expecteds {
			in := ins[i]
			if in.Height != expected.Height ||
				in.Left != expected.Left ||
				in.Right != expected.Right {
				return false
			}
		}
		return true
	}
	getAndCheck := func(ins, expecteds []Segment, t *testing.T) {
		highs := GetHighestValueSegments(ins)
		if !segmentsMatch(highs, expecteds) {
			failMatch(highs, expecteds, t)
		}
	}

	t.Run("zero segments", func(t *testing.T) {
		getAndCheck(nil, nil, t)
	})

	t.Run("one segment", func(t *testing.T) {
		in := []Segment{
			{
				Height: 1,
				Left:   2,
				Right:  3,
			},
		}
		getAndCheck(in, in, t)
	})

	t.Run("two overlapping segments", func(t *testing.T) {
		in := []Segment{
			{
				Height: 7,
				Left:   5,
				Right:  10,
			},
			{
				Height: 3,
				Left:   0,
				Right:  8,
			},
		}
		expected := []Segment{
			{
				Height: 3,
				Left:   0,
				Right:  4,
			},
			{
				Height: 7,
				Left:   5,
				Right:  10,
			},
		}
		getAndCheck(in, expected, t)
	})

	t.Run("three overlapping segments", func(t *testing.T) {
		in := []Segment{
			{
				Height: 7,
				Left:   5,
				Right:  10,
			},
			{
				Height: 3,
				Left:   0,
				Right:  8,
			},
			{
				Height: 2,
				Left:   -1,
				Right:  12,
			},
		}
		expected := []Segment{
			{
				Height: 2,
				Left:   -1,
				Right:  -1,
			},
			{
				Height: 3,
				Left:   0,
				Right:  4,
			},
			{
				Height: 7,
				Left:   5,
				Right:  10,
			},
			{
				Height: 2,
				Left:   11,
				Right:  12,
			},
		}
		getAndCheck(in, expected, t)
	})
}
