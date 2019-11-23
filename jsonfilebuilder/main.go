package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("example.json")
	if err != nil {
		fmt.Print("\n error while opeing file.. Error is", err)
	}
	exampleMap := make(map[int]string, 0)

	for i := 0; i < 100; i++ {

		exampleMap[i] = strconv.Itoa(rand.Intn(200))
	}

	json.NewEncoder(file).Encode(&exampleMap)

}
