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
