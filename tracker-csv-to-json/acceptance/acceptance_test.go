package acceptance_test

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type line struct {
	Name        string   `json:"name"`
	StoryType   string   `json:"story_type"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
	Tasks       []string `json:"tasks"`
}

var _ = Describe("Acceptance", func() {
	It("converts a CSV file to JSONL", func() {
		cmd := exec.Command("go", "run", "github.com/EngineerBetter/new-starters/tracker-csv-to-json", "./fixtures/golang.csv")
		output := bytes.Buffer{}
		cmd.Stdout = &output
		Expect(cmd.Run()).To(Succeed())

		By("not failing", func() {
			Expect(cmd.ProcessState.ExitCode()).To(BeZero())
		})

		By("outputting JSONL", func() {
			lines := strings.Split(strings.TrimSpace(output.String()), "\n")
			content, err := ioutil.ReadFile("fixtures/golang.jsonl")
			Expect(err).NotTo(HaveOccurred())
			fixtureLines := strings.Split(strings.TrimSpace(string(content)), "\n")

			Expect(lines).To(HaveLen(len(fixtureLines)))
			for i := range lines {
				Expect(lines[i]).To(MatchJSON(fixtureLines[i]))
			}
		})
	})
})
