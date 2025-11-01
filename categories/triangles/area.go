package triangles

import (
	"fmt"
	"math"
	"math/rand"
)

type TriangleAreaPoints struct {
	A [3]float64
	B [3]float64
	C [3]float64
}

type TriangleAreaGenerator struct{}

func (g *TriangleAreaGenerator) Category() string {
	return "triangles.area"
}

func (g *TriangleAreaGenerator) Title() string {
	return "Площадь треугольника по трём вершинам"
}

func (g *TriangleAreaGenerator) Generate(r *rand.Rand) TriangleAreaPoints {
	return TriangleAreaPoints{
		A: [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
		B: [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
		C: [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
	}
}

func (g *TriangleAreaGenerator) Validate(t TriangleAreaPoints) bool {
	AB := [3]float64{t.B[0] - t.A[0], t.B[1] - t.A[1], t.B[2] - t.A[2]}
	AC := [3]float64{t.C[0] - t.A[0], t.C[1] - t.A[1], t.C[2] - t.A[2]}
	cross := [3]float64{
		AB[1]*AC[2] - AB[2]*AC[1],
		AB[2]*AC[0] - AB[0]*AC[2],
		AB[0]*AC[1] - AB[1]*AC[0],
	}
	lenCross := math.Sqrt(cross[0]*cross[0] + cross[1]*cross[1] + cross[2]*cross[2])
	return lenCross > 0.01
}

func (g *TriangleAreaGenerator) Statement(t TriangleAreaPoints) string {
	return fmt.Sprintf(
		"Найти площадь треугольника с вершинами $A(%.0f, %.0f, %.0f)$, $B(%.0f, %.0f, %.0f)$, $C(%.0f, %.0f, %.0f)$.",
		t.A[0], t.A[1], t.A[2],
		t.B[0], t.B[1], t.B[2],
		t.C[0], t.C[1], t.C[2],
	)
}

func (g *TriangleAreaGenerator) Solve(t TriangleAreaPoints) (string, error) {
	AB := [3]float64{t.B[0] - t.A[0], t.B[1] - t.A[1], t.B[2] - t.A[2]}
	AC := [3]float64{t.C[0] - t.A[0], t.C[1] - t.A[1], t.C[2] - t.A[2]}
	cross := [3]float64{
		AB[1]*AC[2] - AB[2]*AC[1],
		AB[2]*AC[0] - AB[0]*AC[2],
		AB[0]*AC[1] - AB[1]*AC[0],
	}
	S := 0.5 * math.Sqrt(cross[0]*cross[0]+cross[1]*cross[1]+cross[2]*cross[2])
	return fmt.Sprintf("$S = \\frac{1}{2}|[\\overrightarrow{AB}, \\overrightarrow{AC}]| = %.2f$", S), nil
}
