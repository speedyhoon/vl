package vl

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/speedyhoon/frm"
)

// Float32 validates inp as a float32 input.
func Float32(f *frm.Field, inp ...string) {
	f64, err := strconv.ParseFloat(strings.TrimSpace(inp[0]), 32)
	if err != nil {
		// Return error if input string failed to convert.
		f.Err = err.Error()
		return
	}
	num := float32(f64)
	f.Value = num

	if !f.Required && num == 0 {
		// f.ValueFloat32 is zero by default so assigning zero isn't required.
		return
	}

	if f.Min != nil && num < f.Min.(float32) || f.Max != nil && num > f.Max.(float32) {
		f.Err = fmt.Sprintf("Must be between %f and %f.", f.Min, f.Max)
		return
	}

	if rem := toFixed32(math.Mod(f64, float64(f.Step)), 6); rem != 0 {
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %f and %f.", num-rem, num-rem+f.Step)
	}
}

// Float64 validates inp as a float64 input
func Float64(f *frm.Field, inp ...string) {
	num, err := strconv.ParseFloat(strings.TrimSpace(inp[0]), 64)
	f.Value = num
	if err != nil {
		// Return error if input string failed to convert.
		f.Err = err.Error()
		return
	}

	if !f.Required && num == 0 {
		// f.ValueFloat64 is zero by default so assigning zero isn't required.
		return
	}

	if f.Min != nil && num < f.Min.(float64) || f.Max != nil && num > f.Max.(float64) {
		f.Err = fmt.Sprintf("Must be between %f and %f.", f.Min, f.Max)
		return
	}

	if rem := toFixed(math.Mod(num, float64(f.Step)), 6); rem != 0 {
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %f and %f.", num-rem, num-rem+float64(f.Step))
	}
}

func toFixed32(num, precision float64) float32 {
	output := math.Pow(10, precision)
	return float32(int(num*output)) / float32(output)
}

func toFixed(num, precision float64) float64 {
	output := math.Pow(10, precision)
	return float64(int(num*output)) / output
}
