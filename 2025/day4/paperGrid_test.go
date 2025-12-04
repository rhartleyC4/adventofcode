package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhartleyC4/adventofcode/2025/day4"
)

var _ = Describe("PaperGrid", func() {
	Describe("NewStorage", func() {
		It("successfully parses stock with non-default tokens", func() {
			empty := '*'
			occupied := '#'
			separator := "|"
			stock := fmt.Sprintf("%[1]c%[1]c%[1]c%[3]s%[1]c%[2]c%[1]c%[3]s%[2]c%[2]c%[2]c",
				empty, occupied, separator)

			result, err := main.NewStorage(stock,
				main.WithEmptyRune(empty),
				main.WithOccupiedRune(occupied),
				main.WithRowSeparator(separator))

			Ω(err).ShouldNot(HaveOccurred())
			Ω(result).ShouldNot(BeNil())
		})

		It("successfully parses stock with default tokens", func() {
			stock := ".@.\n..@\n@@@\n..."

			result, err := main.NewStorage(stock)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(result).ShouldNot(BeNil())
		})

		It("errors with stock that does not match the separator", func() {
			stock := "The quick brown fox jumped the fence"

			result, err := main.NewStorage(stock, main.WithRowSeparator("\n"))

			Ω(err).Should(HaveOccurred())
			Ω(result).Should(BeNil())
		})

		It("errors if the stock has tokens not configured", func() {
			stock := "foo\n@@@@|.@a"

			result, err := main.NewStorage(stock)

			Ω(err).Should(HaveOccurred())
			Ω(result).Should(BeNil())
		})

		It("errors if the stock is a jagged array", func() {
			stock := "...\n@@@@\n.@."

			result, err := main.NewStorage(stock)

			Ω(err).Should(HaveOccurred())
			Ω(result).Should(BeNil())
		})
	})

	Describe("ForkLiftAccessibleItems", func() {
		It("finds correct number of items", func() {
			stock := "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@."
			subject, err := main.NewStorage(stock)
			Ω(err).ShouldNot(HaveOccurred())

			// act
			result := subject.ForkLiftAccessibleItems(false)

			Ω(result).Should(Equal(13))
		})
	})
})
