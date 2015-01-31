package twod

// MatSolve returns the scalar x such that (ax - b) is minimized for the given norm
func MatSolve(ad, bd Denser, norm NormType) float64 {
	switch norm {
	default:
		panic("unknown norm type")
	case Frobenius, Frobenius2:
		// Treats a and b as a [4]float64 and minimizes the two norm
		// (a_11  - b_11 x)^2 + ...
		// Take derivative and solve, get x = (\sum_{ij} a_ij b_ij) / (\sum_{ij} aij^2)
		a := ad.Dense2d()
		b := bd.Dense2d()
		aijsq := a[0][0]*a[0][0] +
			a[1][0]*a[1][0] +
			a[0][1]*a[0][1] +
			a[1][1]*a[1][1]
		aijbij := a[0][0]*b[0][0] +
			a[1][0]*b[1][0] +
			a[0][1]*b[0][1] +
			a[1][1]*b[1][1]
		return aijbij / aijsq
	}
}
