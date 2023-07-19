package geometry

import (
	"errors"
	"math"
	"sort"
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

	xTouches := lineTouchLine(box.TopLeft.X, boxBR.X, other.TopLeft.X, otherBR.X)
	yTouches := lineTouchLine(box.TopLeft.Y, boxBR.Y, other.TopLeft.Y, otherBR.Y)

	// totally outside
	if len(xTouches) == 0 || len(yTouches) == 0 {
		return nil, nil
	}

	// broken line-intersection...
	if len(xTouches) > 2 || len(yTouches) > 2 {
		return nil, errors.New("somehow got more than two points describing line-intersections")
	}

	// only touching on one side
	if len(xTouches) == 1 {
		return []*Point{
			{
				X: xTouches[0],
				Y: yTouches[0],
			},
			{
				X: xTouches[0],
				Y: yTouches[1],
			},
		}, nil
	}
	if len(yTouches) == 1 {
		return []*Point{
			{
				X: xTouches[0],
				Y: yTouches[0],
			},
			{
				X: xTouches[1],
				Y: yTouches[0],
			},
		}, nil
	}

	// overlapping on corner-chunks, or side-chunks
	return []*Point{
		{
			X: xTouches[0],
			Y: yTouches[0],
		},
		{
			X: xTouches[1],
			Y: yTouches[1],
		},
	}, nil
}

func lineTouchLine(a1, a2, b1, b2 int) []int {
	smallA, bigA := a1, a2
	if a2 < a1 {
		smallA, bigA = a2, a1
	}
	smallB, bigB := b1, b2
	if b2 < b1 {
		smallB, bigB = b2, b1
	}

	// totally outside
	if bigA < smallB || smallA > bigB {
		return nil
	}

	// total overlap
	if smallA == smallB && bigA == bigB {
		return []int{smallA, bigA}
	}

	smallATouchB := pointTouchLine(smallA, smallB, bigB)
	bigATouchB := pointTouchLine(bigA, smallB, bigB)
	all := append(smallATouchB, bigATouchB...)
	sort.Ints(all)

	// only intersect at one point
	if all[0] == all[1] {
		return []int{all[0]}
	}

	// one line is smaller and inside, or they're each partially touching
	return all
}

func pointTouchLine(a, b1, b2 int) []int {
	leftB, rightB := b1, b2
	if b2 < b1 {
		leftB, rightB = b2, b1
	}

	if a < leftB || a > rightB {
		return nil
	}
	if a == leftB {
		return []int{leftB}
	}
	if a == rightB {
		return []int{rightB}
	}
	return []int{a}
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
