package sims_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSims(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sims Suite")
}
