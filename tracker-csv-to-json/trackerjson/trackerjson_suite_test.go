package trackerjson_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTrackerjson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trackerjson Suite")
}
