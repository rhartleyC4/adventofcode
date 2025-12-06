package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Operator rune

const (
	Add      Operator = '+'
	Multiply Operator = '*'
)

type Problem struct {
	operands []int64
	operator Operator
}

func (p *Problem) AddOperand(operand ...int64) {
	p.operands = append(p.operands, operand...)
}

func (p *Problem) AddOperator(operator Operator) error {
	if !p.operatorSupported(operator) {
		return fmt.Errorf("operator %v is not supported", operator)
	}
	p.operator = operator
	return nil
}

func (p *Problem) operatorSupported(operator Operator) bool {
	switch operator {
	case Add, Multiply:
		return true
	default:
		return false
	}
}

func (p *Problem) Calculate() (int64, error) {
	if len(p.operands) == 0 {
		return 0, errors.New("no operands found")
	}

	switch p.operator {
	case Add:
		var sum int64
		for _, operand := range p.operands {
			sum += operand
		}
		return sum, nil
	case Multiply:
		product := p.operands[0]
		for _, operand := range p.operands[1:] {

			product *= operand
		}
		return product, nil
	default:
		return 0, errors.New("unknown operator")
	}
}

type LineInfo struct {
	Operands  []int64
	Operators []Operator
}

type Homework struct {
	problems []Problem
}

func (h *Homework) AddLine(line string) (*LineInfo, error) {
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return nil, fmt.Errorf("invalid line: %q", line)
	}

	result := &LineInfo{}
	firstOperand, err := strconv.ParseInt(fields[0], 10, 64)
	if err == nil {
		result.Operands = []int64{firstOperand}
		for i, operand := range fields[1:] {
			value, err := strconv.ParseInt(operand, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid operand in col %d: %q", i+2, operand)
			}
			result.Operands = append(result.Operands, value)
		}
	} else {
		for i, operator := range fields {
			if len(operator) != 1 {
				return nil, fmt.Errorf("invalid operand in col %d: %q", i+1, operator)
			}
			result.Operators = append(result.Operators, Operator(operator[0]))
		}
	}

	if h.problems == nil {
		h.problems = make([]Problem, len(fields))
	}

	if len(h.problems) != len(fields) {
		return nil, fmt.Errorf("invalid number of operands/operators: %d != %d", len(h.problems), len(fields))
	}

	if result.Operands != nil {
		for i, operand := range result.Operands {
			h.problems[i].AddOperand(operand)
		}
	}

	if result.Operators != nil {
		for i, operand := range result.Operators {
			err := h.problems[i].AddOperator(operand)
			if err != nil {
				return nil, fmt.Errorf("invalid operand in col %d: %q", i+1, operand)
			}
		}
	}

	return result, nil
}

func (h *Homework) Solve() (int64, error) {
	if len(h.problems) == 0 {
		return 0, errors.New("no problems found")
	}
	var total int64
	for i, problem := range h.problems {
		value, err := problem.Calculate()
		if err != nil {
			return 0, fmt.Errorf("failed to calculate col %d: %v", i+1, err)
		}
		total += value
	}
	return total, nil
}
