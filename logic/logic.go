package logic

// Operand is the template for operands.
type Operand interface {
	Evaluate(ctx interface{}) (bool, error)
}

// Operator is the template for operators.
type Operator struct {
	operands []Operand
}

// AndOperator is the implementation of the boolean AND operators.
type AndOperator struct {
	Operator
}

// Evaluate performs the evaluation of the boolean AND operator.
func (o AndOperator) Evaluate(ctx interface{}) (bool, error) {
	// fmt.Println("And.Evaluate()")
	for _, operand := range o.operands {
		var t bool
		var err error
		if t, err = operand.Evaluate(ctx); err != nil {
			return false, err
		}
		if !t {
			return false, nil
		}
	}
	return true, nil
}

// And generates a new AndOperator with the given set of operands.
func And(operands ...Operand) AndOperator {
	return AndOperator{
		Operator{
			operands,
		},
	}
}

// All generates a new AndOperator with the given set of operands.
func All(operands ...Operand) AndOperator {
	return And(operands...)
}

// OrOperator is the template for boolean OR operators.
type OrOperator struct {
	Operator
}

// Evaluate performs the evaluation of the boolean OR operator.
func (o OrOperator) Evaluate(ctx interface{}) (bool, error) {
	// fmt.Println("Or.Evaluate()")
	for _, operand := range o.operands {
		var t bool
		var err error
		if t, err = operand.Evaluate(ctx); err != nil {
			return false, err
		}
		if t {
			return true, nil
		}
	}
	return false, nil
}

// Or generates a new OrOperator with the given set of operands.
func Or(operands ...Operand) OrOperator {
	return OrOperator{
		Operator{
			operands,
		},
	}
}

// Any generates a new OrOperator with the given set of operands.
func Any(operands ...Operand) OrOperator {
	return Or(operands...)
}

// NotOperator is the implementation of the boolean NOT operators.
type NotOperator struct {
	operand Operand
}

// Evaluate performs the evaluation of the boolean NOT operator.
func (o NotOperator) Evaluate(ctx interface{}) (bool, error) {
	// fmt.Println("Not.Evaluate()")
	result, err := o.operand.Evaluate(ctx)
	return !result, err
}

// Not generates a new NotOperator with the given operand.
func Not(operand Operand) NotOperator {
	return NotOperator{
		operand,
	}
}

// XorOperator is the implementation of the boolean XOR operators.
type XorOperator struct {
	operand1 Operand
	operand2 Operand
}

// Xor generates a new XorOperator with the given operands.
func Xor(operand1, operand2 Operand) XorOperator {
	return XorOperator{
		operand1,
		operand2,
	}
}

// Evaluate performs the evaluation of the boolean XOR operator.
func (o XorOperator) Evaluate(ctx interface{}) (bool, error) {
	// fmt.Println("Xor.Evaluate()")
	result1, err := o.operand1.Evaluate(ctx)
	if err != nil {
		return result1, err
	}
	result2, err := o.operand2.Evaluate(ctx)
	return result1 != result2, err
}

// BoolOperand is the Operand version of the boolean native values.
type BoolOperand bool

const (
	// True represents the boolean true value.
	True BoolOperand = true
	// False represents the boolean false value.
	False BoolOperand = false
)

// Evaluate performs the evaluation of the typedef'd native boolean type.
func (o BoolOperand) Evaluate(ctx interface{}) (bool, error) {
	return bool(o), nil
}

