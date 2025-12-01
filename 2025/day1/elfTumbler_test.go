package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rhartleyC4/adventofcode/2025/day1"
)

var _ = Describe("ElfTumbler", func() {
	Describe("NewElfTumbler", func() {
		It("returns a tumbler given valid min and max", func() {
			subject := main.NewElfTumbler(0, 20)

			Ω(subject).ShouldNot(BeNil())
		})

		It("panics when min is greater than max", func() {
			Ω(func() { main.NewElfTumbler(20, 0) }).Should(Panic())
		})
	})

	Context("methods", func() {
		var (
			minValue    int
			maxValue    int
			midPoint    int
			adjustments = []string{
				"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82", "R501", "R17", "R1000",
			}
			expected = []int{
				82, 52, 0, 95, 55, 0, 99, 0, 14, 32, 33, 50, 50,
			}
			subject *main.ElfTumbler
		)

		BeforeEach(func() {
			minValue = 0
			maxValue = 99
			midPoint = 50
			subject = main.NewElfTumbler(minValue, maxValue)
		})

		Describe("GetValue", func() {
			It("defaults to mid point of min and max", func() {
				Ω(subject.GetValue()).To(Equal(midPoint))
			})
		})

		Describe("AdvanceDial", func() {
			It("makes correct adjustments", func() {
				for index, adjustment := range adjustments {
					// act
					Ω(subject.Adjust(adjustment)).Should(Succeed())

					Ω(subject.GetValue()).Should(Equal(expected[index]))
				}
			})
		})

		Describe("GetPassword", func() {
			It("returns the number of zeros hit in the sequence of adjustments", func() {
				for _, adjustment := range adjustments {
					Ω(subject.Adjust(adjustment)).Should(Succeed())
				}

				Ω(subject.GetPassword()).To(Equal(21))
			})
		})
	})
})
