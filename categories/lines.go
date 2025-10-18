package categories

import (
	"fmt"
	"math"
	"math/rand"
)

type LinePair struct {
	V1 [3]float64
	V2 [3]float64
}

type LineAngleGenerator struct{}

func (g *LineAngleGenerator) Category() string {
	return "lines.angle"
}

func (g *LineAngleGenerator) Generate(r *rand.Rand) LinePair {
	return LinePair{
		V1: [3]float64{
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
		},
		V2: [3]float64{
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
			float64(r.Intn(11) - 5),
		},
	}
}

func (g *LineAngleGenerator) Validate(v LinePair) bool {
	lenV1 := math.Sqrt(v.V1[0]*v.V1[0] + v.V1[1]*v.V1[1] + v.V1[2]*v.V1[2])
	lenV2 := math.Sqrt(v.V2[0]*v.V2[0] + v.V2[1]*v.V2[1] + v.V2[2]*v.V2[2])
	return lenV1 > 0.01 && lenV2 > 0.01
}

func (g *LineAngleGenerator) ToMarkdown(v LinePair) string {
	return fmt.Sprintf(
		"**Задача.** Найти угол между прямыми с векторами направлений "+
			"$\\vec{v_1} = (%.0f, %.0f, %.0f)$ и "+
			"$\\vec{v_2} = (%.0f, %.0f, %.0f)$.",
		v.V1[0], v.V1[1], v.V1[2],
		v.V2[0], v.V2[1], v.V2[2],
	)
}

func (g *LineAngleGenerator) Solve(v LinePair) (string, error) {
	scalar := v.V1[0]*v.V2[0] + v.V1[1]*v.V2[1] + v.V1[2]*v.V2[2]
	lenV1 := math.Sqrt(v.V1[0]*v.V1[0] + v.V1[1]*v.V1[1] + v.V1[2]*v.V1[2])
	lenV2 := math.Sqrt(v.V2[0]*v.V2[0] + v.V2[1]*v.V2[1] + v.V2[2]*v.V2[2])
	angle := math.Acos(scalar/(lenV1*lenV2)) * 180 / math.Pi

	return fmt.Sprintf(
		"$\\cos(\\theta) = \\frac{%.2f}{%.2f \\cdot %.2f} \\Rightarrow \\theta = %.2f^{\\circ}$",
		scalar, lenV1, lenV2, angle,
	), nil
}
