package triangles

import (
	"fmt"
	"math/rand"

	"github.com/shamhi/geomgen/utils"
)

type TrianglePoints struct {
	A [3]float64
	B [3]float64
	C [3]float64
}

type TriangleGenerator struct{}

func (g *TriangleGenerator) Category() string {
	return "triangles.midline"
}

func (g *TriangleGenerator) Title() string {
	return "Средняя линия и медиана треугольника"
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
	Mab := [3]float64{
		(t.A[0] + t.B[0]) / 2,
		(t.A[1] + t.B[1]) / 2,
		(t.A[2] + t.B[2]) / 2,
	}
	Mac := [3]float64{
		(t.A[0] + t.C[0]) / 2,
		(t.A[1] + t.C[1]) / 2,
		(t.A[2] + t.C[2]) / 2,
	}
	dMid := [3]float64{Mac[0] - Mab[0], Mac[1] - Mab[1], Mac[2] - Mab[2]}
	dMed := [3]float64{Mab[0] - t.C[0], Mab[1] - t.C[1], Mab[2] - t.C[2]}
	paramMid := utils.FormatParametric(Mab, dMid)
	canonMid := utils.FormatCanonical(Mab, dMid)
	paramMed := utils.FormatParametric(t.C, dMed)
	canonMed := utils.FormatCanonical(t.C, dMed)
	return fmt.Sprintf("Средняя линия: параметрическое: $%s$; каноническое: $%s$. Медиана к AB: параметрическое: $%s$; каноническое: $%s$.", paramMid, canonMid, paramMed, canonMed), nil
}
