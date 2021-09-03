package trackercsv_test

import (
	"bytes"

	"github.com/EngineerBetter/new-starters/tracker-csv-to-json/trackercsv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trackercsv", func() {
	PIt("normalises the csv", func() {
		csv := bytes.NewBufferString(`Title,Type,Description,Labels\nAAA,BBB,CCC,"DDD,EEE"`)
		normal := trackercsv.New(csv)

		Expect(normal).To(Equal(trackercsv.NormalisedTrackerCSV{
			Title:       "AAA",
			Type:        "BBB",
			Description: "CCC",
			Labels:      `"DDD,EEE"`,
		}))
	})
})
