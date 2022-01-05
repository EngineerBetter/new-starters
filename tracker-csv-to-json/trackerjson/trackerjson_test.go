package trackerjson_test

import (
	"encoding/json"

	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackercsv"
	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackerjson"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trackerjson", func() {
	It("Returns one NormalisedTrackerJSON", func() {
		input := trackercsv.NormalisedTrackerCSV{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []string{},
		}

		normal, err := trackerjson.ConvertCSVStructToJSONSTruct(input)
		res := trackerjson.NormalisedTrackerJSON{
			Name:        "AAA",
			Story_type:  "BBB",
			Description: "CCC",
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []trackerjson.TrackerTask{},
		}

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal(res))
	})

	It("Returns one NormalisedTrackerJSON with blank labels", func() {
		input := trackercsv.NormalisedTrackerCSV{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      []string{},
			Tasks:       []string{},
		}

		normal, err := trackerjson.ConvertCSVStructToJSONSTruct(input)
		res := trackerjson.NormalisedTrackerJSON{
			Name:        "AAA",
			Story_type:  "BBB",
			Description: "CCC",
			Labels:      []string{},
			Tasks:       []trackerjson.TrackerTask{},
		}

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal(res))
	})

	It("Returns two jsonl objects", func() {
		input := []trackercsv.NormalisedTrackerCSV{{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []string{},
		},
			{
				Title:       "TEST2",
				Type:        "TEST3",
				Description: "TEST4",
				Labels:      []string{"TEST5", "TEST6"},
				Tasks:       []string{},
			}}

		normal, err := trackerjson.BulkConvert(input)

		res := []trackerjson.NormalisedTrackerJSON{{
			Name:        "AAA",
			Story_type:  "BBB",
			Description: "CCC",
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []trackerjson.TrackerTask{},
		},
			{
				Name:        "TEST2",
				Story_type:  "TEST3",
				Description: "TEST4",
				Labels:      []string{"TEST5", "TEST6"},
				Tasks:       []trackerjson.TrackerTask{},
			}}

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal(res))
	})

	It("Converts one NormalisedTrackerJSON to json object", func() {
		trackerStruct := []trackerjson.NormalisedTrackerJSON{{
			Name:        "AAA",
			Story_type:  "BBB",
			Description: "CCC",
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []trackerjson.TrackerTask{},
		}}
		jsonStruct, _ := json.Marshal(trackerStruct)

		Expect(jsonStruct).Should(MatchJSON(`[{"name":"AAA","story_type":"BBB","description":"CCC","labels":["DDD","EEE"],"tasks":[]}]`))

	})
})
