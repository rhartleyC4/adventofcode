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
		Entry("finds 2-digit and 3-digit in range", "95-115", []uint64{99, 111}),
		Entry("finds 3-digit and 4-digit in range", "998-1012", []uint64{999, 1010}),
		Entry("finds large value", "1188511880-1188511890", []uint64{1188511885}),
		Entry("finds one value with many repeating characters", "222220-222224", []uint64{222222}),
		Entry("does not find in range", "1698522-1698528", []uint64{}),
		Entry("finds one in split repeating numbers", "446443-446449", []uint64{446446}),
		Entry("finds 4-digit one in range", "38593856-38593862", []uint64{38593859}),
		Entry("finds one 2-digit three-peat", "565653-565659", []uint64{565656}),
		Entry("find one 3-digit three-peat", "824824821-824824827", []uint64{824824824}),
		Entry("finds one 2-digit five-peat", "2121212118-2121212124", []uint64{2121212121}),
	)

	//It("get's expected values", func() {
	//	i := "2121212121"
	//
	//	Ω(main.GCD(len(i), 5)).Should(Equal(5))
	//	Ω(main.GCD(len(i), 4)).Should(Equal(2))
	//	Ω(main.GCD(len(i), 3)).Should(Equal(1))
	//	Ω(main.GCD(len(i), 2)).Should(Equal(2))
	//
	//	Ω(main.GCD(3, 2)).Should(Equal(1))
	//	Ω(main.GCD(3, 1)).Should(Equal(1))
	//})
})
