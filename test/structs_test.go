package structs_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestStream(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Structs Tests")
}
