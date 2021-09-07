package trackerjson

import "github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackercsv"

type NormalisedTrackerJSON struct {
	Name        string   `json:"name"`
	Story_type  string   `json:"story_type"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
	Tasks       []string `json:"tasks"`
}

func ConvertCSVStructToJSONSTruct(trackerInput trackercsv.NormalisedTrackerCSV) (NormalisedTrackerJSON, error) {
	var output NormalisedTrackerJSON

	output.Name = trackerInput.Title
	output.Story_type = trackerInput.Type
	output.Description = trackerInput.Description
	output.Labels = trackerInput.Labels
	output.Tasks = trackerInput.Tasks

	return output, nil
}

func BulkConvert(trackerItems []trackercsv.NormalisedTrackerCSV) ([]NormalisedTrackerJSON, error) {
	var output []NormalisedTrackerJSON

	for _, item := range trackerItems {
		convertedItem, err := ConvertCSVStructToJSONSTruct(item)
		if err != nil {
			return nil, err
		}
		output = append(output, convertedItem)
	}
	return output, nil
}
