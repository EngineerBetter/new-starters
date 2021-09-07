package trackercsv

import (
	"encoding/csv"
	"io"
	"strings"
)

type NormalisedTrackerCSV struct {
	Title       string
	Type        string
	Description string
	Labels      []string
	Tasks       []string
}

func New(r io.Reader) ([]NormalisedTrackerCSV, error) {
	content, err := csv.NewReader(r).ReadAll()

	if err != nil {
		return nil, err
	}

	output := []NormalisedTrackerCSV{}

	for _, line := range content[1:] {
		if len(line) < 5 {
			output = append(output, NormalisedTrackerCSV{
				Title:       line[0],
				Type:        line[1],
				Description: line[2],
				Labels:      strings.Split(line[3], ","),
				Tasks:       []string{},
			})
		} else {
			var task_list []string

			for _, task := range line[4:] {
				if task != "" {
					task_list = append(task_list, task)
				}
			}

			output = append(output, NormalisedTrackerCSV{
				Title:       line[0],
				Type:        line[1],
				Description: line[2],
				Labels:      strings.Split(line[3], ","),
				Tasks:       task_list,
			})

		}
	}

	return output, nil
}
