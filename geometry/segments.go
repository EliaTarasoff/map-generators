package geometry

import "sort"

type Segment struct {
	Left   int
	Right  int
	Height int
}

// GetHighestValueSegments returns an array of segments, sorted from left to right,
// where the height is the highest value from any overlapping segments for that range.
// (Overlapping portions are removed, and new non-overlapping segments returned.
// Like a silhouette of a city sky-line - only the highest shadows are in the silhouette.)
func GetHighestValueSegments(segments []Segment) []Segment {
	switch len(segments) {
	case 0:
		return nil
	case 1:
		return segments
	case 2:
		return getHighestSegments(segments[0], segments[1])
	}

	sortedInputs := make([]Segment, len(segments))
	for i := range segments {
		sortedInputs[i] = segments[i]
	}
	sort.SliceStable(sortedInputs, func(i, j int) bool {
		return sortedInputs[i].Left < sortedInputs[j].Left
	})

	var output []Segment
	for _, segment := range sorted {
	}
}

func CopySlice[K any](in []K) []K {
	out := make([]K, len(in))
	for i := range in {
		out[i] = in[i]
	}
	return out
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
		right.Left = left.Right + 1
	} else {
		left.Right = right.Left - 1
	}

	if left.Right <= left.Left {
		return []Segment{right}
	}
	if right.Left >= right.Right {
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
