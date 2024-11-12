package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/LLIEPJIOK/matrix/internal/householder"
	"github.com/LLIEPJIOK/matrix/internal/matrix"
)

func main() {
	mtr, rhs := matrix.GenerateMatrixAndRHS()
	matrix.PrintMatrixAndRHS(mtr, rhs)

	q, r, xs, err := householder.QRSolve(matrix.Copy2DMatrix(mtr), matrix.Copy2DMatrix(rhs))
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

	fmt.Println("Matrix Q:")
	matrix.Print2DMatrix(q)
	fmt.Println()

	fmt.Println("Matrix R:")
	matrix.Print2DMatrix(r)
	fmt.Println()

	qr, err := matrix.Multiply(q, r)
	if err != nil {
		slog.Error(fmt.Sprintf("multiply Q and R: %s", err))
		os.Exit(1)
	}

	diff, err = matrix.CalculateMatrixDiff(mtr, qr)
	if err != nil {
		slog.Error(fmt.Sprintf("calculate difference for QR: %s", err))
		os.Exit(1)
	}

	fmt.Printf("||Е - QR||: %e\n", diff)
}
