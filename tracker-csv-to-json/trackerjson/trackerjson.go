package trackerjson

import (
	"encoding/json"
)

type NormalisedTrackerCSV struct {
	Title       string `json:"name"`
	Type        string `json:"story_type"`
	Description string `json:"description"`
	Labels      string `json:"labels"`
	Tasks       string `json:"tasks"`
}

func Convert(trackerInput []NormalisedTrackerCSV) ([]byte, error) {
	var output []byte

	for _, list := range trackerInput {
		b, _ := json.Marshal(list)
		output = append(output, b...)

	}
	return output, nil
}
