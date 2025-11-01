package lines

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type LinePair struct {
	V1 [3]float64
	V2 [3]float64
}

type LineAngleGenerator struct{}

const eps = 1e-9

func (g *LineAngleGenerator) Category() string {
	return "lines.angle"
}

func (g *LineAngleGenerator) Title() string {
	return "Угол между прямыми"
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
	lenV1 := math.Hypot(math.Hypot(v.V1[0], v.V1[1]), v.V1[2])
	lenV2 := math.Hypot(math.Hypot(v.V2[0], v.V2[1]), v.V2[2])
	if lenV1 <= eps || lenV2 <= eps {
		return false
	}

	scalar := v.V1[0]*v.V2[0] + v.V1[1]*v.V2[1] + v.V1[2]*v.V2[2]
	den := lenV1 * lenV2
	if den <= eps {
		return false
	}
	c := scalar / den
	if math.IsNaN(c) || math.IsInf(c, 0) {
		return false
	}
	return true
}

func (g *LineAngleGenerator) Statement(v LinePair) string {
	return fmt.Sprintf(
		"**Задача.** Найдите угол между прямыми, задаваемыми вектором направления "+
			"$\\vec{v}_1=(%.0f, %.0f, %.0f)$ и "+
			"$\\vec{v}_2=(%.0f, %.0f, %.0f)$.\n\n"+
			"Используйте формулу: $\\cos\\theta=\\dfrac{\\vec{v}_1\\cdot\\vec{v}_2}{|\\vec{v}_1|\\,|\\vec{v}_2|}$. ",
		v.V1[0], v.V1[1], v.V1[2],
		v.V2[0], v.V2[1], v.V2[2],
	)
}

func (g *LineAngleGenerator) Solve(v LinePair) (string, error) {
	lenV1 := math.Hypot(math.Hypot(v.V1[0], v.V1[1]), v.V1[2])
	lenV2 := math.Hypot(math.Hypot(v.V2[0], v.V2[1]), v.V2[2])
	if lenV1 <= eps || lenV2 <= eps {
		return "", errors.New("недопустимо: один из векторов нулевой (деление на ноль)")
	}

	scalar := v.V1[0]*v.V2[0] + v.V1[1]*v.V2[1] + v.V1[2]*v.V2[2]
	den := lenV1 * lenV2
	if den <= eps {
		return "", errors.New("недопустимо: произведение модулей слишком мало")
	}
	cosTheta := scalar / den
	if cosTheta > 1 {
		cosTheta = 1
	} else if cosTheta < -1 {
		cosTheta = -1
	}
	if math.IsNaN(cosTheta) || math.IsInf(cosTheta, 0) {
		return "", errors.New("числовая ошибка при вычислении косинуса")
	}

	thetaRad := math.Acos(cosTheta)
	thetaDeg := thetaRad * 180.0 / math.Pi

	solution := fmt.Sprintf(
		"1) Скалярное произведение: $\\vec{v}_1\\cdot\\vec{v}_2 = %.2f$.\\\\\n"+
			"2) Длины векторов: $|\\vec{v}_1| = %.4f$, $|\\vec{v}_2| = %.4f$.\\\\\n"+
			"3) Косинус угла: $\\displaystyle \\cos\\theta = \\frac{%.2f}{%.4f\\cdot%.4f} = %.6f$.\\\\\n"+
			"   С учётом численных корректировок: $\\cos\\theta = %.6f$.\\\\\n"+
			"4) Угол: $\\theta = \\arccos(%.6f) = %.6f\\ \\text{рад} = %.4f^{\\circ}$.\\\\\n\n"+
			"\\textbf{Ответ: } $\\theta \\approx %.4f^{\\circ}$.",
		scalar,
		lenV1, lenV2,
		scalar, lenV1, lenV2, scalar/den,
		cosTheta,
		cosTheta, thetaRad, thetaDeg,
		thetaDeg,
	)

	return solution, nil
}
