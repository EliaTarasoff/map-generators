package main

import (
	"map-generators/maps"
)

const SubBlocksPerBlock = 2

func main() {
	random := maps.NewSaneRandomGeneratorNow()
	townGenerator := maps.NewTownGenerator(random)
	town := townGenerator.Generate()
	maps.PrintThings(town)
}
