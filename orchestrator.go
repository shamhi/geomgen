package geomgen

import (
	"fmt"
)

// GenerateWork assembles a work per configuration using the default registry
func GenerateWork(cfg WorkConfig) (WorkResult, error) {
	if cfg.Seed == "" {
		cfg.Seed = "default"
	}
	result := WorkResult{Config: cfg}
	for itemIndex, it := range cfg.Items {
		gen, ok := DefaultRegistry.Get(it.Key)
		if !ok {
			return WorkResult{}, fmt.Errorf("unknown generator key: %s", it.Key)
		}
		count := it.Count
		if count <= 0 {
			count = 1
		}
		for k := 0; k < count; k++ {
			seed := fmt.Sprintf("%s:%d:%d:%s", cfg.Seed, itemIndex, k, it.Key)
			r := NewRand(seed)

			opts := cfg.Options
			if it.Options.Difficulty != 0 {
				opts.Difficulty = it.Options.Difficulty
			}

			stmt, sol, ok := gen.Generate(r, opts)
			if !ok {
				return WorkResult{}, fmt.Errorf("failed to generate problem for key=%s after 10000", it.Key)
			}
			meta := gen.Meta()
			result.Problems = append(result.Problems, Problem{Category: it.Key, Title: meta.Title, Statement: stmt, Solution: sol})
		}
	}
	return result, nil
}
