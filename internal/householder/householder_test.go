package householder_test

import (
	"testing"

	"github.com/LLIEPJIOK/matrix/internal/householder"
	"github.com/LLIEPJIOK/matrix/internal/matrix"
)

func BenchmarkQRSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		mtr, rhs := matrix.GenerateMatrixAndRHS()
		b.StartTimer()

		_, _, _, _ = householder.QRSolve(mtr, rhs)
	}
}
