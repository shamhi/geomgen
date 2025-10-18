package categories

import (
	"fmt"
	"math/rand"
)

type TrianglePoints struct {
	A [3]float64
	B [3]float64
	C [3]float64
}

type TriangleGenerator struct{}

func (g *TriangleGenerator) Category() string {
	return "triangle.median.midline"
}

func (g *TriangleGenerator) Generate(r *rand.Rand) TrianglePoints {
	return TrianglePoints{
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
		C: [3]float64{
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
		},
	}
}

func (g *TriangleGenerator) Validate(t TrianglePoints) bool {
	// Проверка, что точки не коллинеарны (не лежат на одной прямой)
	AB := [3]float64{t.B[0] - t.A[0], t.B[1] - t.A[1], t.B[2] - t.A[2]}
	AC := [3]float64{t.C[0] - t.A[0], t.C[1] - t.A[1], t.C[2] - t.A[2]}
	cross := [3]float64{
		AB[1]*AC[2] - AB[2]*AC[1],
		AB[2]*AC[0] - AB[0]*AC[2],
		AB[0]*AC[1] - AB[1]*AC[0],
	}
	lenCross := cross[0]*cross[0] + cross[1]*cross[1] + cross[2]*cross[2]
	return lenCross > 0.01
}

func (g *TriangleGenerator) Statement(t TrianglePoints) string {
	return fmt.Sprintf(
		"Даны вершины треугольника "+
			"$A(%.0f, %.0f, %.0f)$, $B(%.0f, %.0f, %.0f)$, $C(%.0f, %.0f, %.0f)$. "+
			"Составить каноническое и параметрическое уравнения средней линии, параллельной стороне BC, "+
			"и медианы, проведённой к стороне AB.",
		t.A[0], t.A[1], t.A[2],
		t.B[0], t.B[1], t.B[2],
		t.C[0], t.C[1], t.C[2],
	)
}

func (g *TriangleGenerator) Solve(t TrianglePoints) (string, error) {
	// Средняя линия к BC: проходит через середину AB
	M := [3]float64{
		(t.B[0] + t.C[0]) / 2,
		(t.B[1] + t.C[1]) / 2,
		(t.B[2] + t.C[2]) / 2,
	}
	return fmt.Sprintf(
		"Средняя линия проходит через середину BC: (%.2f, %.2f, %.2f). Медиана к AB: проходит через середину AB и точку C.",
		M[0], M[1], M[2],
	), nil
}
