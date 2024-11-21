package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/LLIEPJIOK/matrix-equations/internal/gauss"
	"github.com/LLIEPJIOK/matrix-equations/internal/matrix"
)

func main() {
	mtr, rhs := matrix.GenerateMatrixAndRHS()
	matrix.PrintMatrixAndRHS(mtr, rhs)

	xs, err := gauss.Solve(matrix.Copy2DMatrix(mtr), matrix.Copy2DMatrix(rhs))
	if err != nil {
		slog.Error(fmt.Sprintf("solve: %s", err))
		os.Exit(1)
	}

	fmt.Println("Vector x:")
	matrix.PrintMatrix(xs)
	fmt.Println()

	diff, err := matrix.CalculateXDiff(mtr, rhs, xs)
	if err != nil {
		slog.Error(fmt.Sprintf("calculate difference for x: %s", err))
		os.Exit(1)
	}

	fmt.Printf("Norm of the residual vector: %e\n", diff)
	fmt.Println()

	// calculating the inverse of a matrix using systems of n equations
	inverse := make([][]float64, matrix.MatrixSize)
	invRHS := make([][]float64, matrix.MatrixSize)

	for i := range matrix.MatrixSize {
		inverse[i] = make([]float64, matrix.MatrixSize)
		invRHS[i] = make([]float64, 1)
	}

	for i := range matrix.MatrixSize {
		invRHS[i][0] = 1

		xs, err := gauss.Solve(matrix.Copy2DMatrix(mtr), matrix.Copy2DMatrix(invRHS))
		if err != nil {
			slog.Error(fmt.Sprintf("solve: %s", err))
			os.Exit(1)
		}

		invRHS[i][0] = 0

		for j := range matrix.MatrixSize {
			inverse[j][i] = xs[j]
		}
	}

	fmt.Println("Matrix A^(-1):")
	matrix.Print2DMatrix(inverse)
	fmt.Println()

	mult, err := matrix.Multiply(mtr, inverse)
	if err != nil {
		slog.Error(fmt.Sprintf("multiply A and A^(-1): %s", err))
		os.Exit(1)
	}

	ident, err := matrix.Identity(matrix.MatrixSize)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to generate identity with size=%d: %s", matrix.MatrixSize, err))
		os.Exit(1)
	}

	diff, err = matrix.CalculateMatrixDiff(ident, mult)
	if err != nil {
		slog.Error(fmt.Sprintf("calculate difference for AA^(-1): %s", err))
		os.Exit(1)
	}

	fmt.Printf("||E - AA^(-1)||: %e\n", diff)
}
