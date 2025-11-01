package coordinates

import (
	"fmt"
	"math"
	"math/rand"
)

type ChangeBasisData struct {
	P  [3]float64
	E1 [3]float64
	E2 [3]float64
	E3 [3]float64
}

type ChangeBasisGenerator struct{}

func (g *ChangeBasisGenerator) Category() string {
	return "coordinates.change_basis"
}

func (g *ChangeBasisGenerator) Title() string {
	return "Переход к новому базису"
}

func (g *ChangeBasisGenerator) Generate(r *rand.Rand) ChangeBasisData {
	return ChangeBasisData{
		P:  [3]float64{float64(r.Intn(9) - 4), float64(r.Intn(9) - 4), float64(r.Intn(9) - 4)},
		E1: [3]float64{1, 0, float64(r.Intn(3))},
		E2: [3]float64{0, 1, float64(r.Intn(3))},
		E3: [3]float64{float64(r.Intn(3)), float64(r.Intn(3)), 1},
	}
}

func (g *ChangeBasisGenerator) Validate(d ChangeBasisData) bool {
	det := d.E1[0]*(d.E2[1]*d.E3[2]-d.E2[2]*d.E3[1]) -
		d.E1[1]*(d.E2[0]*d.E3[2]-d.E2[2]*d.E3[0]) +
		d.E1[2]*(d.E2[0]*d.E3[1]-d.E2[1]*d.E3[0])
	return math.Abs(det) > 1e-3
}

func (g *ChangeBasisGenerator) Statement(d ChangeBasisData) string {
	return fmt.Sprintf(
		"Дана точка $M(%.0f, %.0f, %.0f)$ и новый базис с векторами "+
			"$\\vec{e_1}=(%.0f, %.0f, %.0f)$, $\\vec{e_2}=(%.0f, %.0f, %.0f)$, $\\vec{e_3}=(%.0f, %.0f, %.0f)$. "+
			"Найти координаты точки $M$ в новом базисе.",
		d.P[0], d.P[1], d.P[2],
		d.E1[0], d.E1[1], d.E1[2],
		d.E2[0], d.E2[1], d.E2[2],
		d.E3[0], d.E3[1], d.E3[2],
	)
}

func (g *ChangeBasisGenerator) Solve(d ChangeBasisData) (string, error) {
	A := [3][3]float64{
		{d.E1[0], d.E2[0], d.E3[0]},
		{d.E1[1], d.E2[1], d.E3[1]},
		{d.E1[2], d.E2[2], d.E3[2]},
	}
	b := d.P
	detA := A[0][0]*(A[1][1]*A[2][2]-A[1][2]*A[2][1]) -
		A[0][1]*(A[1][0]*A[2][2]-A[1][2]*A[2][0]) +
		A[0][2]*(A[1][0]*A[2][1]-A[1][1]*A[2][0])
	if math.Abs(detA) < 1e-6 {
		return "", fmt.Errorf("determinant is zero")
	}

	detX := b[0]*(A[1][1]*A[2][2]-A[1][2]*A[2][1]) -
		A[0][1]*(b[1]*A[2][2]-A[1][2]*b[2]) +
		A[0][2]*(b[1]*A[2][1]-A[1][1]*b[2])
	detY := A[0][0]*(b[1]*A[2][2]-A[1][2]*b[2]) -
		b[0]*(A[1][0]*A[2][2]-A[1][2]*A[2][0]) +
		A[0][2]*(A[1][0]*b[2]-b[1]*A[2][0])
	detZ := A[0][0]*(A[1][1]*b[2]-b[1]*A[2][1]) -
		A[0][1]*(A[1][0]*b[2]-b[1]*A[2][0]) +
		b[0]*(A[1][0]*A[2][1]-A[1][1]*A[2][0])
	x, y, z := detX/detA, detY/detA, detZ/detA
	return fmt.Sprintf("$[M]_{new} = (%.2f, %.2f, %.2f)$", x, y, z), nil
}
