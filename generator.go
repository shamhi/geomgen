package geomgen

import (
	"math/rand"
	"time"
)

type ExpressionGenerator[T any] interface {
	Category() string
	Generate(r *rand.Rand) T
	Validate(expr T) bool
	Statement(expr T) string
	Solve(expr T) (string, error)
}

func GenerateValidExpression[T any](gen ExpressionGenerator[T], seed string) Expression[T] {
	r := NewRand(seed)
	for {
		expr := gen.Generate(r)
		if gen.Validate(expr) {
			statement := gen.Statement(expr)
			solution, err := gen.Solve(expr)
			if err != nil {
				continue
			}

			return Expression[T]{
				Category:  gen.Category(),
				Data:      expr,
				Statement: statement,
				Solution:  solution,
				Valid:     true,
				Seed:      seed,
				CreatedAt: time.Now(),
			}
		}
	}
}
