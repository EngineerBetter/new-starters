package trackerjson_test

import (
	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackerjson"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trackerjson", func() {
	It("Returns one NormalisedTrackerJSON", func() {
		input := trackerjson.NormalisedTrackerCSV{
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
			Tasks:       []string{},
		}

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal(res))
	})

	It("Returns two jsonl objects", func() {
		input := []trackerjson.NormalisedTrackerCSV{{
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

		normal, err := trackerjson.Convert(input)

		res := []trackerjson.NormalisedTrackerJSON{{
			Name:        "AAA",
			Story_type:  "BBB",
			Description: "CCC",
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []string{},
		},
			{
				Name:        "TEST2",
				Story_type:  "TEST3",
				Description: "TEST4",
				Labels:      []string{"TEST5", "TEST6"},
				Tasks:       []string{},
			}}

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal(res))
	})
})
