package main

import (
	"fmt"
	"sort"
)

func main() {
	// Define a map
	states := make(map[string]string)
	fmt.Println(states)

	// ADD VALUES TO THE MAP
	states["KL"] = "KERALA"
	states["HP"] = "HIMACHAL PRADESH"
	states["JK"] = "KASHMIR"
	states["AS"] = "ASSAM"
	states["TN"] = "TAMILNADU"

	fmt.Println(states)

	// GET A VALUE FROM THE MAP
	kerala := states["KL"]
	fmt.Println(kerala)

	// DELETE A VALUE FROM THE MAP
	delete(states, "AS")
	fmt.Println(states)

	//SORTING OF A MAP BASED ON KEYS

	// CREATE A SLICE and copy all the keys of the map to the slice
	keys := make([]string, len(states))
	i := 0

	// ITERATE OVER A MAP
	for k := range states {
		keys[i] = k
		i++
	}

	// SORT THE keys array
	sort.Strings(keys)
	fmt.Println(keys)

	// PRINTING THE SORTED MAP USING SORTED KEY LIST
	for x := range keys {
		fmt.Println(states[keys[x]])
	}

}
