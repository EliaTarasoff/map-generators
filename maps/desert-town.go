package maps

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

func GenerateDesertTown() []MapThing {
	var town []MapThing
}

func addBuildingToTown(town []MapThing) []MapThing {
	room := getRoom()
	if len(town) == 0 {
		return []MapThing{room}
	}

}

func getRoom() MapThing {}

type MapThing interface{}
