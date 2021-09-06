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

	// PIt("Returns two jsonl objects", func() {
	// 	input := trackerjson.NormalisedTrackerCSV{{
	// 		Title:       "AAA",
	// 		Type:        "BBB",
	// 		Description: "CCC",
	// 		Labels:      []string{"DDD", "EEE"},
	// 		Tasks:       []string{},
	// 	},
	// 		{
	// 			Title:       "TEST2",
	// 			Type:        "TEST3",
	// 			Description: "TEST4",
	// 			Labels:      []string{"TEST5", "TEST6"},
	// 			Tasks:       []string{},
	// 		}}

	// 	normal, err := trackerjson.ConvertCSVStructToJSONSTruct(input)
	// 	res := []byte(`{"name":"AAA","story_type":"BBB","description":"CCC","labels":"DDD,EEE","tasks":""}{"name":"TEST2","story_type":"TEST3","description":"TEST4","labels":"TEST5, TEST6","tasks":""}`)

	// 	Expect(err).NotTo(HaveOccurred())
	// 	Expect(normal).To(Equal(res))
	// })
})
