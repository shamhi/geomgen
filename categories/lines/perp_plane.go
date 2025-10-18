package lines

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type PointPlane struct {
	P     [3]float64
	Plane [4]float64
}

type LinePerpPlaneGenerator struct{}

func (g *LinePerpPlaneGenerator) Category() string {
	return "lines.perp_plane"
}

func (g *LinePerpPlaneGenerator) Generate(r *rand.Rand) PointPlane {
	pp := PointPlane{
		P:     [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
		Plane: [4]float64{float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(11) - 5)},
	}
	if math.Abs(pp.Plane[0])+math.Abs(pp.Plane[1])+math.Abs(pp.Plane[2]) < 1e-6 {
		pp.Plane[0] = 1
	}
	return pp
}

func (g *LinePerpPlaneGenerator) Validate(pp PointPlane) bool {
	return math.Abs(pp.Plane[0])+math.Abs(pp.Plane[1])+math.Abs(pp.Plane[2]) > 1e-6
}

func (g *LinePerpPlaneGenerator) Statement(pp PointPlane) string {
	return fmt.Sprintf("Составить уравнение прямой, проходящей через точку $A(%.0f, %.0f, %.0f)$ перпендикулярно плоскости $%.0fx \\; %+0.0fy \\; %+0.0fz \\; %+0.0f = 0$.", pp.P[0], pp.P[1], pp.P[2], pp.Plane[0], pp.Plane[1], pp.Plane[2], pp.Plane[3])
}

func (g *LinePerpPlaneGenerator) Solve(pp PointPlane) (string, error) {
	d := [3]float64{pp.Plane[0], pp.Plane[1], pp.Plane[2]}
	param := formatParametric(pp.P, d)
	canon := formatCanonical(pp.P, d)
	return fmt.Sprintf("Параметрическое: $%s$; каноническое: $%s$", param, canon), nil
}

func formatParametric(p [3]float64, d [3]float64) string {
	return fmt.Sprintf("x=%.2f + t\\cdot(%.2f),\\; y=%.2f + t\\cdot(%.2f),\\; z=%.2f + t\\cdot(%.2f)", p[0], d[0], p[1], d[1], p[2], d[2])
}

func formatCanonical(p [3]float64, d [3]float64) string {
	eps := 1e-9
	ratio := make([]string, 0, 3)
	fixed := make([]string, 0, 3)
	if math.Abs(d[0]) > eps {
		ratio = append(ratio, fmt.Sprintf("\\dfrac{x-%.2f}{%.2f}", p[0], d[0]))
	} else {
		fixed = append(fixed, fmt.Sprintf("x=%.2f", p[0]))
	}
	if math.Abs(d[1]) > eps {
		ratio = append(ratio, fmt.Sprintf("\\dfrac{y-%.2f}{%.2f}", p[1], d[1]))
	} else {
		fixed = append(fixed, fmt.Sprintf("y=%.2f", p[1]))
	}
	if math.Abs(d[2]) > eps {
		ratio = append(ratio, fmt.Sprintf("\\dfrac{z-%.2f}{%.2f}", p[2], d[2]))
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
