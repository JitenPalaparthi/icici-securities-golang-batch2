package main

import "fmt"

func main() {

	var map1 map[string]any // declartion but not instantiated
	if map1 == nil {
		map1 = make(map[string]any, 5) // instantiated
	}

	map1["560086"] = "Bangalore-1"
	map1["560096"] = "Bangalore-2"
	map1["560034"] = "Bangalore-3"
	map1["560067"] = "Bangalore-4"
	map1["522001"] = 522002

	for k, v := range map1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	v := map1["5600863"]
	fmt.Println(v)

	map2 := make(map[int]int)
	map2[1] = 0
	map2[2] = 100
	map2[3] = 300

	v1 := map2[1]
	println(v1)
	v2, ok := map2[10]
	if ok {
		println(v2)
	} else {
		println("key does not exist")
	}

	clear(map2)
	fmt.Println(map2) // debug print , it just prints for the developer
	delete(map1, "522001")
	// delete(map[string]any{}, "1") // no op
	// //add([]int{10, 20})

	// //map3 := map[string]any{"1": 1, "2": 2} // short hand declaration
	// //slice := []int{10, 20}

	// var map4 map[string]any
	// delete(map4, "213")
}

// map has a header
// ptr
// size

// What all types can be used as key == operator

// create a function, to delete a key from the map
// if map is nil return error, saying nil map
// if key does not exist return error with a message key does not exist
