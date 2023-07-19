package maps

import (
	"log"
	"sort"

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

func (town *TownGenerator) putBuildingOnSide(building *SquareRoom, s side) {
	if building == nil {
		log.Fatalf("putBuildingOnSide() can't use a null building")
		return
	}
	if !building.walls.SizeIsValid() {
		log.Fatalf("putBuildingOnSide() building has invalid size")
		return
	}

	if s == north {
		town.prepBuildingForNorthSide(building)
	}
	if s == south {
		town.prepBuildingForSouthSide(building)
	}
	if s == east {
		town.prepBuildingForEastSide(building)
	}
	if s == west {
		town.prepBuildingForWestSide(building)
	}
	town.buildings = append(town.buildings, building)
}

func (town *TownGenerator) prepBuildingForNorthSide(building *SquareRoom) {
	var widths []int
	var eastWests []int

	if len(town.buildings) < 1 {
		return
	}

	for _, townBuilding := range town.buildings {
		widths = append(widths, townBuilding.walls.Width)
		bottomRight, err := townBuilding.walls.BottomRight()
		if err != nil {
			log.Fatalf("prepBuildingForNorthSide() failed: %s", err.Error())
			return
		}
		eastWests = append(eastWests, townBuilding.walls.TopLeft.X, bottomRight.X)
	}

	westest, eastest := geometry.MinMax(eastWests...)
	westest = westest - building.walls.Width + 1
	eastest = eastest + building.walls.Width - 1

	sort.Ints(widths)
	widest := widths[len(widths)-1]
	if building.walls.Width > widest {
		mid := (eastest + westest) / 2
		building.walls.TopLeft.X = mid - (building.walls.Width / 2)
		return
	}
}

func (town *TownGenerator) prepBuildingForSouthSide(building *SquareRoom) {
}

func (town *TownGenerator) prepBuildingForEastSide(building *SquareRoom) {
}

func (town *TownGenerator) prepBuildingForWestSide(building *SquareRoom) {
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

type side string

const (
	north side = "north"
	south side = "south"
	east  side = "east"
	west  side = "west"
)

var allSides = []side{north, south, east, west}
