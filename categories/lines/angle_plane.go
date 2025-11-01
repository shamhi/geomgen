package lines

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type LineAndPlane struct {
	P  [3]float64
	D  [3]float64
	Pl [4]float64
}

type AngleLinePlaneGenerator struct{}

func (g *AngleLinePlaneGenerator) Category() string {
	return "lines.angle_plane"
}

func (g *AngleLinePlaneGenerator) Title() string {
	return "Угол между прямой и плоскостью"
}

func (g *AngleLinePlaneGenerator) Generate(r *rand.Rand) LineAndPlane {
	return LineAndPlane{
		P:  [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
		D:  [3]float64{float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(7) - 3)},
		Pl: [4]float64{float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(11) - 5)},
	}
}

func (g *AngleLinePlaneGenerator) Validate(lp LineAndPlane) bool {
	lenD := math.Sqrt(lp.D[0]*lp.D[0] + lp.D[1]*lp.D[1] + lp.D[2]*lp.D[2])
	lenN := math.Sqrt(lp.Pl[0]*lp.Pl[0] + lp.Pl[1]*lp.Pl[1] + lp.Pl[2]*lp.Pl[2])
	return lenD > 1e-6 && lenN > 1e-6
}

func (g *AngleLinePlaneGenerator) Statement(lp LineAndPlane) string {
	return fmt.Sprintf(
		"Найти угол между прямой $\\dfrac{x-%.0f}{%.0f}=\\dfrac{y-%.0f}{%.0f}=\\dfrac{z-%.0f}{%.0f}$ и плоскостью $%.0fx \\; %+0.0fy \\; %+0.0fz \\; %+0.0f = 0$.",
		lp.P[0], lp.D[0], lp.P[1], lp.D[1], lp.P[2], lp.D[2],
		lp.Pl[0], lp.Pl[1], lp.Pl[2], lp.Pl[3],
	)
}

func (g *AngleLinePlaneGenerator) Solve(lp LineAndPlane) (string, error) {
	n := [3]float64{lp.Pl[0], lp.Pl[1], lp.Pl[2]}
	lenD := math.Sqrt(lp.D[0]*lp.D[0] + lp.D[1]*lp.D[1] + lp.D[2]*lp.D[2])
	lenN := math.Sqrt(n[0]*n[0] + n[1]*n[1] + n[2]*n[2])
	if lenD < 1e-6 || lenN < 1e-6 {
		return "", errors.New("некорректные данные: нулевой вектор")
	}
	scalar := math.Abs(n[0]*lp.D[0] + n[1]*lp.D[1] + n[2]*lp.D[2])
	cosAlpha := scalar / (lenD * lenN)
	if cosAlpha > 1 {
		cosAlpha = 1
	}
	alpha := math.Asin(cosAlpha)
	alphaDeg := alpha * 180 / math.Pi

	return fmt.Sprintf(
		"1) Нормальный вектор плоскости: $\\vec{n}=(%.0f, %.0f, %.0f)$.\\\\\n"+
			"2) Направляющий вектор прямой: $\\vec{d}=(%.0f, %.0f, %.0f)$.\\\\\n"+
			"3) Косинус угла между $\\vec{n}$ и $\\vec{d}$: $\\cos\\alpha=\\dfrac{|%.2f|}{%.4f\\cdot%.4f}=%.5f$.\\\\\n"+
			"4) Угол между прямой и плоскостью: $\\varphi=90^{\\circ}-\\alpha=\\arcsin(\\cos\\alpha)=%.4f^{\\circ}$.\\\\\n"+
			"\\textbf{Ответ: } $\\varphi\\approx%.4f^{\\circ}$.",
		n[0], n[1], n[2],
		lp.D[0], lp.D[1], lp.D[2],
		scalar, lenD, lenN, cosAlpha,
		alphaDeg, alphaDeg,
	), nil
}
