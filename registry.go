package geomgen

import (
	"math/rand"
	"sort"

	"github.com/shamhi/geomgen/v2/categories/coordinates"
	"github.com/shamhi/geomgen/v2/categories/lines"
	"github.com/shamhi/geomgen/v2/categories/planes"
	"github.com/shamhi/geomgen/v2/categories/points"
	"github.com/shamhi/geomgen/v2/categories/triangles"
	"github.com/shamhi/geomgen/v2/categories/vectors"
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
	for i := 0; i < 10000; i++ {
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
		M:    GeneratorMeta{Title: "Угол между векторами", Tags: []string{"vectors", "geometry", "angle"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[lines.LinePair]{
		Impl: &lines.LineAngleGenerator{},
		M:    GeneratorMeta{Title: "Угол между прямыми", Tags: []string{"lines", "geometry", "angle"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[lines.PointPlane]{
		Impl: &lines.LinePerpPlaneGenerator{},
		M:    GeneratorMeta{Title: "Прямая, перпендикулярная плоскости через точку", Tags: []string{"lines", "planes", "perpendicular"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[lines.LineAndPlane]{
		Impl: &lines.AngleLinePlaneGenerator{},
		M:    GeneratorMeta{Title: "Угол между прямой и плоскостью", Tags: []string{"lines", "planes", "angle"}, MinDiff: 2, MaxDiff: 4},
	})
	reg.Register(GeneratorAdapter[triangles.TrianglePoints]{
		Impl: &triangles.TriangleMidlineMedianGenerator{},
		M:    GeneratorMeta{Title: "Средняя линия и медиана треугольника", Tags: []string{"triangles", "geometry", "lines"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[triangles.TriangleAreaPoints]{
		Impl: &triangles.TriangleAreaGenerator{},
		M:    GeneratorMeta{Title: "Площадь треугольника по координатам", Tags: []string{"triangles", "geometry", "area"}, MinDiff: 1, MaxDiff: 3},
	})
	reg.Register(GeneratorAdapter[planes.TwoPointsVector]{
		Impl: &planes.PlaneTwoPointsParallelVectorGenerator{},
		M:    GeneratorMeta{Title: "Плоскость через две точки, параллельная вектору", Tags: []string{"planes", "vectors", "geometry"}, MinDiff: 2, MaxDiff: 4},
	})
	reg.Register(GeneratorAdapter[coordinates.ChangeBasisData]{
		Impl: &coordinates.ChangeBasisGenerator{},
		M:    GeneratorMeta{Title: "Переход к новому базису координат", Tags: []string{"coordinates", "basis", "linear-algebra"}, MinDiff: 3, MaxDiff: 5},
	})
	reg.Register(GeneratorAdapter[points.MirrorPlaneData]{
		Impl: &points.MirrorPlaneGenerator{},
		M:    GeneratorMeta{Title: "Отражение точки относительно плоскости", Tags: []string{"points", "planes", "mirror", "geometry"}, MinDiff: 2, MaxDiff: 4},
	})

	return reg
}()
