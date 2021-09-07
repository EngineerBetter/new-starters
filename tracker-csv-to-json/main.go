package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackercsv"
	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackerjson"
)

func main() {
	file := os.Args[1]

	csvData, err := ioutil.ReadFile(file) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	r := bytes.NewReader(csvData)
	normalisedTrackerCsvStruct, err := trackercsv.New(r)

	if err != nil {
		fmt.Println(err)
	}

	normalisedJsonStructs, err := trackerjson.BulkConvert(normalisedTrackerCsvStruct)

	if err != nil {
		fmt.Println(err)
	}

	for _, line := range normalisedJsonStructs {
		jsonOutput, err := json.Marshal(line)

		if err != nil {
			fmt.Println("ERROR OCCURING: ", err)
		}

		os.Stdout.Write(jsonOutput[:])
		fmt.Println("")
	}

}
