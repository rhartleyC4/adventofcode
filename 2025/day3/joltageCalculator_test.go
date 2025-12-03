package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhartleyC4/adventofcode/2025/day3"
)

var _ = Describe("JoltageCalculator", func() {
	Describe("AddBank", func() {
		DescribeTable("successful additions",
			func(input string, expected uint64) {
				subject := main.NewJoltageCalculator(12)

				// act
				result, err := subject.AddBank(input)

				Ω(err).ShouldNot(HaveOccurred())
				Ω(result).Should(Equal(expected))
			},
			Entry("calculates input 1", "987654321111111", uint64(987654321111)),
			Entry("calculates input 2", "811111111111119", uint64(811111111119)),
			Entry("calculates input 3", "234234234234278", uint64(434234234278)),
			Entry("calculates input 4", "818181911112111", uint64(888911112111)),
			Entry("calculates input 5", "7657222591427217122445272425253565561122226426262235211484443362522725526264152632322121122211226432", uint64(987666336432)),
		)
	})

	Describe("TotalJolts", func() {
		It("sums all bank jolts", func() {
			subject := main.NewJoltageCalculator(12)

			Ω(subject.AddBank("987654321111111")).Error().ShouldNot(HaveOccurred())
			Ω(subject.AddBank("811111111111119")).Error().ShouldNot(HaveOccurred())
			Ω(subject.AddBank("234234234234278")).Error().ShouldNot(HaveOccurred())
			Ω(subject.AddBank("818181911112111")).Error().ShouldNot(HaveOccurred())

			Ω(subject.TotalJolts()).Should(Equal(uint64(3121910778619)))
		})
	})

	It("fails if a bank does not have numeric values", func() {
		subject := main.NewJoltageCalculator(1)
		_, err := subject.AddBank("1A34")

		Ω(err).Should(HaveOccurred())
	})
})
