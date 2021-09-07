package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackercsv"
	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackerjson"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path) // just pass the file name
	bailOnError(err, "open file")
	defer file.Close()

	normalisedTrackerCsvStruct, err := trackercsv.New(file)

	bailOnError(err, "normalise csv")

	normalisedJsonStructs, err := trackerjson.BulkConvert(normalisedTrackerCsvStruct)

	bailOnError(err, "convert csv to json")

	for i := len(normalisedJsonStructs) - 1; i >= 0; i-- {
		jsonOutput, err := json.Marshal(normalisedJsonStructs[i])

		bailOnError(err, "marshal output")

		fmt.Println(string(jsonOutput))

	}

}

func bailOnError(err error, message string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", message, err)
		os.Exit(1)
	}
}
