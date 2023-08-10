package geometry

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetHighestValueSegments(t *testing.T) {
	failMatch := func(got, expected []Segment) {
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

	var in []Segment
	var expected []Segment
	highest := GetHighestValueSegments(in)
	if !segmentsMatch(highest, expected) {
		failMatch(highest, expected)
	}

	in = []Segment{
		{
			Height: 1,
			Left:   2,
			Right:  3,
		},
	}
	expected = in
	highest = GetHighestValueSegments(in)
	if !segmentsMatch(highest, expected) {
		failMatch(highest, expected)
	}

	in = []Segment{
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
	expected = []Segment{
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
	highest = GetHighestValueSegments(in)
	if !segmentsMatch(highest, expected) {
		failMatch(highest, expected)
	}

	in = []Segment{
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
	expected = []Segment{
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
	highest = GetHighestValueSegments(in)
	if !segmentsMatch(highest, expected) {
		failMatch(highest, expected)
	}
}
