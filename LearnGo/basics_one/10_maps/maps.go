package main

import (
	"fmt"
	"maps"
)

func main() {
	mapOne := make(map[string]string)
	mapTwo := make(map[string]string)
	mapInt := make(map[string]int)

	mapOne["Name"] = "Golang"
	mapOne["Path"] = "backend"

	fmt.Println(mapOne["Name"], mapOne["Path"])
	fmt.Println(len(mapTwo))

	fmt.Println(mapOne["Emtpy || Zero"])
	fmt.Println(mapInt["Emtpy || Zero"])

	delete(mapOne, "Name")
	fmt.Println(mapOne)

	clear(mapOne)
	fmt.Println(mapOne)

	shortHand := map[string]int{}

	shortHand["Start"] = 10
	fmt.Println(shortHand["Start"])

	value, ok := mapOne["Name"]
	fmt.Println(value)

	if ok {
		fmt.Println("found")
	} else {
		fmt.Println("Nope")
	}

	fmt.Println(maps.Equal(mapOne, mapTwo))
}
