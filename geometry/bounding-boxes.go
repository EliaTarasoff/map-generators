package geometry

import (
	"errors"
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

	// boxes are equal
	if ((box.TopLeft.X == other.TopLeft.X) && (box.TopLeft.Y == other.TopLeft.Y)) &&
		((boxBR.X == otherBR.X) && (boxBR.Y == otherBR.Y)) {
		return []*Point{
			box.TopLeft,
			boxBR,
		}, nil
	}

	return nil, errors.New("UNIMPLEMENTED")
}

type Point struct {
	X int
	Y int
}
