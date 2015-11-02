package matrices

import (
    "fmt"
    "math/rand"
    "errors"
)

// Vector represends one-dimensional field
type Vector []float64

// Rows returns number of rows in vector - 1
func (v Vector) Rows() int {
    return 1
}

// Cols returns number of columns in vector
func (v Vector) Cols() int {
    return len(v)
}

// InitVector creates new vector with given length
func InitVector(length int) Vector {
    return Vector(make([]float64, length))
}

// InitVectorWithValues creates new vector with given values
func InitVectorWithValues(values []float64) Vector {
    return Vector(values)
}

// RandInitVector creates new vector with given length and random values
func RandInitVector(length int) Vector {
    v := InitVector(length)
    for i := range v {
        v[i] = rand.Float64()
    }
    return v
}

func (v Vector) checkIndex(index int) bool {
    return index >= 0 && index < len(v)
}

// At returns item that is in vector at given index
func (v Vector) At(i int) (float64, error) {
    if !v.checkIndex(i) {
        return 0, errors.New("matrices: cannot get value outside of vector")
    }
    return v[i], nil
}

// Set sets item in vector to given value
func (v Vector) Set(i int, value float64) error {
    if !v.checkIndex(i) {
        return errors.New("matrices: cannot set value outside of vector")
    }
    v[i] = value
    return nil
}

func (v Vector) operate(w Vector, operation func(float64, float64) float64) (Vector, error) {
    var result Vector
    if len(v) != len(w) {
        return result, errors.New("matrices: operating on two vectors with different lengths")
    }
    result = InitVector(len(v))
    for i := range result {
        result[i] = operation(v[i], w[i])
    }
    return result, nil
}

// Add adds two vectors
func (v Vector) Add(w Vector) (Vector, error) {
    return v.operate(w, func (x, y float64) float64 { return x + y; })
}

// Sub subtracts two vectors
func (v Vector) Sub(w Vector) (Vector, error) {
    return v.operate(w, func (x, y float64) float64 { return x - y; })
}

// Dot returns dot product of two vectors
func (v Vector) Dot(w Vector) (float64, error) {
    if len(v) != len(w) {
        return 0, errors.New("matrices: cannot do dot product of vectors with different sizes")
    }
    sum := float64(0)
    for i := range v {
        sum += v[i] * w[i]
    }
    return sum, nil
}

// Apply applies function to each element of vector
func (v Vector) Apply(operation func(float64) float64) Vector {
    result := InitVector(len(v))
    for i, val := range v {
        result[i] = operation(val)
    }
    return result
}

// Max returns biggest value in vector
func (v Vector) Max() (float64, error) {
    if len(v) == 0 {
        return 0, errors.New("matrices: can't get max value in empty vector")
    }
    maxval := v[0]
    for _, val := range v {
        if val > maxval {
            maxval = val
        }
    }
    return maxval, nil
}

// Min returns smallest value in vector
func (v Vector) Min() (float64, error) {
    if len(v) == 0 {
        return 0, errors.New("matrices: can't get min value in empty vector")
    }
    minval := v[0]
    for _, val := range v {
        if val < minval {
            minval = val
        }
    }
    return minval, nil
}

func (v Vector) String() (result string) {
    if len(v) == 0 {
        return "empty Vector"
    }
    maxval, _ := v.Max()
    flen := numberlen(maxval)
    floatfmt := fmt.Sprintf("%%%d.2f", flen + 6)
    for _, val := range v {
        result += fmt.Sprintf(floatfmt, val)
    }
    result = "[ " + result + " ]"
    return
}
