package gauss_test

import (
	"testing"

	"github.com/LLIEPJIOK/matrix/internal/gauss"
	"github.com/LLIEPJIOK/matrix/internal/matrix"
)

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		mtr, rhs := matrix.GenerateMatrixAndRHS()
		b.StartTimer()

		_, _ = gauss.Solve(mtr, rhs)
	}
}
