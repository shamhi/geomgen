package categories

import (
	"github.com/shamhi/geomgen/categories/lines"
	"github.com/shamhi/geomgen/categories/planes"
	"github.com/shamhi/geomgen/categories/triangles"
	"github.com/shamhi/geomgen/categories/vectors"
)

type VectorAngleGenerator = vectors.VectorAngleGenerator
type LineAngleGenerator = lines.LineAngleGenerator
type LinePerpPlaneGenerator = lines.LinePerpPlaneGenerator
type AngleLinePlaneGenerator = lines.AngleLinePlaneGenerator
type TriangleGenerator = triangles.TriangleGenerator
type PlaneThroughTwoPointsParallelVectorGenerator = planes.PlaneThroughTwoPointsParallelVectorGenerator
