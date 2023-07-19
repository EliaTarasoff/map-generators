package maps

func PrintThings(name string, things []MapThing) {
	println(name)
	for _, thing := range things {
		// TODO this is actually wrong, since they could be boxes, not just lines or points
		println("    ", thing.ToString())
	}
}
