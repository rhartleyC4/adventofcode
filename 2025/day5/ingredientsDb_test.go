package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhartleyC4/adventofcode/2025/day5"
)

var _ = Describe("IngredientsDb", func() {
	var subject main.IngredientsDb

	BeforeEach(func() {
		subject = main.IngredientsDb{}
	})

	Describe("AddFresh", func() {
		It("successfully adds a valid range string", func() {
			expectedMin := uint64(474206951121632)
			expectedMax := uint64(478696506672479)
			input := fmt.Sprintf("%d-%d", expectedMin, expectedMax)

			// act
			result, err := subject.AddFresh(input)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(result.Min).Should(Equal(expectedMin))
			Ω(result.Max).Should(Equal(expectedMax))
		})
	})

	Context("with sample fresh", func() {
		BeforeEach(func() {
			Ω(subject.AddFresh("3-5")).Error().ShouldNot(HaveOccurred())
			Ω(subject.AddFresh("10-14")).Error().ShouldNot(HaveOccurred())
			Ω(subject.AddFresh("16-20")).Error().ShouldNot(HaveOccurred())
			Ω(subject.AddFresh("12-18")).Error().ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			subject.Reset()
		})

		Describe("CheckIngredient", func() {
			DescribeTable("with ingredient",
				func(input string, expected bool) {
					Ω(subject.CheckIngredient(input)).Should(Equal(expected))
				},
				Entry("returns false for spoiled ingredient 1", "1", false),
				Entry("returns true for fresh ingredient 5", "5", true),
				Entry("returns false for spoiled ingredient 8", "8", false),
				Entry("returns true for fresh ingredient 11", "11", true),
				Entry("returns true for fresh ingredient 17", "17", true),
				Entry("returns false for spoiled ingredient 32", "32", false),
			)
		})

		Describe("TotalFresh", func() {
			BeforeEach(func() {
				Ω(subject.CheckIngredient("1")).Error().ShouldNot(HaveOccurred())
				Ω(subject.CheckIngredient("5")).Error().ShouldNot(HaveOccurred())
				Ω(subject.CheckIngredient("8")).Error().ShouldNot(HaveOccurred())
				Ω(subject.CheckIngredient("11")).Error().ShouldNot(HaveOccurred())
				Ω(subject.CheckIngredient("17")).Error().ShouldNot(HaveOccurred())
				Ω(subject.CheckIngredient("32")).Error().ShouldNot(HaveOccurred())
			})

			It("should give the correct number of fresh ingredients", func() {
				Ω(subject.TotalFresh()).Should(Equal(3))
			})
		})

		Describe("TotalConsideredFresh", func() {
			It("should give the correct number of possible fresh ingredients", func() {
				Ω(subject.TotalConsideredFresh()).Should(Equal(uint64(14)))
			})
		})
	})
})
