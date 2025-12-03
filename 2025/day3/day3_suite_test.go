package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDay3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day3 Suite")
}
