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

	t.Run("zero segments", func(t *testing.T) {
		highs := GetHighestValueSegments(nil)
		if !segmentsMatch(highs, nil) {
			failMatch(highs, nil, t)
		}
	})

	t.Run("one segment", func(t *testing.T) {
		ins := []Segment{
			{
				Height: 1,
				Left:   2,
				Right:  3,
			},
		}
		highs := GetHighestValueSegments(ins)
		if !segmentsMatch(highs, ins) {
			failMatch(highs, ins, t)
		}
	})

	t.Run("two overlapping segments", func(t *testing.T) {
		ins := []Segment{
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
		expecteds := []Segment{
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
		highs := GetHighestValueSegments(ins)
		if !segmentsMatch(highs, expecteds) {
			failMatch(highs, expecteds, t)
		}
	})

	t.Run("three overlapping segments", func(t *testing.T) {
		ins := []Segment{
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
		expecteds := []Segment{
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
		highs := GetHighestValueSegments(ins)
		if !segmentsMatch(highs, expecteds) {
			failMatch(highs, expecteds, t)
		}
	})
}
