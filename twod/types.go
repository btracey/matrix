package twod

import "math"

var panicRowBound = "twod: matrix row access out of bounds"
var panicColBound = "twod: matrix column access out of bounds"

var panicUnknownNorm = "two: unknown norm type"

func matCheckBounds(i, j int) {
	if i < 0 || i > 1 {
		panic(panicRowBound)
	}
	if j < 0 || j > 1 {
		panic(panicColBound)
	}
}

// Dense represents a two-dimensional dense matrix
type Dense [2][2]float64

func (d Dense) Dims() (float64, float64) {
	return 2, 2
}

// At gets an element from the 2-D matrix
func (d Dense) At(i, j int) float64 {
	matCheckBounds(i, j)
	return d[i][j]
}

func (d *Dense) Set(i, j int, v float64) {
	matCheckBounds(i, j)
	d[i][j] = v
}

func (d *Dense) Norm(o NormType) float64 {
	switch o {
	case Frobenius:
		sum := d[0][0]*d[0][0] + d[0][1]*d[0][1] + d[1][0]*d[1][0] + d[1][1]*d[1][1]
		return math.Sqrt(sum)
	default:
		panic(panicUnknownNorm)
	}
}

func (d Dense) Dense2d() Dense {
	return d
}

// Split divides the dense into its symmetric and anti-symmetric components
func (d Dense) Split() (Symmetric, SkewSymmetric) {
	offDiag := (d[1][0] + d[0][1]) / 2
	s := Symmetric{d[0][0], offDiag, d[1][1]}

	return s, SkewSymmetric((d[0][1] - d[1][0]) / 2)
}

//  represents a two-dimensional tensor
type Symmetric [3]float64

// At gets an element from the 2-D matrix
func (s Symmetric) At(i, j int) float64 {
	matCheckBounds(i, j)
	switch {
	case i == 0 && j == 0:
		return s[0]
	case i == 1 && j == 0, j == 1 && i == 0:
		return s[1]
	case j == 1 && i == 1:
		return s[2]
	default:
		panic("unreachable")
	}
}

func (s *Symmetric) Set(i, j int, v float64) {
	matCheckBounds(i, j)
	switch {
	case i == 0 && j == 0:
		s[0] = v
	case i == 1 && j == 0, j == 1 && i == 0:
		s[1] = v
	case j == 1 && i == 1:
		s[2] = v
	default:
		panic("unreachable")
	}
	return
}

func (s Symmetric) Dense2d() Dense {
	return Dense{
		{s[0], s[1]},
		{s[1], s[2]},
	}
}

func (s Symmetric) Norm(o NormType) float64 {
	switch o {
	case Frobenius:
		sum := s[0]*s[0] + 2*s[1]*s[1] + s[2]*s[2]
		return math.Sqrt(sum)
	default:
		panic(panicUnknownNorm)
	}
}

type Vector [2]float64

func (v Vector) Vector() Vector {
	return v
}

type ColumnVector Vector

type RowVector Vector

// SkewSymmetric represents a two dimensional skew-symmetric matrix. The underlying
// float64 is the value of (0,1). The value of (1,0) is the negative
type SkewSymmetric float64

// At gets an element from the 2-D matrix
func (s SkewSymmetric) At(i, j int) float64 {
	matCheckBounds(i, j)
	switch {
	case i == 0 && j == 0:
		return 0
	case i == 0 && j == 1:
		return float64(s)
	case i == 1 && j == 0:
		return float64(-s)
	case j == 1 && i == 1:
		return 0
	default:
		panic("unreachable")
	}
}

func (s SkewSymmetric) Dense2d() Dense {
	return Dense{
		{0, float64(s)},
		{-float64(s), 0},
	}
}
