package twod

// returns a * b
func Mul(a, b Denser) Dense {
	c := a.Dense2d()
	d := b.Dense2d()

	d00 := c[0][0]*d[0][0] + c[0][1]*d[1][0]
	d01 := c[0][0]*d[0][1] + c[0][1]*d[1][1]
	d10 := c[1][0]*d[0][0] + c[1][1]*d[1][0]
	d11 := c[1][0]*d[0][1] + c[1][1]*d[1][1]

	return Dense{
		{d00, d01},
		{d10, d11},
	}
}
