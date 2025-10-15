package categories

import (
	"fmt"
	"math"
	"math/rand"
)

func GenerateVectorAngle(r *rand.Rand) (string, string) {
	a := [3]int{r.Intn(11) - 5, r.Intn(11) - 5, r.Intn(11) - 5}
	b := [3]int{r.Intn(11) - 5, r.Intn(11) - 5, r.Intn(11) - 5}
	scalar := a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
	lenA := math.Sqrt(float64(a[0]*a[0] + a[1]*a[1] + a[2]*a[2]))
	lenB := math.Sqrt(float64(b[0]*b[0] + b[1]*b[1] + b[2]*b[2]))
	angle := math.Acos(float64(scalar)/(lenA*lenB)) * 180 / math.Pi

	statement := fmt.Sprintf("Найти угол между векторами a=(%d,%d,%d) и b=(%d,%d,%d).",
		a[0], a[1], a[2], b[0], b[1], b[2])

	solution := fmt.Sprintf("cos(θ) = %d / (%.2f × %.2f) ⇒ θ = %.2f°",
		scalar, lenA, lenB, angle)

	return statement, solution
}
