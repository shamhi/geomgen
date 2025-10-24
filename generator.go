package geomgen

import (
	"math/rand"
	"time"
)

// ExpressionGenerator[T] base interface for all generators
type ExpressionGenerator[T any] interface {
	Category() string
	Title() string
	Generate(r *rand.Rand) T
	Validate(expr T) bool
	Statement(expr T) string
	Solve(expr T) (string, error)
}

// GenerateValidExpression[T] main method for generating expressions based on `seed`
func GenerateValidExpression[T any](gen ExpressionGenerator[T], seed string) Expression[T] {
	r := NewRand(seed)
	const maxAttempts = 10000
	for i := 0; i < maxAttempts; i++ {
		expr := gen.Generate(r)
		if !gen.Validate(expr) {
			continue
		}
		statement := gen.Statement(expr)
		solution, err := gen.Solve(expr)
		if err != nil {
			continue
		}
		return Expression[T]{
			Category:  gen.Category(),
			Title:     gen.Title(),
			Data:      expr,
			Statement: statement,
			Solution:  solution,
			Valid:     true,
			Seed:      seed,
			CreatedAt: time.Now(),
		}
	}
	return Expression[T]{
		Category:  gen.Category(),
		Title:     gen.Title(),
		Valid:     false,
		Seed:      seed,
		CreatedAt: time.Now(),
	}
}
