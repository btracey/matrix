package twod

// Vector can represent itself as a 2-D vector
type Vectorer interface {
	Vector() Vector
}

// Symmetric can represent itself as a 2-D symmetric tensor
type Symmetricer interface {
	Symmetric() Symmetric
}

type Denser interface {
	Dense() Dense
}

type NormType int

const (
	Frobenius = iota
)
