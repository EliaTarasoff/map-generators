package main

import (
	"map-generators/maps"
)

func main() {
	random := maps.NewSaneRandomGeneratorNow()
	townGenerator := maps.NewTownGenerator(random, 1, 5, 3, 10, 5)
	town := townGenerator.Generate()
	maps.PrintThings("town", town)
}
