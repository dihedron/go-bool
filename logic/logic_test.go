package logic

import (
	"fmt"
	"testing"
)

type test struct {
	label    string
	operand  Operand
	expected bool
	err      error
}

func runTestImpl(t *testing.T, test test) {
	t.Logf("running test %q...\n", test.label)
	actual, err := test.operand.Evaluate(nil)
	if err != test.err {
		t.Fatalf("test %q: expected error %v, got error %v\n", test.label, test.err, err)
	}
	if err == nil && actual != test.expected {
		t.Fatalf("test %q: expected value %t, got value %t\n", test.label, test.expected, actual)
	}
}

func TestAnd(t *testing.T) {
	for _, test := range []test{
		test{
			"and (t) -> t",
			And(
				True,
			),
			true,
			nil,
		},
		test{
			"and (f) -> f",
			And(
				False,
			),
			false,
			nil,
		},
		test{
			"and (t,t) -> t",
			And(
				True,
				True,
			),
			true,
			nil,
		},
		test{
			"and (t,f) -> f",
			And(
				True,
				False,
			),
			false,
			nil,
		},
		test{
			"and (f,t) -> f",
			And(
				False,
				True,
			),
			false,
			nil,
		},
		test{
			"and (f,f) -> f",
			And(
				False,
				False,
			),
			false,
			nil,
		},
		test{
			"and (t,t,f) -> f",
			And(
				True,
				True,
				False,
			),
			false,
			nil,
		},
		test{
			"and (t,t,t) -> t",
			And(
				True,
				True,
				True,
			),
			true,
			nil,
		},
		test{
			"all (t,f,f) -> f",
			All(
				True,
				False,
				False,
			),
			false,
			nil,
		},
		test{
			"all (t,t,t) -> t",
			All(
				True,
				True,
				True,
			),
			true,
			nil,
		},
	} {
		runTestImpl(t, test)
	}
}

func TestOr(t *testing.T) {
	for _, test := range []test{
		test{
			"or (t) -> t",
			Or(
				True,
			),
			true,
			nil,
		},
		test{
			"or (f) -> f",
			Or(
				False,
			),
			false,
			nil,
		},
		test{
			"or (t,t) -> t",
			Or(
				True,
				True,
			),
			true,
			nil,
		},
		test{
			"or (t,f) -> t",
			Or(
				True,
				False,
			),
			true,
			nil,
		},
		test{
			"or (f,t) -> t",
			Or(
				False,
				True,
			),
			true,
			nil,
		},
		test{
			"or (f,f) -> f",
			Or(
				False,
				False,
			),
			false,
			nil,
		},
		test{
			"or (t,t,f) -> t",
			Or(
				True,
				True,
				False,
			),
			true,
			nil,
		},
		test{
			"or (f,f,f) -> f",
			Or(
				False,
				False,
				False,
			),
			false,
			nil,
		},
		test{
			"any (t,t,f) -> t",
			Any(
				True,
				False,
				False,
			),
			true,
			nil,
		},
		test{
			"any (f,f,f) -> f",
			Any(
				False,
				False,
				False,
			),
			false,
			nil,
		},
	} {
		runTestImpl(t, test)
	}
}

func TestNot(t *testing.T) {
	for _, test := range []test{
		test{
			"not (f) -> t",
			Not(
				False,
			),
			true,
			nil,
		},
		test{
			"not (t) -> f",
			Not(
				True,
			),
			false,
			nil,
		},
	} {
		runTestImpl(t, test)
	}
}

func TestXor(t *testing.T) {
	for _, test := range []test{
		test{
			"xor (f, f) -> f",
			Xor(
				False,
				False,
			),
			false,
			nil,
		},
		test{
			"xor (t, t) -> f",
			Xor(
				True,
				True,
			),
			false,
			nil,
		},
		test{
			"xor (f, t) -> t",
			Xor(
				False,
				True,
			),
			true,
			nil,
		},
		test{
			"xor (t, f) -> t",
			Xor(
				True,
				False,
			),
			true,
			nil,
		},
	} {
		runTestImpl(t, test)
	}
}

func TestComplex(t *testing.T) {
	for _, test := range []test{
		test{
			"not ( and ( or (t,f,f), t)) -> f",
			Not(
				And(
					Or(
						True,
						False,
						False,
					),
					True,
				),
			),
			false,
			nil,
		},
	} {
		runTestImpl(t, test)
	}
}

type failing struct{}

var failure = fmt.Errorf("internal error")

func (f failing) Evaluate(ctx interface{}) (bool, error) {
	return false, failure
}

func TestFailure(t *testing.T) {
	for _, test := range []test{
		test{
			"not ( and ( or (*,f,t), t)) -> ?, err",
			Not(
				And(
					Or(
						failing{},
						False,
						True,
					),
					True,
				),
			),
			false,
			failure,
		},
		test{
			"xor (*,t) -> ?, err",
			Xor(
				failing{},
				True,
			),
			false,
			failure,
		},
		test{
			"xor (t,*) -> ?, err",
			Xor(
				True,
				failing{},
			),
			false,
			failure,
		},
	} {
		runTestImpl(t, test)
	}
}
