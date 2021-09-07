package acceptance_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os/exec"

	. "github.com/onsi/ginkgo"
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
			scanner := bufio.NewScanner(&output)
			for scanner.Scan() {
				l := line{}
				Expect(json.Unmarshal(scanner.Bytes(), &l)).To(Succeed())
			}
		})
	})
})
