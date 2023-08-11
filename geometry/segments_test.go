package geometry

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetHighestValueSegments(t *testing.T) {
	failMatch := func(got, expecteds []Segment, t *testing.T) {
		var goodsOrBads []string
		for i, expected := range expecteds {
			if i >= len(got) {
				goodsOrBads = append(goodsOrBads, "missing")
				continue
			}
			strGot := spew.Sdump(got[i])
			strExpected := spew.Sdump(expected)
			if strGot == strExpected {
				goodsOrBads = append(goodsOrBads, "good")
				continue
			}
			goodsOrBads = append(goodsOrBads, strGot)
		}
		t.Errorf("mismatch: [%s]", strings.Join(goodsOrBads, ", "))
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

	t.Run("two overlapping segments, out of order", func(t *testing.T) {
		ins := []Segment{
			{
				Height: 9,
				Left:   1,
				Right:  3,
			},
			{
				Height: 8,
				Left:   0,
				Right:  2,
			},
		}
		expecteds := []Segment{
			{
				Height: 8,
				Left:   0,
				Right:  0,
			},
			{
				Height: 9,
				Left:   1,
				Right:  3,
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

	t.Run("several overlapping segments", func(t *testing.T) {
		ins := []Segment{
			{
				Height: 11,
				Left:   5,
				Right:  100,
			},
			{
				Height: 33,
				Left:   95,
				Right:  200,
			},
			{
				Height: 22,
				Left:   195,
				Right:  300,
			},
			{
				Height: 44,
				Left:   295,
				Right:  400,
			},
		}
		expecteds := []Segment{
			{
				Height: 11,
				Left:   5,
				Right:  94,
			},
			{
				Height: 33,
				Left:   95,
				Right:  200,
			},
			{
				Height: 22,
				Left:   201,
				Right:  294,
			},
			{
				Height: 44,
				Left:   295,
				Right:  400,
			},
		}
		highs := GetHighestValueSegments(ins)
		if !segmentsMatch(highs, expecteds) {
			failMatch(highs, expecteds, t)
		}
	})
}
