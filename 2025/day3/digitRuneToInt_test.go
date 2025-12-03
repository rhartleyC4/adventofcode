package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rhartleyC4/adventofcode/2025/day3"
)

var _ = Describe("DigitRuneToInt", func() {
	DescribeTable("conversion",
		func(input rune, expectedInt int, expectedBool bool) {
			resultInt, resultBool := main.DigitRuneToInt(input)

			Ω(resultInt).To(Equal(expectedInt))
			Ω(resultBool).To(Equal(expectedBool))
		},
		Entry("A", 'A', 0, false),
		Entry("0", '0', 0, true),
		Entry("1", '1', 1, true),
		Entry("2", '2', 2, true),
		Entry("3", '3', 3, true),
		Entry("4", '4', 4, true),
		Entry("5", '5', 5, true),
		Entry("6", '6', 6, true),
		Entry("7", '7', 7, true),
		Entry("8", '8', 8, true),
		Entry("9", '9', 9, true),
	)
})
