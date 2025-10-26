package matrix3

import "math"	
import "fmt"

// Matrix3 represents a 3x3 matrix.
type Matrix3 [3][3]float64
// NewIdentity creates and returns an identity matrix.	
export func NewIdentity() Matrix3 {
	return Matrix3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}