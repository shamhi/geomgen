package geomgen

import (
	"fmt"
	"time"

	"github.com/shamhi/geomgen/categories"
)

func Generate(category Category, seed string) Task {
	r := NewRand(seed)
	var statement, solution string
	switch category {
	case Vectors:
		statement, solution = categories.GenerateVectorAngle(r)
	default:
		statement, solution = "Unknown category", ""
	}
	return Task{
		ID:        fmt.Sprintf("%x", SeedFromString(seed)),
		Category:  string(category),
		Statement: statement,
		Solution:  solution,
		Seed:      seed,
		CreatedAt: time.Now(),
	}
}
