package matrix3

// Matrix3 represents a 3x3 matrix.
type Mtx3 [3][3]float64

// NewIdentity creates and returns an identity matrix.
func NewIdentity() Mtx3 {
	return Mtx3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}
