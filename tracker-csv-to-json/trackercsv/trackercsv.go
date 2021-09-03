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

	for _, line := range content[1:] {
		if len(line) != 5 {
			output = append(output, NormalisedTrackerCSV{
				Title:       line[0],
				Type:        line[1],
				Description: line[2],
				Labels:      line[3],
			})
		} else {
			output = append(output, NormalisedTrackerCSV{
				Title:       line[0],
				Type:        line[1],
				Description: line[2],
				Labels:      line[3],
				Tasks:       "[" + line[4] + "]",
			})
		}
	}

	return output, nil
}
