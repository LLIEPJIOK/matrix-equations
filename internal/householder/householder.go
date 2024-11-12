package householder

import (
	"fmt"

	"github.com/LLIEPJIOK/matrix/internal/matrix"
)

func sgn(n float64) int {
	if n >= 0 {
		return 1
	}

	return -1
}

func QRSolve(mtr [][]float64, rhs [][]float64) ([][]float64, [][]float64, []float64, error) {
	// checking the correctness of the input data
	if err := matrix.Validate(mtr, rhs); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid matrix: %w", err)
	}

	n := len(mtr)

	// matrix Q
	matrixQ, err := matrix.Identity(n)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to make identity matrix with size=%d: %w", n, err)
	}

	for i := range n {
		// vector s_i
		vs := make([][]float64, n-i)
		for j := i; j < n; j++ {
			vs[j-i] = make([]float64, 1)
			vs[j-i][0] = mtr[j][i]
		}

		// norm of vector s_i
		vsNorm, err := matrix.SpectralNormForColumn(vs)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to calculate spectral norm for column: %w", err)
		}

		// vector e_i
		ve := make([][]float64, n-i)
		for j := range len(ve) {
			ve[j] = make([]float64, 1)
		}

		ve[0][0] = 1

		// calculating w_i
		matrix.MultiplyByNumber(ve, vsNorm*float64(sgn(mtr[i][i])))

		err = matrix.Add(vs, ve)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to add matrix: %w", err)
		}

		vsNorm, err = matrix.SpectralNormForColumn(vs)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to calculate spectral norm for column: %w", err)
		}

		matrix.MultiplyByNumber(vs, 1.0/vsNorm)

		// calculating 2*w*w'
		transposeVs, err := matrix.Transpose(vs)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to transpose: %w", err)
		}

		vsMult, err := matrix.Multiply(vs, transposeVs)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to multiply vs and transpose vs: %w", err)
		}

		matrix.MultiplyByNumber(vsMult, 2.0)

		// calculating matrix H
		mtrH, err := matrix.Identity(n - i)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to make identity matrix with size=%d: %w", n-i, err)
		}

		err = matrix.Sub(mtrH, vsMult)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to subtract matrix: %w", err)
		}

		// recalculating matrix A
		partA := make([][]float64, n-i)
		
		for a := range n - i {
			partA[a] = make([]float64, n-i)

			for b := range n - i {
				partA[a][b] = mtr[i+a][i+b]
			}
		}

		partA, err = matrix.Multiply(mtrH, partA)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to multiply H and partA matrices: %w", err)
		}

		for a := range n - i {
			for b := range n - i {
				mtr[i+a][i+b] = partA[a][b]
			}
		}

		// calculating matrix Q_i
		mtrQ, err := matrix.Identity(n)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to make identity matrix with size=%d: %w", n, err)
		}

		for j := i; j < n; j++ {
			for k := i; k < n; k++ {
				mtrQ[j][k] = mtrH[j-i][k-i]
			}
		}

		// recalculating right-hand side vector b
		rhs, err = matrix.Multiply(mtrQ, rhs)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to multiply q and rhs matrices: %w", err)
		}

		// calculating Q_i^(-1) = Q_i'
		mtrQ, err = matrix.Transpose(mtrQ)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to transpose mtrQ: %w", err)
		}

		// recalculating Q
		matrixQ, err = matrix.Multiply(matrixQ, mtrQ)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("failed to multiply qs matrices: %w", err)
		}
	}

	// getting answer
	xs := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		right := rhs[i][0]

		for j := n - 1; j > i; j-- {
			right -= mtr[i][j] * xs[j]
		}

		xs[i] = right / mtr[i][i]
	}

	return matrixQ, mtr, xs, nil
}
