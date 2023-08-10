package geometry

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetHighestValueSegments(t *testing.T) {
	failMatch := func(got, expected []Segment, t *testing.T) {
		t.Error(spew.Sprintf("got %v, expected %v", got, expected))
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

	t.Run("one low segment totally overlaps a small high segment", func(t *testing.T) {
		ins := []Segment{
			{
				Height: 5,
				Left:   4,
				Right:  6,
			},
			{
				Height: 3,
				Left:   0,
				Right:  10,
			},
		}
		expecteds := []Segment{
			{
				Height: 3,
				Left:   0,
				Right:  3,
			},
			{
				Height: 5,
				Left:   4,
				Right:  6,
			},
			{
				Height: 3,
				Left:   7,
				Right:  10,
			},
		}
		highs := GetHighestValueSegments(ins)
		if !segmentsMatch(highs, expecteds) {
			failMatch(highs, expecteds, t)
		}
	})

	t.Run("length-1 segment is under and touching longer segment", func(t *testing.T) {
		ins := []Segment{
			{
				Height: 0,
				Left:   0,
				Right:  0,
			},
			{
				Height: 99,
				Left:   0,
				Right:  99,
			},
		}
		expecteds := []Segment{
			{
				Height: 99,
				Left:   0,
				Right:  99,
			},
		}
		highs := GetHighestValueSegments(ins)
		if !segmentsMatch(highs, expecteds) {
			failMatch(highs, expecteds, t)
		}
	})

	t.Run("length-1 segment is over and touching longer segment", func(t *testing.T) {
		ins := []Segment{
			{
				Height: 42,
				Left:   0,
				Right:  0,
			},
			{
				Height: 3,
				Left:   0,
				Right:  99,
			},
		}
		expecteds := []Segment{
			{
				Height: 42,
				Left:   0,
				Right:  0,
			},
			{
				Height: 3,
				Left:   1,
				Right:  99,
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
