package trackercsv_test

import (
	"bytes"

	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackercsv"
	. "github.com/onsi/ginkgo/v2"
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
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []string{},
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
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []string{},
		}, {
			Title:       "FFF",
			Type:        "GGG",
			Description: "HHH",
			Labels:      []string{"III", "JJJ"},
			Tasks:       []string{},
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
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []string{"Completed: Methods"},
		}}))
	})

	It("normalises many tasks", func() {
		csv := bytes.NewBufferString("Title,Type,Description,Labels, Task, Task\nAAA,BBB,CCC,\"DDD,EEE\",Completed: AAA,Completed: BBB")

		normal, err := trackercsv.New(csv)

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal([]trackercsv.NormalisedTrackerCSV{{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []string{"Completed: AAA", "Completed: BBB"},
		}}))
	})

	It("normalises one task, when the file has many task columns", func() {
		csv := bytes.NewBufferString("Title,Type,Description,Labels, Task, Task, Task\nAAA,BBB,CCC,\"DDD,EEE\",Completed: AAA,Completed: BBB,\"\"")

		normal, err := trackercsv.New(csv)

		Expect(err).NotTo(HaveOccurred())
		Expect(normal).To(Equal([]trackercsv.NormalisedTrackerCSV{{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      []string{"DDD", "EEE"},
			Tasks:       []string{"Completed: AAA", "Completed: BBB"},
		}}))
	})

	It("Does not include a labels key/value pair if labels are blank", func() {
		csv := bytes.NewBufferString("Title,Type,Description,Labels, Task, Task, Task\nAAA,BBB,CCC,\"\",Completed: AAA,Completed: BBB,\"\"")
		normal, err := trackercsv.New(csv)

		Expect(err).NotTo(HaveOccurred())

		Expect(normal).To(Equal([]trackercsv.NormalisedTrackerCSV{{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      []string{},
			Tasks:       []string{"Completed: AAA", "Completed: BBB"},
		}}))
	})
})
