package twod

// Vector can represent itself as a 2-D vector
type Vectorer interface {
	Vector2d() Vector
}

// Symmetric can represent itself as a 2-D symmetric tensor
type Symmetricer interface {
	Symmetric2d() Symmetric
}

type Denser interface {
	Dense2d() Dense
}

type NormType int

const (
	Frobenius = iota
)
