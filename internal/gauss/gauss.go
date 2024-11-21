package gauss

import (
	"fmt"
	"math"

	"github.com/LLIEPJIOK/matrix-equations/internal/matrix"
)

func getMaxID(matrix [][]float64, row, colFrom int) int {
	maxID := colFrom
	for i := maxID + 1; i < len(matrix[row]); i++ {
		if math.Abs(matrix[row][i]) > math.Abs(matrix[row][maxID]) {
			maxID = i
		}
	}

	return maxID
}

func swapCols(matrix [][]float64, first, second int) {
	if first == second {
		return
	}

	for i := range matrix {
		matrix[i][first], matrix[i][second] = matrix[i][second], matrix[i][first]
	}
}

func removeElements(matrix [][]float64, rhs [][]float64, row, col int) {
	for i := row + 1; i < len(matrix); i++ {
		multiplier := matrix[i][col] / matrix[row][col]

		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] -= matrix[row][j] * multiplier
		}

		rhs[i][0] -= rhs[row][0] * multiplier
	}
}

func Solve(mtr [][]float64, rhs [][]float64) ([]float64, error) {
	// checking the correctness of the input data
	if err := matrix.Validate(mtr, rhs); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	n := len(mtr)

	// nums - vector for storing column indices
	nums := make([]int, n)

	for i := range n {
		nums[i] = i
	}

	// gauss part
	for i := range n {
		maxID := getMaxID(mtr, i, i)
		swapCols(mtr, i, maxID)
		nums[i], nums[maxID] = nums[maxID], nums[i]
		removeElements(mtr, rhs, i, i)
	}

	// getting answer
	xs := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		right := rhs[i][0]

		for j := n - 1; j > i; j-- {
			right -= mtr[i][j] * xs[nums[j]]
		}

		xs[nums[i]] = right / mtr[i][i]
	}

	return xs, nil
}
