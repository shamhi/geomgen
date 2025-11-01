package lines

import (
	"errors"
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

func (g *LinePerpPlaneGenerator) Title() string {
	return "Прямая ⟂ плоскости через точку"
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
	return fmt.Sprintf(
		"Составить уравнение прямой, проходящей через точку $A(%.0f, %.0f, %.0f)$ и перпендикулярной плоскости $%.0fx \\; %+0.0fy \\; %+0.0fz \\; %+0.0f = 0$.",
		pp.P[0], pp.P[1], pp.P[2],
		pp.Plane[0], pp.Plane[1], pp.Plane[2], pp.Plane[3],
	)
}

func (g *LinePerpPlaneGenerator) Solve(pp PointPlane) (string, error) {
	if math.Abs(pp.Plane[0])+math.Abs(pp.Plane[1])+math.Abs(pp.Plane[2]) < 1e-6 {
		return "", errors.New("некорректная плоскость (нулевой нормальный вектор)")
	}
	d := [3]float64{pp.Plane[0], pp.Plane[1], pp.Plane[2]}
	param := utils.FormatParametric(pp.P, d)
	canon := utils.FormatCanonical(pp.P, d)

	return fmt.Sprintf(
		"1) Нормальный вектор плоскости: $\\vec{n}=(%.0f, %.0f, %.0f)$ — это и есть направляющий вектор прямой.\\\\\n"+
			"2) Прямая проходит через $A(%.0f, %.0f, %.0f)$.\\\\\n"+
			"3) Параметрические уравнения прямой: $%s$.\\\\\n"+
			"4) Каноническое уравнение: $%s$.\\\\\n"+
			"\\textbf{Ответ: } прямая $A\\vec{n}$ перпендикулярна плоскости.",
		d[0], d[1], d[2],
		pp.P[0], pp.P[1], pp.P[2],
		param, canon,
	), nil
}
