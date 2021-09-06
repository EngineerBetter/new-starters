package trackerjson_test

import (
	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackerjson"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trackerjson", func() {
	It("Returns one jsonl object", func() {
		input := []trackerjson.NormalisedTrackerCSV{{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      "DDD,EEE",
			Tasks:       "",
		}}

		normal, err := trackerjson.Convert(input)
		res := []byte(`{"name":"AAA","story_type":"BBB","description":"CCC","labels":"DDD,EEE","tasks":""}`)

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal(res))
	})

	It("Returns two jsonl objects", func() {
		input := []trackerjson.NormalisedTrackerCSV{{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      "DDD,EEE",
			Tasks:       "",
		},
			{
				Title:       "TEST2",
				Type:        "TEST3",
				Description: "TEST4",
				Labels:      "TEST5,TEST6",
				Tasks:       "",
			}}

		normal, err := trackerjson.Convert(input)
		res := []byte(`{"name":"AAA","story_type":"BBB","description":"CCC","labels":"DDD,EEE","tasks":""}{"name":"TEST2","story_type":"TEST3","description":"TEST4","labels":"TEST5,TEST6","tasks":""}`)

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal(res))
	})
})
