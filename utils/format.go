package utils

import (
	"fmt"
	"math"
	"strings"
)

// FormatParametric returns parametric equation string for a line through point p with direction d
func FormatParametric(p [3]float64, d [3]float64) string {
	return fmt.Sprintf("x=%.2f + t\\cdot(%.2f),\\; y=%.2f + t\\cdot(%.2f),\\; z=%.2f + t\\cdot(%.2f)", p[0], d[0], p[1], d[1], p[2], d[2])
}

// termSub formats without a double minus
func termSub(varName string, v float64) string {
	if v >= 0 {
		return fmt.Sprintf("%s-%.2f", varName, v)
	}
	return fmt.Sprintf("%s+%.2f", varName, -v)
}

// FormatCanonical returns canonical equation string for a line through point p with direction d
func FormatCanonical(p [3]float64, d [3]float64) string {
	eps := 1e-9
	ratio := make([]string, 0, 3)
	fixed := make([]string, 0, 3)
	if math.Abs(d[0]) > eps {
		ratio = append(ratio, fmt.Sprintf("\\dfrac{%s}{%.2f}", termSub("x", p[0]), d[0]))
	} else {
		fixed = append(fixed, fmt.Sprintf("x=%.2f", p[0]))
	}
	if math.Abs(d[1]) > eps {
		ratio = append(ratio, fmt.Sprintf("\\dfrac{%s}{%.2f}", termSub("y", p[1]), d[1]))
	} else {
		fixed = append(fixed, fmt.Sprintf("y=%.2f", p[1]))
	}
	if math.Abs(d[2]) > eps {
		ratio = append(ratio, fmt.Sprintf("\\dfrac{%s}{%.2f}", termSub("z", p[2]), d[2]))
	} else {
		fixed = append(fixed, fmt.Sprintf("z=%.2f", p[2]))
	}
	var b strings.Builder
	if len(ratio) > 0 {
		b.WriteString(strings.Join(ratio, " = "))
	}
	if len(fixed) > 0 {
		if b.Len() > 0 {
			b.WriteString("; ")
		}
		b.WriteString(strings.Join(fixed, ", "))
	}
	return b.String()
}
