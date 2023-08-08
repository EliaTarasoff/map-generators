package maps

import (
	"errors"
	"math"

	"map-generators/geometry"
)

/*
## three buildings
xxx xxxxx
x | |   x
xxx x   x
    x   x
xxx x   x
x x xxxxx
x-x

## buildings with plants in between
xxx xxxxx
x | |   x
xxx x   x
    x   x
ppp x   x
ppp x   x
    x   x
xxx x   x
x x xxxxx
x-x
*/

func NewTownGenerator(random *SaneRandomGenerator, minBuildings, maxBuildings, minBuildingSize, maxBuildingSize, oasisSize int) *TownGenerator {
	return &TownGenerator{
		random:          random,
		minBuildings:    minBuildings,
		maxBuildings:    maxBuildings,
		minBuildingSize: minBuildingSize,
		maxBuildingSize: maxBuildingSize,

		waterSource: &Oasis{
			position: &geometry.Point{},
			radius:   oasisSize,
		},
	}
}

type TownGenerator struct {
	random          *SaneRandomGenerator
	minBuildings    int
	maxBuildings    int
	minBuildingSize int
	maxBuildingSize int

	buildings   []*SquareRoom
	waterSource WaterSource
}

func (town *TownGenerator) Generate() []MapThing {
	numBuildings := town.random.Int(town.minBuildings, town.maxBuildings)
	for i := 0; i < numBuildings; i++ {
		town.addBuilding()
	}
	var things []MapThing
	for _, room := range town.buildings {
		things = append(things, MapThing(room))
	}
	return things
}

func (town *TownGenerator) addBuilding() {
	building := &SquareRoom{
		walls: &geometry.AxisAlignedBoundingBox{
			Width:  town.random.Int(town.minBuildingSize, town.maxBuildingSize),
			Height: town.random.Int(town.minBuildingSize, town.maxBuildingSize),
		},
	}

	town.chooseBuildingType(building)
	town.buildAroundWater(building)
	town.buildings = append(town.buildings)
}

func (town *TownGenerator) chooseBuildingType(building *SquareRoom) {
	/*
		TODO
			house
			market
			apartment
			temple/church
			graveyard
	*/
}

func (town *TownGenerator) buildAroundWater(building *SquareRoom) {
	building.walls.MoveBottomTo(0)
	for {
		if town.waterSource.DistanceTo(building.walls.) {

		}
	}

	// go clockwise, finding shortest distance to water or other buildings
	// put it there
}

type MapThing interface {
	ToString() string
}

type SquareRoom struct {
	walls *geometry.AxisAlignedBoundingBox
	doors []geometry.Point
}

func (room *SquareRoom) ToString() string {
	return "TODO"
}

type WaterSource interface {
	Width() int
	DistanceTo(pos *geometry.Point) (float64, error)
}

type Oasis struct {
	position *geometry.Point
	radius   int
}

func (oasis *Oasis) DistanceTo(pos *geometry.Point) (float64, error) {
	if oasis == nil || oasis.position == nil || pos == nil {
		return 0, errors.New("can't get distance to nulls")
	}

	dX := math.Abs(float64(oasis.position.X - pos.X))
	dY := math.Abs(float64(oasis.position.Y - pos.Y))
	distanceToCenter := math.Pow((dX*dX)+(dY*dY), 0.5)
	return distanceToCenter - float64(oasis.radius), nil
}

func (oasis *Oasis) Width() int {
	return oasis.radius * 2
}
