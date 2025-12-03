package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhartleyC4/adventofcode/2025/day3"
)

var _ = Describe("JoltageCalculator", func() {
	Describe("AddBank", func() {
		DescribeTable("successful additions",
			func(input string, expected int) {
				subject := main.NewJoltageCalculator()

				// act
				result, err := subject.AddBank(input)

				Ω(err).ShouldNot(HaveOccurred())
				Ω(result).Should(Equal(expected))
			},
			Entry("calculates input 1", "987654321111111", 98),
			Entry("calculates input 2", "811111111111119", 89),
			Entry("calculates input 3", "234234234234278", 78),
			Entry("calculates input 4", "818181911112111", 92),
		)
	})

	Describe("TotalJolts", func() {
		It("sums all bank jolts", func() {
			subject := main.NewJoltageCalculator()

			Ω(subject.AddBank("987654321111111")).Error().ShouldNot(HaveOccurred())
			Ω(subject.AddBank("811111111111119")).Error().ShouldNot(HaveOccurred())
			Ω(subject.AddBank("234234234234278")).Error().ShouldNot(HaveOccurred())
			Ω(subject.AddBank("818181911112111")).Error().ShouldNot(HaveOccurred())

			Ω(subject.TotalJolts()).Should(Equal(357))
		})
	})

	It("fails if a bank does not have numeric values", func() {
		subject := main.NewJoltageCalculator()
		result, err := subject.AddBank("1A34")

		Ω(err).Should(HaveOccurred())
		Ω(result).Should(Equal(0))
	})
})
