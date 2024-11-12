package matrix

import (
	"fmt"
	"math"
)

func validateMatrix(matrix [][]float64) error {
	if len(matrix) == 0 {
		return NewErrMatrix("matrix is empty")
	}

	n, m := len(matrix), len(matrix[0])

	for i := range n {
		if len(matrix[i]) != m {
			return NewErrMatrix("matrix is not rectangular")
		}
	}

	if m == 0 {
		return NewErrMatrix("matrix is empty")
	}

	return nil
}

func Add(first, second [][]float64) error {
	if err := validateMatrix(first); err != nil {
		return fmt.Errorf("invalid first matrix: %w", err)
	}

	if err := validateMatrix(second); err != nil {
		return fmt.Errorf("invalid second matrix: %w", err)
	}

	m, n := len(first), len(first[0])

	if m != len(second) || n != len(second[0]) {
		return NewErrMatrix("matrix must have equal size")
	}

	for i := range m {
		for j := range n {
			first[i][j] += second[i][j]
		}
	}

	return nil
}

func Sub(first, second [][]float64) error {
	if err := validateMatrix(first); err != nil {
		return fmt.Errorf("invalid first matrix: %w", err)
	}

	if err := validateMatrix(second); err != nil {
		return fmt.Errorf("invalid second matrix: %w", err)
	}

	m, n := len(first), len(first[0])

	if m != len(second) || n != len(second[0]) {
		return NewErrMatrix("matrix must have equal size")
	}

	for i := range m {
		for j := range n {
			first[i][j] -= second[i][j]
		}
	}

	return nil
}

func MultiplyByNumber(matrix [][]float64, number float64) {
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			matrix[i][j] *= number
		}
	}
}

func Multiply(first, second [][]float64) ([][]float64, error) {
	if err := validateMatrix(first); err != nil {
		return nil, fmt.Errorf("invalid first matrix: %w", err)
	}

	if err := validateMatrix(second); err != nil {
		return nil, fmt.Errorf("invalid second matrix: %w", err)
	}

	if len(first[0]) != len(second) {
		return nil, NewErrMatrix("these matrices cannot be multiplied")
	}

	mult := make([][]float64, len(first))

	for i := range len(first) {
		mult[i] = make([]float64, len(second[0]))
		for j := range len(second[0]) {
			for k := range len(second) {
				mult[i][j] += first[i][k] * second[k][j]
			}
		}
	}

	return mult, nil
}

func Transpose(matrix [][]float64) ([][]float64, error) {
	if err := validateMatrix(matrix); err != nil {
		return nil, fmt.Errorf("invalid first matrix: %w", err)
	}

	m, n := len(matrix), len(matrix[0])
	transposeMatrix := make([][]float64, n)

	for j := range n {
		transposeMatrix[j] = make([]float64, m)

		for i := range m {
			transposeMatrix[j][i] = matrix[i][j]
		}
	}

	return transposeMatrix, nil
}

func Identity(n int) ([][]float64, error) {
	if n <= 0 {
		return nil, NewErrMatrix("matrix size must be positive")
	}

	matrix := make([][]float64, n)
	for i := range n {
		matrix[i] = make([]float64, n)
		matrix[i][i] = 1
	}

	return matrix, nil
}

// спектральная норма для столбца s_i
func SpectralNormForColumn(matrix [][]float64) (float64, error) {
	if len(matrix) == 0 {
		return 0.0, NewErrMatrix("matrix is empty")
	}

	n, m := len(matrix), len(matrix[0])

	if m != 1 {
		return 0.0, NewErrMatrix("matrix should be a column")
	}

	for i := range n {
		if len(matrix[i]) != m {
			return 0.0, NewErrMatrix("matrix is not rectangular")
		}
	}

	norm := 0.0

	for i := range n {
		norm += math.Pow(matrix[i][0], 2)
	}

	norm = math.Sqrt(norm)

	return norm, nil
}

func CopyMatrix(matrix []float64) []float64 {
	matrixCopy := make([]float64, len(matrix))

	for i := range len(matrix) {
		matrixCopy[i] = matrix[i]
	}

	return matrixCopy
}

func Copy2DMatrix(matrix [][]float64) [][]float64 {
	matrixCopy := make([][]float64, len(matrix))

	for i := range len(matrix) {
		matrixCopy[i] = CopyMatrix(matrix[i])
	}

	return matrixCopy
}

func PrintMatrix(matrix []float64) {
	for i := range len(matrix) {
		fmt.Printf("%10.5f ", matrix[i])
	}

	fmt.Println()
}

func Print2DMatrix(matrix [][]float64) {
	for i := range len(matrix) {
		PrintMatrix(matrix[i])
	}
}

func PrintVector(vector [][]float64) {
	for i := range len(vector) {
		fmt.Printf("%10.5f ", vector[i][0])
	}

	fmt.Println()
}

// generating matrix and right-hand side vector
const MatrixSize = 15

func GenerateMatrixAndRHS() ([][]float64, [][]float64) {
	mtr := make([][]float64, MatrixSize)

	for i := range MatrixSize {
		mtr[i] = make([]float64, MatrixSize)

		for j := range MatrixSize {
			if i == j {
				mtr[i][j] = float64(5 * (i + 1))
			} else {
				mtr[i][j] = -(float64(i+1) + math.Sqrt(float64(j+1)))
			}
		}
	}

	rhs := make([][]float64, MatrixSize)

	for i := range MatrixSize {
		rhs[i] = make([]float64, 1)
		rhs[i][0] = 3.0 * math.Sqrt(float64(i+1))
	}

	return mtr, rhs
}

func PrintMatrixAndRHS(mtr, rhs [][]float64) {
	fmt.Println("Initial matrix A:")
	Print2DMatrix(mtr)
	fmt.Println()

	fmt.Println("Initial right-hand side vector b:")
	PrintVector(rhs)
	fmt.Println()
}

// checking correctness of data
func Validate(matrix [][]float64, rhs [][]float64) error {
	if len(matrix) == 0 {
		return NewErrMatrix("matrix is empty")
	}

	n := len(matrix)

	if len(rhs) != n {
		return NewErrRHS("right-hand side vector must have the same size as matrix")
	}

	for i := range n {
		if len(matrix[i]) != n {
			return NewErrMatrix("matrix is not square")
		}

		if len(rhs[i]) != 1 {
			return NewErrRHS("right-hand side vector must be a column")
		}
	}

	return nil
}

// calculating the cubic norm for a vector difference
func CalculateXDiff(matrix, rhs [][]float64, xs []float64) (float64, error) {
	if len(matrix) == 0 {
		return 0.0, NewErrMatrix("matrix is empty")
	}

	n := len(matrix)

	for i := range n {
		if len(matrix[i]) != n {
			return 0.0, NewErrMatrix("matrix is not square")
		}
	}

	if len(rhs) != n {
		return 0.0, NewErrRHS("right-hand side vector must have the same size as matrix")
	}

	mx := 0.0

	for i := range len(matrix) {
		rowSum := 0.0

		for j := range len(matrix[i]) {
			rowSum += matrix[i][j] * xs[j]
		}

		mx = max(mx, math.Abs(rowSum-rhs[i][0]))
	}

	return mx, nil
}

// calculating the cubic norm for a matrix difference
func CalculateMatrixDiff(first [][]float64, second [][]float64) (float64, error) {
	err := Sub(first, second)
	if err != nil {
		return 0.0, fmt.Errorf("sub first and second: %w", err)
	}

	mx := 0.0
	for i := range len(first) {
		cur := 0.0

		for j := range len(first[i]) {
			cur += math.Abs(first[i][j])
		}

		mx = max(mx, cur)
	}

	return mx, nil
}
