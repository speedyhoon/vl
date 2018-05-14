package v8

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/speedyhoon/forms"
)

//Float32 validates inp as a float32 input
func Float32(f *forms.Field, inp ...string) {
	f64, err := strconv.ParseFloat(strings.TrimSpace(inp[0]), 32)
	if err != nil {
		//Return error if input string failed to convert.
		f.Err = err.Error()
		return
	}
	num := float32(f64)

	if !f.Required && num == 0 {
		//f.ValueFloat32 is zero by default so assigning zero isn't required
		return
	}
	if num < float32(f.Min) || num > float32(f.Max) {
		f.Err = fmt.Sprintf("Must be between %v and %v.", f.Min, f.Max)
		return
	}

	if rem := toFixed(math.Mod(f64, float64(f.Step)), 6); rem != 0 {
		f.Err = fmt.Sprintf("Please enter a valid value. The two nearest values are %v and %v.", num-rem, num-rem+f.Step)
		return
	}
	f.Value = num
}

func toFixed(num, precision float64) float32 {
	output := math.Pow(10, precision)
	return float32(int(num*output)) / float32(output)
}
