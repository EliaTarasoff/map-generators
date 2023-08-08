package geometry

import "sort"

type Segment struct {
	Left   int
	Right  int
	Height int
}

// GetHighestValueSegments returns an array of segments, sorted from left to right,
// where the height is the highest of all input segments dor that range.
func GetHighestValueSegments(segments []Segment) []Segment {
	sorted := make([]Segment, len(segments))
	for i := range segments {
		sorted[i] = segments[i]
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Left < sorted[j].Left
	})

	var output []Segment
	for _, segment := range sorted {
	}
}

func getHighestSegments(a, b Segment) []Segment {
	if a.Right < b.Left {
		return []Segment{a, b}
	}
	if b.Right < a.Left {
		return []Segment{b, a}
	}

	left := Segment{}
	right := Segment{}

	if a.Left < b.Left {
		left = copySegment(a)
	} else {
		left = copySegment(b)
	}

	if a.Right > b.Right {
		right = copySegment(a)
	} else {
		right = copySegment(b)
	}

	if left.Height > right.Height {
		right.Left += 1
	} else {
		left.Right -= 1
	}

	if left.Left == left.Right {
		return []Segment{right}
	}
	if right.Left == right.Right {
		return []Segment{left}
	}
	return []Segment{left, right}
}

func copySegment(a Segment) Segment {
	return Segment{
		Left:   a.Left,
		Right:  a.Right,
		Height: a.Height,
	}
}
