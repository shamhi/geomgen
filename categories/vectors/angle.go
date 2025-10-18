package vectors

import (
	"fmt"
	"math"
	"math/rand"
)

type VectorPair struct {
	A [3]float64
	B [3]float64
}

type VectorAngleGenerator struct{}

func (g *VectorAngleGenerator) Category() string {
	return "vectors.angle"
}

func (g *VectorAngleGenerator) Generate(r *rand.Rand) VectorPair {
	return VectorPair{
		A: [3]float64{
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
		},
		B: [3]float64{
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
		},
	}
}

func (g *VectorAngleGenerator) Validate(v VectorPair) bool {
	lenA := math.Sqrt(v.A[0]*v.A[0] + v.A[1]*v.A[1] + v.A[2]*v.A[2])
	lenB := math.Sqrt(v.B[0]*v.B[0] + v.B[1]*v.B[1] + v.B[2]*v.B[2])
	return lenA > 0.01 && lenB > 0.01
}

func (g *VectorAngleGenerator) Statement(v VectorPair) string {
	return fmt.Sprintf(
		"Найти угол между векторами "+
			"$\\mathbf{a} = (%.0f, %.0f, %.0f)$ и "+
			"$\\mathbf{b} = (%.0f, %.0f, %.0f)$.",
		v.A[0], v.A[1], v.A[2],
		v.B[0], v.B[1], v.B[2],
	)
}

func (g *VectorAngleGenerator) Solve(v VectorPair) (string, error) {
	scalar := v.A[0]*v.B[0] + v.A[1]*v.B[1] + v.A[2]*v.B[2]
	lenA := math.Sqrt(v.A[0]*v.A[0] + v.A[1]*v.A[1] + v.A[2]*v.A[2])
	lenB := math.Sqrt(v.B[0]*v.B[0] + v.B[1]*v.B[1] + v.B[2]*v.B[2])
	c := scalar / (lenA * lenB)
	if c > 1 {
		c = 1
	} else if c < -1 {
		c = -1
	}
	angle := math.Acos(c) * 180 / math.Pi

	return fmt.Sprintf(
		"$\\cos(\\theta) = \\frac{%.2f}{%.2f \\cdot %.2f}$ "+
			"$\\Rightarrow \\theta = %.2f^{\\circ}$",
		scalar, lenA, lenB, angle,
	), nil
}
