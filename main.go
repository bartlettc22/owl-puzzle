package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

var iterations int

func main() {

	// Start out at position -1 because we want the first piece added to be at 0
	board := Board{CurrentPosition: -1}

	dataPath := "./data.json"
	jsonFile, err := ioutil.ReadFile(dataPath)
	if err != nil {
		log.Fatalf("Error opening %s", dataPath)
	}

	err = json.Unmarshal(jsonFile, &board.Squares)
	if err != nil {
		log.Fatalf("Error unmarshalling %s", dataPath)
	}

	start := time.Now()
	board.solve()
	board.Print()
	duration := time.Since(start)

	fmt.Printf("Solution took %dms and %d iterations\n", duration.Milliseconds(), iterations)

}
