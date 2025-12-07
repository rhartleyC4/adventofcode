package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rhartleyC4/adventofcode/2025/day6"
)

var _ = Describe("Problem", func() {
	Describe("Calculate", func() {
		DescribeTable("test cases",
			func(operands [][]main.OperandPart, operator main.Operator, expectedValue int64, shouldError bool) {
				subject := main.NewProblem(operator)
				subject.AddOperand(operands...)

				result, err := subject.Calculate()

				if shouldError {
					Expect(err).To(HaveOccurred())
				} else {
					Expect(err).NotTo(HaveOccurred())
				}
				Ω(result).Should(Equal(expectedValue))
			},
			Entry("add two 1s",
				[][]main.OperandPart{
					{{0, 1}}, // 1
					{{0, 1}}, // 1
				}, main.Add, int64(2), false),
			Entry("adds multiple numbers",
				[][]main.OperandPart{
					{{0, 0}, {1, 0}, {2, 8}}, // 8
					{{0, 0}, {1, 0}, {2, 2}}, // 2
					{{0, 0}, {1, 1}, {2, 0}}, // 10
					{{0, 0}, {1, 0}, {2, 4}}, // 4
					{{0, 1}, {1, 0}, {2, 0}}, // 100
				}, main.Add, int64(124), false),
			Entry("multiplies two 1s",
				[][]main.OperandPart{
					{{0, 1}}, // 1
					{{0, 1}}, // 1
				}, main.Multiply, int64(1), false),
			Entry("multiplies multiple numbers",
				[][]main.OperandPart{
					{{0, 0}, {1, 5}}, // 5
					{{0, 0}, {1, 5}}, // 5
					{{0, 1}, {1, 0}}, // 10
				}, main.Multiply, int64(250), false),
			Entry("errors when there are no operands",
				[][]main.OperandPart{}, main.Add, int64(0), true),
			Entry("errors with unsupported operator",
				[][]main.OperandPart{
					{{0, 1}}, // 1
					{{0, 1}}, // 1
				}, main.Operator(5), int64(0), true),
		)
	})
})

var _ = Describe("Homework", func() {
	Describe("AddLine", func() {
		DescribeTable("called with valid input",
			func(line string, expectError bool) {
				subject := main.Homework{}

				err := subject.AddLine(line)

				if expectError {
					Ω(err).Should(HaveOccurred())
				} else {
					Ω(err).ShouldNot(HaveOccurred())
				}
			},
			Entry(`"123 328  51 64 " succeeds`, "123 328  51 64 ", false),
			Entry(`" 45 64  387 23 " succeeds`, " 45 64  387 23 ", false),
			Entry(`"  6 98  215 314" succeeds`, "  6 98  215 314", false),
			Entry(`"*   +   *   +  " succeeds`, "*   +   *   +  ", false),
			Entry(`"  2 BC    1   D" fails`, "A   BC    1   D", true),
			Entry(`"Hello`, "Hello", true),
			Entry(`"H e l l o`, "H e l l o", true),
			Entry(`"" fails`, "", true),
		)
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
			Ω(result).Should(Equal(int64(3263827)))
		})

		It("fails when there are no operands", func() {
			subject := main.Homework{}

			Ω(subject.Solve()).Error().Should(HaveOccurred())
		})
	})
})
