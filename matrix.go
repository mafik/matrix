package matrix

import (
	"math"

	"mrogalski.eu/go/vec"
)

// Matrix represents a 3x3 matrix with an implicit 0,0,1 row at the bottom.
//
// Only the first two rows are actually stored.
//
// Example:
//
//  [ #0 #2 #4 ]
//  [ #1 #3 #5 ]
//  [  0  0  1 ]
type Matrix [6]float64

// Identity is an identity matrix.
var Identity = Matrix{1, 0, 0, 1, 0, 0}

// Translation is a transform matrix that adds the vector `t`.
func Translation(t vec.Vec) Matrix {
	return Matrix{1, 0, 0, 1, t.X, t.Y}
}

// Scale is a transform matrix that multiplies by the factor `s`.
func Scale(s float64) Matrix {
	return Matrix{s, 0, 0, s, 0, 0}
}

// Rotation is a transform matrix that rotates by `r` radians.
func Rotation(r float64) Matrix {
	s, c := math.Sin(r), math.Cos(r)
	return Matrix{c, s, -s, c, 0, 0}
}

// Determinant of this matrix.
func (m Matrix) Determinant() float64 {
	return m[0]*m[3] - m[1]*m[2]
}

// Inverse of the matrix `m`.
func (m Matrix) Inverse() Matrix {
	d := m.Determinant()
	return Matrix{
		m[3] / d, -m[1] / d,
		-m[2] / d, m[0] / d,
		(m[2]*m[5] - m[3]*m[4]) / d,
		(m[1]*m[4] - m[0]*m[5]) / d,
	}
}

// Mul multiplies the matrix `m` by the matrix `n`.
func (m Matrix) Mul(n Matrix) Matrix {
	return Matrix{
		m[0]*n[0] + m[1]*n[2],
		m[1]*n[3] + m[0]*n[1],
		m[2]*n[0] + m[3]*n[2],
		m[3]*n[3] + m[2]*n[1],
		m[4]*n[0] + m[5]*n[2] + n[4],
		m[5]*n[3] + m[4]*n[1] + n[5],
	}
}

// Transform vector `a` by applying matrix `m`.
func (m Matrix) Transform(a vec.Vec) vec.Vec {
	return vec.New(
		a.X*m[0]+a.Y*m[2]+m[4],
		a.X*m[1]+a.Y*m[3]+m[5])
}
