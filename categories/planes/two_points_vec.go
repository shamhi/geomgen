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

type PlaneThroughTwoPointsParallelVectorGenerator struct{}

func (g *PlaneThroughTwoPointsParallelVectorGenerator) Category() string {
	return "planes.two_points_vec"
}

func (g *PlaneThroughTwoPointsParallelVectorGenerator) Generate(r *rand.Rand) TwoPointsVector {
	return TwoPointsVector{
		A: [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
		B: [3]float64{float64(r.Intn(11) - 5), float64(r.Intn(11) - 5), float64(r.Intn(11) - 5)},
		V: [3]float64{float64(r.Intn(7) - 3), float64(r.Intn(7) - 3), float64(r.Intn(7) - 3)},
	}
}

func (g *PlaneThroughTwoPointsParallelVectorGenerator) Validate(p TwoPointsVector) bool {
	lenV := math.Sqrt(p.V[0]*p.V[0] + p.V[1]*p.V[1] + p.V[2]*p.V[2])
	diff := [3]float64{p.B[0] - p.A[0], p.B[1] - p.A[1], p.B[2] - p.A[2]}
	lenDiff := math.Sqrt(diff[0]*diff[0] + diff[1]*diff[1] + diff[2]*diff[2])
	if lenV < 0.01 || lenDiff < 0.01 {
		return false
	}
	cross := [3]float64{
		diff[1]*p.V[2] - diff[2]*p.V[1],
		diff[2]*p.V[0] - diff[0]*p.V[2],
		diff[0]*p.V[1] - diff[1]*p.V[0],
	}
	lenCross := math.Sqrt(cross[0]*cross[0] + cross[1]*cross[1] + cross[2]*cross[2])
	return lenCross > 0.01
}

func (g *PlaneThroughTwoPointsParallelVectorGenerator) Statement(p TwoPointsVector) string {
	return fmt.Sprintf("Составить уравнение плоскости, проходящей через точки $A(%.0f,%.0f,%.0f)$ и $B(%.0f,%.0f,%.0f)$, параллельно вектору $\\vec{a}=(%.0f,%.0f,%.0f)$.", p.A[0], p.A[1], p.A[2], p.B[0], p.B[1], p.B[2], p.V[0], p.V[1], p.V[2])
}

func (g *PlaneThroughTwoPointsParallelVectorGenerator) Solve(p TwoPointsVector) (string, error) {
	u := [3]float64{p.B[0] - p.A[0], p.B[1] - p.A[1], p.B[2] - p.A[2]}
	v := p.V
	n := [3]float64{
		u[1]*v[2] - u[2]*v[1],
		u[2]*v[0] - u[0]*v[2],
		u[0]*v[1] - u[1]*v[0],
	}
	D := -(n[0]*p.A[0] + n[1]*p.A[1] + n[2]*p.A[2])
	return fmt.Sprintf("$%.0f x \\; %+0.0f y \\; %+0.0f z \\; %+0.0f = 0$", n[0], n[1], n[2], D), nil
}
