package points

import (
	"fmt"
	"math"
	"math/rand"
)

type MirrorPlaneData struct {
	P     [3]float64
	Plane [4]float64 // Ax + By + Cz + D = 0
}

type MirrorPlaneGenerator struct{}

func (g *MirrorPlaneGenerator) Category() string {
	return "points.mirror_plane"
}

func (g *MirrorPlaneGenerator) Title() string {
	return "Точка, симметричная относительно плоскости"
}

func (g *MirrorPlaneGenerator) Generate(r *rand.Rand) MirrorPlaneData {
	return MirrorPlaneData{
		P:     [3]float64{float64(r.Intn(9) - 4), float64(r.Intn(9) - 4), float64(r.Intn(9) - 4)},
		Plane: [4]float64{float64(r.Intn(5) + 1), float64(r.Intn(5) + 1), float64(r.Intn(5) + 1), float64(r.Intn(7) - 3)},
	}
}

func (g *MirrorPlaneGenerator) Validate(d MirrorPlaneData) bool {
	n := math.Sqrt(d.Plane[0]*d.Plane[0] + d.Plane[1]*d.Plane[1] + d.Plane[2]*d.Plane[2])
	return n > 1e-3
}

func (g *MirrorPlaneGenerator) Statement(d MirrorPlaneData) string {
	return fmt.Sprintf(
		"Найти координаты точки, симметричной точке $M(%.0f, %.0f, %.0f)$ относительно плоскости $%.0fx %+0.0fy %+0.0fz %+0.0f = 0$.",
		d.P[0], d.P[1], d.P[2],
		d.Plane[0], d.Plane[1], d.Plane[2], d.Plane[3],
	)
}

func (g *MirrorPlaneGenerator) Solve(d MirrorPlaneData) (string, error) {
	A, B, C, D := d.Plane[0], d.Plane[1], d.Plane[2], d.Plane[3]
	x0, y0, z0 := d.P[0], d.P[1], d.P[2]
	t := -(A*x0 + B*y0 + C*z0 + D) / (A*A + B*B + C*C)
	x1 := x0 + 2*A*t
	y1 := y0 + 2*B*t
	z1 := z0 + 2*C*t
	return fmt.Sprintf("$M'(%.2f, %.2f, %.2f)$", x1, y1, z1), nil
}
