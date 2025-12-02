package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDay2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day2 Suite")
}
