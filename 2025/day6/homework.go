package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type Operator rune

const (
	Add      Operator = '+'
	Multiply Operator = '*'
)

type Problem struct {
	operands [][]OperandPart
	operator Operator
}

func (p *Problem) AddOperand(operand ...[]OperandPart) {
	p.operands = append(p.operands, operand...)
}

func (p *Problem) DigitSize() int {
	if len(p.operands) == 0 {
		panic(errors.New("no operands"))
	}
	return len(p.operands[0])
}

func (p *Problem) Calculate() (int64, error) {
	if len(p.operands) == 0 {
		return 0, errors.New("no operands found")
	}

	switch p.operator {
	case Add:
		var sum int64
		for _, operand := range p.operands {
			sum += p.getNumber(operand)
		}
		return sum, nil
	case Multiply:
		product := p.getNumber(p.operands[0])
		for _, operand := range p.operands[1:] {
			product *= p.getNumber(operand)
		}
		return product, nil
	default:
		return 0, errors.New("unknown operator")
	}
}

func (p *Problem) getNumber(input []OperandPart) int64 {
	digitCount := len(input)
	number := int64(0)
	for _, part := range input {
		p := digitCount - part.Significance - 1
		number += int64(part.Digit) * int64(math.Pow10(p))
	}
	return number
}

func NewProblem(operator Operator) *Problem {
	return &Problem{
		operator: operator,
	}
}

func IsSupportedOperator(operator Operator) bool {
	switch operator {
	case Add, Multiply:
		return true
	default:
		return false
	}
}

type OperandPart struct {
	Significance int
	Digit        int
}

type Homework struct {
	problems [][]uint8
}

func (h *Homework) AddLine(line string) error {
	if len(line) == 0 {
		return fmt.Errorf("invalid line: %q", line)
	}

	if h.problems == nil {
		h.problems = make([][]uint8, 0)
	}

	characters := make([]uint8, len(line))
	for i := 0; i < len(line); i++ {
		c := line[i]
		if IsSupportedOperator(Operator(c)) || c == ' ' {
			characters[i] = c
			continue
		}
		if c >= '0' && c <= '9' {
			characters[i] = c - '0'
		} else {
			return fmt.Errorf("invalid entry at col: %d", i+1)
		}
	}

	h.problems = append(h.problems, characters)
	return nil
}

func (h *Homework) Solve() (int64, error) {
	if len(h.problems) == 0 {
		return 0, errors.New("no problems found")
	}
	var total int64
	operators, err := h.getOperators()
	if err != nil {
		return 0, err
	}
	allOperands := h.problems[:len(h.problems)-1]
	currentNumber := 0
	for _, operator := range operators {
		problem := NewProblem(operator)
		for {
			var operands []OperandPart
			for _, operand := range allOperands {
				if currentNumber == len(operand) {
					break
				}
				p := operand[currentNumber]
				if p == ' ' {
					continue
				}
				d := OperandPart{
					Significance: len(operands),
					Digit:        int(p),
				}
				operands = append(operands, d)
			}
			currentNumber++
			if len(operands) == 0 {
				// can calculate
				value, err := problem.Calculate()
				if err != nil {
					return 0, err
				}
				total += value
				break
			} else {
				problem.AddOperand(operands)
			}
		}
	}
	return total, nil
}

func (h *Homework) getOperators() ([]Operator, error) {
	operators := strings.Fields(string(h.problems[len(h.problems)-1]))
	result := make([]Operator, len(operators))
	for i, o := range operators {
		if len(o) != 1 {
			return nil, fmt.Errorf("invalid operator: %v", o)
		}
		result[i] = Operator(o[0])
	}
	return result, nil
}
