// Copyright 2020 Paul Greenberg (greenpau@outlook.com)

package calculator

import (
	"fmt"
	"testing"
)

type TestInput struct {
	description string
	input       []uint64
	reg         Register
	shouldFail  bool // Whether test should result in a failure
	shouldErr   bool // Whether parsing of a response should result in error
	errMessage  string
}

func evalTestResults(t *testing.T, i int, k string, test TestInput, output *Cell, err error) (bool, error) {
	if !test.shouldErr {
		return true, fmt.Errorf(
			"FAIL: Test %d: input (%s): input: %v, register: %v, output register: %v, error: %v",
			i, k, test.input, test.reg, output.Register, err,
		)
	}
	if test.errMessage != err.Error() {
		return true, fmt.Errorf(
			"FAIL: Test %d: input (%s): %v, expected different error: %s (expected) vs. %s (received)",
			i, k, test.input, test.errMessage, err,
		)
	}
	if test.shouldFail {
		return true, fmt.Errorf(
			"FAIL: Test %d: input (%s): %v, expected failure but passed",
			i, k, test.input,
		)
	}

	t.Logf("PASS: Test %d: input (%s): %v", i, k, test)
	return false, nil
}

func TestUint64Calculator(t *testing.T) {
	testFailed := 0

	for i, test := range []TestInput{
		{
			input: []uint64{1, 2, 3, 4, 5, 6},
			reg: Register{
				Total:             21.0,
				Mean:              3.5,
				Median:            3.5,
				Range:             5.0,
				Variance:          2.9166666666666665,
				StandardDeviation: 1.0,
				SortedMedian:      3.5,
				MaxIndices:        []int{5},
				MaxValue:          6.0,
				MinIndices:        []int{0},
				MinValue:          1.0,
				Modes:             []float64{1.0},
				ModeRepeatCount:   0,
			},
			shouldFail: false,
			shouldErr:  false,
			errMessage: "",
		},
	} {
		calc := NewUint64(test.input)
		calc.RunAll()

		if calc.Register.Total != test.reg.Total {
			abort, err := evalTestResults(t, i, "total", test, calc, nil)
			if err != nil {
				t.Logf("%s", err)
				testFailed++
			}
			if abort {
				continue
			}
		}
		if calc.Register.Mean != test.reg.Mean {
			abort, err := evalTestResults(t, i, "mean", test, calc, nil)
			if err != nil {
				t.Logf("%s", err)
				testFailed++
			}
			if abort {
				continue
			}
		}
		if calc.Register.SortedMedian != test.reg.SortedMedian {
			abort, err := evalTestResults(t, i, "sorted_median", test, calc, nil)
			if err != nil {
				t.Logf("%s", err)
				testFailed++
			}
			if abort {
				continue
			}
		}
		if calc.Register.Median != test.reg.Median {
			abort, err := evalTestResults(t, i, "median", test, calc, nil)
			if err != nil {
				t.Logf("%s", err)
				testFailed++
			}
			if abort {
				continue
			}
		}
		if calc.Register.Variance != test.reg.Variance {
			abort, err := evalTestResults(t, i, "variance", test, calc, nil)
			if err != nil {
				t.Logf("%s", err)
				testFailed++
			}
			if abort {
				continue
			}
		}
		if calc.Register.Range != test.reg.Range {
			abort, err := evalTestResults(t, i, "range", test, calc, nil)
			if err != nil {
				t.Logf("%s", err)
				testFailed++
			}
			if abort {
				continue
			}
		}

		if calc.Register.MaxValue != test.reg.MaxValue {
			abort, err := evalTestResults(t, i, "max_value", test, calc, nil)
			if err != nil {
				t.Logf("%s", err)
				testFailed++
			}
			if abort {
				continue
			}
		}

		if calc.Register.MinValue != test.reg.MinValue {
			abort, err := evalTestResults(t, i, "min_value", test, calc, nil)
			if err != nil {
				t.Logf("%s", err)
				testFailed++
			}
			if abort {
				continue
			}
		}

		t.Logf("PASS: Test %d: input: %v", i, test.input)

	}
	if testFailed > 0 {
		t.Fatalf("Failed %d tests", testFailed)
	}
}