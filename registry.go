package geomgen

import (
	"math/rand"
	"sort"

	"github.com/shamhi/geomgen/categories/lines"
	"github.com/shamhi/geomgen/categories/planes"
	"github.com/shamhi/geomgen/categories/triangles"
	"github.com/shamhi/geomgen/categories/vectors"
)

// UnifiedGenerator is a non-generic facade over concrete generators
type UnifiedGenerator interface {
	Key() string
	Generate(r *rand.Rand, opts Options) (statement string, solution string, ok bool)
	Meta() GeneratorMeta
}

type GeneratorMeta struct {
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
	MinDiff int      `json:"min_diff"`
	MaxDiff int      `json:"max_diff"`
}

// Generic adapter wrapping existing ExpressionGenerator[T]
type GeneratorAdapter[T any] struct {
	Impl ExpressionGenerator[T]
	M    GeneratorMeta
}

func (a GeneratorAdapter[T]) Key() string { return a.Impl.Category() }

func (a GeneratorAdapter[T]) Meta() GeneratorMeta {
	if a.M.Title == "" {
		return GeneratorMeta{
			Title:   a.Impl.Title(),
			Tags:    a.M.Tags,
			MinDiff: a.M.MinDiff,
			MaxDiff: a.M.MaxDiff,
		}
	}
	return a.M
}

func (a GeneratorAdapter[T]) Generate(r *rand.Rand, opts Options) (stmt string, sol string, ok bool) {
	maxAttempts := opts.MaxAttempts
	if maxAttempts <= 0 {
		maxAttempts = 10000
	}
	for i := 0; i < maxAttempts; i++ {
		expr := a.Impl.Generate(r)
		if !a.Impl.Validate(expr) {
			continue
		}
		stmt = a.Impl.Statement(expr)
		var err error
		sol, err = a.Impl.Solve(expr)
		if err != nil {
			continue
		}
		ok = true
		return
	}
	return "", "", false
}

// Registry holds available generators
type Registry struct {
	byKey map[string]UnifiedGenerator
	order []string
}

func NewRegistry() *Registry { return &Registry{byKey: map[string]UnifiedGenerator{}} }

func (r *Registry) Register(g UnifiedGenerator) {
	key := g.Key()
	if _, exists := r.byKey[key]; !exists {
		r.order = append(r.order, key)
		sort.Strings(r.order)
	}
	r.byKey[key] = g
}

func (r *Registry) Get(key string) (UnifiedGenerator, bool) { g, ok := r.byKey[key]; return g, ok }

func (r *Registry) Keys() []string { return append([]string(nil), r.order...) }

// Default registry pre-populated with available generators
var DefaultRegistry = func() *Registry {
	reg := NewRegistry()

	reg.Register(GeneratorAdapter[vectors.VectorPair]{
		Impl: &vectors.VectorAngleGenerator{},
		M:    GeneratorMeta{Title: "Угол между векторами", Tags: []string{"vectors", "angle"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[lines.LinePair]{
		Impl: &lines.LineAngleGenerator{},
		M:    GeneratorMeta{Title: "Угол между прямыми", Tags: []string{"lines", "angle"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[lines.PointPlane]{
		Impl: &lines.LinePerpPlaneGenerator{},
		M:    GeneratorMeta{Title: "Прямая ⟂ плоскости через точку", Tags: []string{"lines", "planes"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[lines.LineAndPlane]{
		Impl: &lines.AngleLinePlaneGenerator{},
		M:    GeneratorMeta{Title: "Угол между прямой и плоскостью", Tags: []string{"lines", "planes", "angle"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[triangles.TrianglePoints]{
		Impl: &triangles.TriangleGenerator{},
		M:    GeneratorMeta{Title: "Средняя линия и медиана треугольника", Tags: []string{"triangles", "lines"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[planes.TwoPointsVector]{
		Impl: &planes.PlaneThroughTwoPointsParallelVectorGenerator{},
		M:    GeneratorMeta{Title: "Плоскость через две точки ∥ вектору", Tags: []string{"planes"}, MinDiff: 1, MaxDiff: 3},
	})

	return reg
}()
