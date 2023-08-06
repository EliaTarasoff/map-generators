package maps

import (
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

func NewTownGenerator(random *SaneRandomGenerator, minBuildings, maxBuildings, minBuildingSize, maxBuildingSize int) *TownGenerator {
	return &TownGenerator{
		random:          random,
		minBuildings:    minBuildings,
		maxBuildings:    maxBuildings,
		minBuildingSize: minBuildingSize,
		maxBuildingSize: maxBuildingSize,
	}
}

type TownGenerator struct {
	random          *SaneRandomGenerator
	minBuildings    int
	maxBuildings    int
	minBuildingSize int
	maxBuildingSize int

	buildings []*SquareRoom
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

	sideIndex := town.random.Int(0, 3)
	buildingSide := allSides[sideIndex]
	town.putBuildingOnSide(building, buildingSide)

	town.buildings = append(town.buildings)
}

func (town *TownGenerator) placeBuildingCloseToWater() {

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
	DistanceTo(pos geometry.Point) int
}

type Oasis struct {
	position *geometry.Point
	radius   int
}

func (oasis *Oasis) DistanceTo(pos geometry.Point) int {

}
