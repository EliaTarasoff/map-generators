package geometry

import (
	"errors"
	"math"
)

type AxisAlignedBoundingBox struct {
	TopLeft *Point
	Width   int
	Height  int
}

func (box *AxisAlignedBoundingBox) sizeIsValid() bool {
	if box == nil {
		return false
	}

	return box.Width > 0 && box.Height > 0
}

func (box *AxisAlignedBoundingBox) BottomRight() (*Point, error) {
	if box == nil || box.TopLeft == nil {
		return nil, errors.New("can't compute bottom-right corner on nulls")
	}

	if !box.sizeIsValid() {
		return nil, errors.New("can't compute bottom-right with a negative size")
	}

	return &Point{
		X: box.TopLeft.X + box.Width,
		Y: box.TopLeft.Y + box.Height,
	}, nil
}

func (box *AxisAlignedBoundingBox) Intersection(other *AxisAlignedBoundingBox) (intersections []*Point, err error) {
	if box == nil || other == nil {
		return nil, errors.New("can't do intersection on a null shape")
	}

	if !box.sizeIsValid() || !other.sizeIsValid() {
		return nil, errors.New("can't do intersection on shapes with negative size")
	}

	boxBR, err := box.BottomRight()
	if err != nil {
		return nil, errors.Join(errors.New("can't compute intersection with bad bottom-right"), err)
	}

	otherBR, err := other.BottomRight()
	if err != nil {
		return nil, errors.Join(errors.New("can't compute intersection with bad bottom-right"), err)
	}

	// boxes are completely outside of each other
	if boxBR.X < other.TopLeft.X || box.TopLeft.X > otherBR.X ||
		boxBR.Y < other.TopLeft.Y || box.TopLeft.Y > otherBR.Y {
		return nil, nil
	}

	// totally overlapping (equal) boxes
	if ((box.TopLeft.X == other.TopLeft.X) && (box.TopLeft.Y == other.TopLeft.Y)) &&
		((boxBR.X == otherBR.X) && (boxBR.Y == otherBR.Y)) {
		return []*Point{
			box.TopLeft,
			boxBR,
		}, nil
	}

	// equal width
	if (box.TopLeft.X == other.TopLeft.X) && (boxBR.X == otherBR.X) {
		// overlap one edge
		if box.TopLeft.Y == otherBR.Y {
			return , nil
		}
		if boxBR.Y == other.TopLeft.Y {
			return , nil
		}

		all := []int{box.TopLeft.Y, boxBR.Y, other.TopLeft.Y, otherBR.Y}
		top := min(all...)
		bottom := max(all...)
		return []*Point{
			{
				X: box.TopLeft.X,
				Y: top,
			},
			{
				X: boxBR.X,
				Y: bottom,
			},
		}, nil
	}

	// equal height
	if (box.TopLeft.Y == other.TopLeft.Y) && (boxBR.Y == otherBR.Y) {
		// overlap one edge
		if box.TopLeft.X == otherBR.X {
			return , nil
		}
		if boxBR.X == other.TopLeft.X {
			return , nil
		}

		all := []int{box.TopLeft.X, boxBR.X, other.TopLeft.X, otherBR.X}
		left := min(all...)
		right := max(all...)
		return []*Point{
			{
				X: left,
				Y: box.TopLeft.Y,
			},
			{
				X: right,
				Y: boxBR.Y,
			},
		}, nil
	}

	return nil, errors.New("UNIMPLEMENTED")
}

type Point struct {
	X int
	Y int
}

func min(nums ...int) int {
	small := math.MaxInt
	for _, num := range nums {
		if num < small {
			small = num
		}
	}
	return small
}

func max(nums ...int) int {
	big := -(math.MaxInt - 1)
	for _, num := range nums {
		if num > big {
			big = num
		}
	}
	return big
}
