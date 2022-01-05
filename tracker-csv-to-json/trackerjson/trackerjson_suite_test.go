package trackerjson_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTrackerjson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trackerjson Suite")
}
