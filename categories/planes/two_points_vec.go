package planes

import (
	"fmt"
	"math"
	"math/rand"
)

type TwoPointsVector struct {
	A [3]float64
	B [3]float64
	V [3]float64
}

type PlaneTwoPointsParallelVectorGenerator struct{}

func (g *PlaneTwoPointsParallelVectorGenerator) Category() string {
	return "planes.two_points_vec"
}

func (g *PlaneTwoPointsParallelVectorGenerator) Title() string {
	return "Плоскость через две точки ∥ заданному вектору"
}

func (g *PlaneTwoPointsParallelVectorGenerator) Generate(r *rand.Rand) TwoPointsVector {
	return TwoPointsVector{
		A: [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
		B: [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
		V: [3]float64{float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(7) - 3)},
	}
}

func (g *PlaneTwoPointsParallelVectorGenerator) Validate(p TwoPointsVector) bool {
	u := [3]float64{p.B[0] - p.A[0], p.B[1] - p.A[1], p.B[2] - p.A[2]}
	lenU := math.Sqrt(u[0]*u[0] + u[1]*u[1] + u[2]*u[2])
	lenV := math.Sqrt(p.V[0]*p.V[0] + p.V[1]*p.V[1] + p.V[2]*p.V[2])
	if lenU < 1e-3 || lenV < 1e-3 {
		return false
	}

	cross := [3]float64{
		u[1]*p.V[2] - u[2]*p.V[1],
		u[2]*p.V[0] - u[0]*p.V[2],
		u[0]*p.V[1] - u[1]*p.V[0],
	}
	lenCross := math.Sqrt(cross[0]*cross[0] + cross[1]*cross[1] + cross[2]*cross[2])
	return lenCross > 1e-3
}

func (g *PlaneTwoPointsParallelVectorGenerator) Statement(p TwoPointsVector) string {
	return fmt.Sprintf(
		"Составить уравнение плоскости, проходящей через точки "+
			"$A(%.0f, %.0f, %.0f)$ и $B(%.0f, %.0f, %.0f)$, параллельно вектору "+
			"$\\vec{a} = (%.0f, %.0f, %.0f)$.",
		p.A[0], p.A[1], p.A[2],
		p.B[0], p.B[1], p.B[2],
		p.V[0], p.V[1], p.V[2],
	)
}

func (g *PlaneTwoPointsParallelVectorGenerator) Solve(p TwoPointsVector) (string, error) {
	u := [3]float64{p.B[0] - p.A[0], p.B[1] - p.A[1], p.B[2] - p.A[2]}
	v := p.V
	n := [3]float64{
		u[1]*v[2] - u[2]*v[1],
		u[2]*v[0] - u[0]*v[2],
		u[0]*v[1] - u[1]*v[0],
	}
	D := -(n[0]*p.A[0] + n[1]*p.A[1] + n[2]*p.A[2])

	return fmt.Sprintf(
		"1. Направляющие вектора плоскости: $\\vec{AB} = (%.0f, %.0f, %.0f)$ и $\\vec{a} = (%.0f, %.0f, %.0f)$. \n"+
			"2. Нормаль $\\vec{n} = [\\vec{AB}, \\vec{a}] = (%.0f, %.0f, %.0f)$. \n"+
			"3. Уравнение плоскости: $%.0fx %+0.0fy %+0.0fz %+0.0f = 0$.",
		u[0], u[1], u[2], v[0], v[1], v[2],
		n[0], n[1], n[2],
		n[0], n[1], n[2], D,
	), nil
}
