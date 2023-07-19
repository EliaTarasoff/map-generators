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
	buildings       []*SquareRoom
}

func (town *TownGenerator) Generate() []MapThing {
	var things []MapThing
	return things
}

func (town *TownGenerator) addBuilding() {
	town.buildings = append(town.buildings, &SquareRoom{
		walls: &geometry.AxisAlignedBoundingBox{
			Width:  town.random.Int(town.minBuildingSize, town.maxBuildingSize),
			Height: town.random.Int(town.minBuildingSize, town.maxBuildingSize),
		},
	})
}

type MapThing interface {
	ToString() string
}

type SquareRoom struct {
	walls *geometry.AxisAlignedBoundingBox
	doors []*geometry.Point
}

func (room *SquareRoom) ToString() string {
	return "TODO"
}
