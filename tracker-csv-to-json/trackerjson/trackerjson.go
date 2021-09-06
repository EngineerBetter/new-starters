package trackerjson

type NormalisedTrackerCSV struct {
	Title       string
	Type        string
	Description string
	Labels      []string
	Tasks       []string
}

type NormalisedTrackerJSON struct {
	Name        string
	Story_type  string
	Description string
	Labels      []string
	Tasks       []string
}

func ConvertCSVStructToJSONSTruct(trackerInput NormalisedTrackerCSV) (NormalisedTrackerJSON, error) {
	var output NormalisedTrackerJSON

	output.Name = trackerInput.Title
	output.Story_type = trackerInput.Type
	output.Description = trackerInput.Description
	output.Labels = trackerInput.Labels
	output.Tasks = trackerInput.Tasks

	return output, nil
}
