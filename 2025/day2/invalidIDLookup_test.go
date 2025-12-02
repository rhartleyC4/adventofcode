package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhartleyC4/adventofcode/2025/day2"
)

var _ = Describe("InvalidIDLookup", func() {
	DescribeTable("given ranges", func(r string, expected []uint64) {
		Ω(main.InvalidIDLookup(r)).Should(Equal(expected))
	},
		Entry("finds two values at end of range", "11-22", []uint64{11, 22}),
		Entry("finds 2-digit one in range", "95-115", []uint64{99}),
		Entry("finds 4-digit value towards end of range", "998-1012", []uint64{1010}),
		Entry("finds large value", "1188511880-1188511890", []uint64{1188511885}),
		Entry("finds one value with many repeating characters", "222220-222224", []uint64{222222}),
		Entry("does not find in range", "1698522-1698528", []uint64{}),
		Entry("finds one in split repeating numbers", "446443-446449", []uint64{446446}),
		Entry("finds 4-digit one in range", "38593856-38593862", []uint64{38593859}),
		Entry("does not find in range 1", "565653-565659", []uint64{}),
		Entry("does not find in range 2", "824824821-824824827", []uint64{}),
		Entry("does not find in range 3", "2121212118-2121212124", []uint64{}),
	)
})
