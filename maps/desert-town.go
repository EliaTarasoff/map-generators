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

func NewTownGenerator(random *SaneRandomGenerator) *TownGenerator {
	return &TownGenerator{
		random: random,
	}
}

type TownGenerator struct {
	random *SaneRandomGenerator
}

func (town *TownGenerator) Generate() []MapThing {
	var things []MapThing
}

func (town *TownGenerator) addBuildingToTown(things []MapThing) []MapThing {
	room := town.getRoom()
	if len(things) == 0 {
		return []MapThing{room}
	}
}

func (town *TownGenerator) getRoom(minSize, maxSize int) geometry.AxisAlignedBoundingBox {
	return geometry.AxisAlignedBoundingBox{
		Width:  town.random.Int(minSize, maxSize),
		Height: town.random.Int(minSize, maxSize),
	}
}

type MapThing interface{}
