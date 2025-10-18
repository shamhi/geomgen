package lines

import (
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

func (g *AngleLinePlaneGenerator) Generate(r *rand.Rand) LineAndPlane {
	return LineAndPlane{
		P:  [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
		D:  [3]float64{float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(7) - 3)},
		Pl: [4]float64{float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(11) - 5)},
	}
}

func (g *AngleLinePlaneGenerator) Validate(lp LineAndPlane) bool {
	lenD := math.Sqrt(lp.D[0]*lp.D[0] + lp.D[1]*lp.D[1] + lp.D[2]*lp.D[2])
	lenN := math.Abs(lp.Pl[0]) + math.Abs(lp.Pl[1]) + math.Abs(lp.Pl[2])
	return lenD > 0.01 && lenN > 1e-6
}

func (g *AngleLinePlaneGenerator) Statement(lp LineAndPlane) string {
	return fmt.Sprintf("Найти угол между прямой $\\frac{x-%.0f}{%.0f}=\\frac{y-%.0f}{%.0f}=\\frac{z-%.0f}{%.0f}$ и плоскостью $%.0fx \\; %+0.0fy \\; %+0.0fz \\; %+0.0f = 0$.", lp.P[0], lp.D[0], lp.P[1], lp.D[1], lp.P[2], lp.D[2], lp.Pl[0], lp.Pl[1], lp.Pl[2], lp.Pl[3])
}

func (g *AngleLinePlaneGenerator) Solve(lp LineAndPlane) (string, error) {
	n := [3]float64{lp.Pl[0], lp.Pl[1], lp.Pl[2]}
	lenD := math.Sqrt(lp.D[0]*lp.D[0] + lp.D[1]*lp.D[1] + lp.D[2]*lp.D[2])
	lenN := math.Sqrt(n[0]*n[0] + n[1]*n[1] + n[2]*n[2])
	cosAngle := math.Abs(n[0]*lp.D[0]+n[1]*lp.D[1]+n[2]*lp.D[2]) / (lenD * lenN)
	if cosAngle > 1 {
		cosAngle = 1
	}
	angleRad := math.Asin(cosAngle)
	phi := angleRad * 180 / math.Pi
	return fmt.Sprintf("Угол между прямой и плоскостью: %.2f^\\circ", phi), nil
}
