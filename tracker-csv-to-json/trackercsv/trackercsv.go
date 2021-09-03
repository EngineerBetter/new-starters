package trackercsv

import (
	"encoding/csv"
	"io"
)

type NormalisedTrackerCSV struct {
	Title       string
	Type        string
	Description string
	Labels      string
	Tasks       string
}

func New(r io.Reader) ([]NormalisedTrackerCSV, error) {
	content, err := csv.NewReader(r).ReadAll()

	if err != nil {
		return nil, err
	}

	output := []NormalisedTrackerCSV{}
	output = append(output, NormalisedTrackerCSV{
		Title:       content[1][0],
		Type:        content[1][1],
		Description: content[1][2],
		Labels:      content[1][3],
	})

	return output, nil
}
