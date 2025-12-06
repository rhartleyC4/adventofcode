package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhartleyC4/adventofcode/2025/day6"
)

var _ = Describe("Problem", func() {
	Describe("Calculate", func() {
		DescribeTable("test cases",
			func(operands []int64, operator main.Operator, expectedValue int64, shouldError bool) {
				subject := main.Problem{}
				subject.AddOperand(operands...)
				_ = subject.AddOperator(operator)

				result, err := subject.Calculate()

				if shouldError {
					Expect(err).To(HaveOccurred())
				} else {
					Expect(err).NotTo(HaveOccurred())
				}
				Ω(result).Should(Equal(expectedValue))
			},
			Entry("add two 1s", []int64{1, 1}, main.Add, int64(2), false),
			Entry("adds multiple numbers", []int64{8, 2, 1, 4, 100}, main.Add, int64(115), false),
			Entry("multiplies two 1s", []int64{1, 1}, main.Multiply, int64(1), false),
			Entry("multiplies multiple numbers", []int64{5, 5, 10}, main.Multiply, int64(250), false),
			Entry("errors when there are no operands", []int64{}, main.Add, int64(0), true),
			Entry("errors with unsupported operator", []int64{1, 1}, main.Operator(5), int64(0), true),
		)

		Describe("AddOperator", func() {
			DescribeTable("adding operator", func(operator main.Operator, shouldError bool) {
				subject := main.Problem{}
				if shouldError {
					Ω(subject.AddOperator(operator)).Should(HaveOccurred())
				} else {
					Ω(subject.AddOperator(operator)).Should(Succeed())
				}
			},
				Entry("supports Add", main.Add, false),
				Entry("supports Multiply", main.Multiply, false),
				Entry("does not support A", main.Operator('A'), true))
		})
	})
})

var _ = Describe("Homework", func() {
	Describe("AddLine", func() {
		DescribeTable("called with valid input",
			func(line string, operands []int64, operators []main.Operator, expectError bool) {
				subject := main.Homework{}

				result, err := subject.AddLine(line)

				if expectError {
					Ω(err).Should(HaveOccurred())
					Ω(result).Should(BeNil())
				} else {
					Ω(err).ShouldNot(HaveOccurred())
					Ω(result.Operators).Should(Equal(operators))
					Ω(result.Operands).Should(Equal(operands))
				}
			},
			Entry(`"123 328  51 64 " succeeds"`, "123 328  51 64 ", []int64{123, 328, 51, 64}, nil, false),
			Entry(`" 45 64  387 23 " succeeds`, " 45 64  387 23 ", []int64{45, 64, 387, 23}, nil, false),
			Entry(`"  6 98  215 314" succeeds`, "  6 98  215 314", []int64{6, 98, 215, 314}, nil, false),
			Entry(`"*   +   *   +  " succeeds`, "*   +   *   +  ", nil, []main.Operator{'*', '+', '*', '+'}, false),
			Entry(`"*   15  *   +  " fails`, "*   15  *   +  ", nil, nil, true),
			Entry(`"  2 BC    1   D" fails`, "A   BC    1   D", nil, nil, true),
			Entry(`"Hello`, "Hello", nil, nil, true),
			Entry(`"H e l l o`, "H e l l o", nil, nil, true),
			Entry(`"" fails`, "", nil, nil, true),
		)

		It("fails if you had another line that does not match the number of the first line", func() {
			subject := main.Homework{}
			Ω(subject.AddLine("1 2 3")).Error().ShouldNot(HaveOccurred())

			Ω(subject.AddLine("1 2")).Error().Should(HaveOccurred())
		})
	})

	Describe("Solve", func() {
		It("sums all problem calculations", func() {
			subject := main.Homework{}
			Ω(subject.AddLine("123 328  51 64 ")).Error().Should(Succeed())
			Ω(subject.AddLine(" 45 64  387 23 ")).Error().Should(Succeed())
			Ω(subject.AddLine("  6 98  215 314")).Error().Should(Succeed())
			Ω(subject.AddLine("*   +   *   +  ")).Error().Should(Succeed())

			result, err := subject.Solve()

			Ω(err).ShouldNot(HaveOccurred())
			Ω(result).Should(Equal(int64(4277556)))
		})

		It("fails when there are no operands", func() {
			subject := main.Homework{}

			Ω(subject.Solve()).Error().Should(HaveOccurred())
		})
	})
})
