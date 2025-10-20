package lines

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/shamhi/geomgen/utils"
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
	param := utils.FormatParametric(pp.P, d)
	canon := utils.FormatCanonical(pp.P, d)
	return fmt.Sprintf("Параметрическое: $%s$; каноническое: $%s$", param, canon), nil
}
