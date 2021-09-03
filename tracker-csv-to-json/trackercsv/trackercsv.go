package trackercsv

import "io"

type NormalisedTrackerCSV struct {
	Title       string
	Type        string
	Description string
	Labels      string
	Tasks       string
}

func New(r io.Reader) NormalisedTrackerCSV {
	return NormalisedTrackerCSV{}
}
