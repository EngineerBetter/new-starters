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
})
