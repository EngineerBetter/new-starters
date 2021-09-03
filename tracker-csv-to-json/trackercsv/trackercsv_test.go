package trackercsv_test

import (
	"bytes"

	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackercsv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trackercsv", func() {
	It("normalises the csv", func() {
		csv := bytes.NewBufferString("Title,Type,Description,Labels\nAAA,BBB,CCC,\"DDD,EEE\"")
		normal, err := trackercsv.New(csv)

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal([]trackercsv.NormalisedTrackerCSV{{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      "DDD,EEE",
			Tasks:       "",
		}}))
	})

	It("normalises two lines of the csv", func() {
		csv := bytes.NewBufferString("Title,Type,Description,Labels\nAAA,BBB,CCC,\"DDD,EEE\"\nFFF,GGG,HHH,\"III,JJJ\"")
		normal, err := trackercsv.New(csv)

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal([]trackercsv.NormalisedTrackerCSV{{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      "DDD,EEE",
			Tasks:       "",
		}, {
			Title:       "FFF",
			Type:        "GGG",
			Description: "HHH",
			Labels:      "III,JJJ",
			Tasks:       "",
		}}))
	})

	It("normalises one task", func() {
		csv := bytes.NewBufferString("Title,Type,Description,Labels, Tasks\nAAA,BBB,CCC,\"DDD,EEE\",Completed: Methods")

		normal, err := trackercsv.New(csv)

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal([]trackercsv.NormalisedTrackerCSV{{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      "DDD,EEE",
			Tasks:       "Completed: Methods",
		}}))
	})

	// One task

	// Many tasks

})
