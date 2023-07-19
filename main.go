package main

import (
	"map-generators/maps"
)

const SubBlocksPerBlock = 2

func main() {
	random := maps.NewSaneRandomGeneratorNow()
	townGenerator := maps.NewTownGenerator(random, 1, 5, 3, 10)
	town := townGenerator.Generate()
	maps.PrintThings("town", town)
}
