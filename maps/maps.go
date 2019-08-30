package maps

// GetMap returns a map
func GetMap() map[string]int {
	mapTest := make(map[string]int)
	mapTest["victor"] = 26
	mapTest["otro"] = 2
	mapTest["asfasf"] = 3
	delete(mapTest, "otro")

	return mapTest
}

// GetAgeFromName returns the age of the name
func GetAgeFromName(name string) int {
	mapTest := map[string]int{
		"victor":    26,
		"antonio":   64,
		"antoñito":  101,
		"guadalupe": 2,
	}

	return mapTest[name]
}
