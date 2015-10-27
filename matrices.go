package matrices

import (
    "fmt"
    "math/rand"
    "errors"
    "strings"
)

// Matrix represents two-dimensional field
type Matrix struct {
    Rows, Cols int
    values []float64
}

// Shape returns dimensions of matrix in slice
func (m *Matrix) Shape() []int {
    return []int{m.Rows, m.Cols}
}

// InitMatrix initializes Matrix structure to have required number of rows and columns
func InitMatrix(rows, cols int) *Matrix {
    m := Matrix{Rows: rows, Cols: cols}
    m.values = make([]float64, rows*cols)
    return &m
}

// RandInitMatrix initializes Matrix structure and fills it with random numbers from random generator
func RandInitMatrix(rows, cols int, r *rand.Rand) *Matrix {
    m := InitMatrix(rows, cols)
    for i := 0; i < rows * cols; i++ {
        m.values[i] = r.Float64()
    }
    return m
}

// InitMatrixWithValues initializes Matrix with given dimensions and values
func InitMatrixWithValues(rows, cols int, values []float64) (*Matrix, error) {
    if len(values) != rows*cols {
        return nil, errors.New("matrices: bad dimensions of matrix")
    }
    m := Matrix{Rows: rows, Cols: cols, values: values}
    return &m, nil
}

func (m *Matrix) checkRowCol(row, col int) bool {
    return row < m.Rows && col < m.Cols && row >= 0 && col >= 0
}

func (m *Matrix) at(row, col int) float64 {
    return m.values[row * m.Cols + col]
}

func (m *Matrix) set(row, col int, value float64) {
    m.values[row * m.Cols + col] = value
}

// At returns item that is in matrix at given coordinates
func (m *Matrix) At(row, col int) (float64, error) {
    if !m.checkRowCol(row, col) {
        return 0, errors.New("matrices: cannot get value outside of matrix")
    }
    return m.at(row, col), nil
}

// Set sets item in matrix to given value
func (m *Matrix) Set(row, col int, value float64) error {
    if !m.checkRowCol(row, col) {
        return errors.New("matrices: cannot set value outside of matrix")
    }
    m.set(row, col, value)
    return nil
}

func (m *Matrix) operate(n *Matrix, operation func(float64, float64) float64) (*Matrix, error) {
    if m.Rows != n.Rows || m.Cols != n.Cols {
        return nil, errors.New("matrices: operating on two matrices with different dimensions")
    }
    result := InitMatrix(m.Rows, m.Cols)
    for i := range m.values {
        result.values[i] = operation(m.values[i], n.values[i])
    }
    return result, nil
}

// Add adds two matrices
func (m *Matrix) Add(n *Matrix) (*Matrix, error) {
    return m.operate(n, func (x, y float64) float64 { return x + y; })
}

// Sub subtracts two matrices
func (m *Matrix) Sub(n *Matrix) (*Matrix, error) {
    return m.operate(n, func (x, y float64) float64 { return x - y; })
}

// Apply applies function to each element of Matrix
func (m *Matrix) Apply(operation func(float64) float64) *Matrix {
    result := InitMatrix(m.Rows, m.Cols)
    for i, val := range m.values {
        result.values[i] = operation(val)
    }
    return result
}

// MatrixMult multiplies two matrices
func (m *Matrix) MatrixMult(n *Matrix) (*Matrix, error) {
    if m.Cols != n.Rows {
        return nil, errors.New("matrices: for matrix multiplication, first matrix cols == second matrix rows")
    }
    result := InitMatrix(m.Rows, n.Cols)
    for i := 0; i < result.Rows; i++ {
        for j := 0; j < result.Cols; j++ {
            sum := 0.0
            for counter := 0; counter < m.Cols; counter++ {
                sum += m.at(i, counter) * n.at(counter, j)
            }
            result.set(i, j, sum)
        }
    }
    return result, nil
}

// Transpose creates transposed matrix of original matrix
func (m *Matrix) Transpose() *Matrix {
    result := InitMatrix(m.Cols, m.Rows)
    for i := 0; i < m.Rows; i++ {
        for j := 0; j < m.Cols; j++ {
            val := m.at(i, j)
            result.set(j, i, val)
        }
    }
    return result
}

// Max returns biggest value in matrix
func (m *Matrix) Max() (float64, error) {
    if m.Rows == 0 || m.Cols == 0 {
        return 0, errors.New("matrices: can't return max value in empty matrix")
    }
    maxval := m.values[0]
    for _, val := range m.values {
        if val > maxval {
            maxval = val
        }
    }
    return maxval, nil
}

// Min returns smallest value in matrix
func (m *Matrix) Min() (float64, error) {
    if m.Rows == 0 || m.Cols == 0 {
        return 0, errors.New("matrices: can't return min value in empty matrix")
    }
    minval := m.values[0]
    for _, val := range m.values {
        if val < minval {
            minval = val
        }
    }
    return minval, nil
}

// String converts matrix to string
func (m *Matrix) String() (result string) {
    if m.Rows == 0 || m.Cols == 0 {
        return "empty Matrix"
    }
    maxval, _ := m.Max()
    flen := numberlen(maxval)
    floatfmt := fmt.Sprintf("%%%d.2f", flen + 6)
    rows := make([]string, m.Rows)

    for i := 0; i < m.Rows; i++ {
        row := ""
        for j := 0; j < m.Cols; j++ {
            val := m.at(i, j)
            row += fmt.Sprintf(floatfmt, val)
        }
        rows[i] = "| " + row + " |"
    }
    result = strings.Join(rows, "\n")
    return
}
