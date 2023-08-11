package geometry

import (
	"math"
)

type Segment struct {
	Left   int
	Right  int
	Height int
}

func (s *Segment) Copy() *Segment {
	if s == nil {
		return nil
	}

	return &Segment{
		Left:   s.Left,
		Right:  s.Right,
		Height: s.Height,
	}
}

func (s *Segment) Size() int {
	if s.Right < s.Left || s.Left > s.Right {
		return -1
	}
	return s.Right - s.Left + 1
}

func (s *Segment) ShrinkLeft() *Segment {
	if s.Size() == 1 {
		return nil
	}
	s.Left += 1
	return s
}

func (s *Segment) ShrinkRight() *Segment {
	if s.Size() == 1 {
		return nil
	}
	s.Right -= 1
	return s
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

	numSeg := len(segments)
	half := numSeg / 2
	left := make([]Segment, half)
	right := make([]Segment, half)

	for i, segment := range segments {
		if i < half {
			left = append(left, segment)
		} else {
			right = append(right, segment)
		}
	}
	unmergedLeft := GetHighestValueSegments(left)
	unmergedRight := GetHighestValueSegments(right)
	return mergeSkylines(unmergedLeft, unmergedRight)
}

func mergeSkylines(leftSkyline, rightSkyline []Segment) []Segment {
	numLeft := len(leftSkyline)
	numRight := len(rightSkyline)
	if numLeft == 0 && numRight == 0 {
		return nil
	}
	if numLeft == 0 {
		return rightSkyline
	}
	if numRight == 0 {
		return leftSkyline
	}

	type tempSkylinePos struct {
		valid  bool
		pos    int
		height int
	}
	left := math.MaxInt
	right := math.MinInt
	for _, building := range leftSkyline {
		if building.Left < left {
			left = building.Left
		}
	}
	for _, building := range rightSkyline {
		if building.Right < right {
			right = building.Right
		}
	}
	var tempSkylinePoss []*tempSkylinePos
	for p := left; p <= right; p++ {
		tempSkylinePoss = append(tempSkylinePoss, &tempSkylinePos{
			pos: p,
		})
	}

	posToArray := func(pos int) int {
		return pos - left
	}
	for _, building := range leftSkyline {
		for pos := building.Left; pos <= building.Right; pos++ {
			temp := tempSkylinePoss[posToArray(pos)]
			temp.valid = true
			temp.height = building.Height
		}
	}
	for _, building := range rightSkyline {
		for pos := building.Left; pos <= building.Right; pos++ {
			temp := tempSkylinePoss[posToArray(pos)]
			temp.valid = true
			temp.height = building.Height
		}
	}

	type tempSkylineSegment struct {
		hasLeft bool
		seg     Segment
	}
	tempSegment := tempSkylineSegment{}
	tempSkylineSegments := make([]tempSkylineSegment, 0)
	for _, temp := range tempSkylinePoss {
		if !tempSegment.hasLeft {
			tempSegment.seg.Left = temp.pos
			tempSegment.seg.Height = temp.height
			tempSegment.hasLeft = true
		}
		if temp.valid {
			tempSegment.seg.Right = temp.pos
		} else {
			tempSkylineSegments = append(tempSkylineSegments, tempSegment)
		}
	}

	outs := make([]Segment, len(tempSkylineSegments))
	for i, segment := range tempSkylineSegments {
		outs[i] = segment.seg
	}
	return outs
}

func getHighestSegments(a, b Segment) []Segment {
	// no overlap at all
	if a.Right < b.Left {
		return []Segment{a, b}
	}
	if b.Right < a.Left {
		return []Segment{b, a}
	}

	lower := a
	higher := b
	if a.Height > b.Height {
		lower = b
		higher = a
	}

	// total overlap
	if a.Left == b.Left && a.Right == b.Right {
		return []Segment{higher}
	}

	bounds := Segment{}
	if a.Left < b.Left {
		bounds.Left = a.Left
	} else {
		bounds.Left = b.Left
	}
	if a.Right > b.Right {
		bounds.Right = a.Right
	} else {
		bounds.Right = b.Right
	}

	// equal height
	if a.Height == b.Height {
		bounds.Height = a.Height
		return []Segment{bounds}
	}

	// one segment totally inside the other
	aInside := a.Left > bounds.Left && a.Right < bounds.Right
	bInside := b.Left > bounds.Left && b.Right < bounds.Right
	outer := Segment{}
	if aInside {
		// totally overshadowed
		if a.Height <= b.Height {
			return []Segment{b}
		}
		outer = b
	}
	if bInside {
		// totally overshadowed
		if b.Height <= a.Height {
			return []Segment{b}
		}
		outer = a
	}
	if aInside || bInside {
		return []Segment{
			{
				Height: lower.Height,
				Left:   outer.Left,
				Right:  higher.Left - 1,
			},
			{
				Height: higher.Height,
				Left:   higher.Left,
				Right:  higher.Right,
			},
			{
				Height: lower.Height,
				Left:   higher.Right + 1,
				Right:  outer.Right,
			},
		}
	}

	pickLeftRight := func(a, b Segment) (left Segment, right Segment) {
		left = a
		right = b
		if b.Left < a.Left {
			left = b
			right = a
		}
		return
	}

	// barely overlapping on one edge
	if a.Right == b.Left || a.Left == b.Right {
		left, right := pickLeftRight(a, b)

		if left.Height > right.Height {
			if right.Size() == 1 {
				return []Segment{left}
			}
			return []Segment{
				left,
				*right.Copy().ShrinkLeft(),
			}
		} else if right.Height > left.Height {
			if left.Size() == 1 {
				return []Segment{right}
			}
			return []Segment{
				*left.Copy().ShrinkRight(),
				right,
			}
		}
	}

	// overlapping and share one edge
	if a.Right == b.Right || b.Left == a.Left {
		if a.Size() == 1 && b.Height > a.Height {
			return []Segment{b}
		}
		if b.Size() == 1 && a.Height > b.Height {
			return []Segment{a}
		}
		left, right := pickLeftRight(a, b)

		if left.Height > right.Height {
			return []Segment{
				left,
				*right.Copy().ShrinkLeft(),
			}
		}
		return []Segment{
			*left.Copy().ShrinkRight(),
			right,
		}
	}

	left, right := pickLeftRight(a, b)
	if left.Height < right.Height {
		left.Right = right.Left - 1
	} else if right.Height < left.Height {
		right.Left = left.Right + 1
	}
	return []Segment{left, right}
}
