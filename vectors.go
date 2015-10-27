package matrices

import (
    "fmt"
    "math/rand"
    "errors"
)

// Vector represends one-dimensional field
type Vector struct {
    Length int
    values []float64
}

// Shape returns returns length of vector
func (v *Vector) Shape() []int {
    return []int{v.Length}
}

// InitVector creates new vector with given length
func InitVector(length int) *Vector {
    v := Vector{Length: length}
    v.values = make([]float64, length)
    return &v
}

// InitVectorWithValues creates new vector with given values
func InitVectorWithValues(values []float64) *Vector {
    v := Vector{Length: len(values), values: values}
    return &v
}

// RandInitVector creates new vector with given length and random values
func RandInitVector(length int, r *rand.Rand) *Vector {
    v := InitVector(length)
    for i := range v.values {
        v.values[i] = r.Float64()
    }
    return v
}

func (v *Vector) checkIndex(index int) bool {
    return index >= 0 && index < v.Length
}

// At returns item that is in vector at given index
func (v *Vector) At(i int) (float64, error) {
    if !v.checkIndex(i) {
        return 0, errors.New("matrices: cannot get value outside of vector")
    }
    return v.values[i], nil
}

// Set sets item in vector to given value
func (v *Vector) Set(i int, value float64) error {
    if !v.checkIndex(i) {
        return errors.New("matrices: cannot set value outside of vector")
    }
    v.values[i] = value
    return nil
}

func (v *Vector) operate(w *Vector, operation func(float64, float64) float64) (*Vector, error) {
    if v.Length != w.Length {
        return nil, errors.New("matrices: operating on two vectors with different lengths")
    }
    result := InitVector(v.Length)
    for i := range result.values {
        result.values[i] = operation(v.values[i], w.values[i])
    }
    return result, nil
}

// Add adds two vectors
func (v *Vector) Add(w *Vector) (*Vector, error) {
    return v.operate(w, func (x, y float64) float64 { return x + y; })
}

// Sub subtracts two vectors
func (v *Vector) Sub(w *Vector) (*Vector, error) {
    return v.operate(w, func (x, y float64) float64 { return x - y; })
}

// Dot returns dot product of two vectors
func (v *Vector) Dot(w *Vector) (float64, error) {
    if v.Length != w.Length {
        return 0, errors.New("matrices: cannot do dot product of vectors with different sizes")
    }
    sum := float64(0)
    for i := range v.values {
        sum += v.values[i] * w.values[i]
    }
    return sum, nil
}

// Apply applies function to each element of vector
func (v *Vector) Apply(operation func(float64) float64) *Vector {
    result := InitVector(v.Length)
    for i, val := range v.values {
        result.values[i] = operation(val)
    }
    return result
}

// Max returns biggest value in vector
func (v *Vector) Max() (float64, error) {
    if v.Length == 0 {
        return 0, errors.New("matrices: can't get max value in empty vector")
    }
    maxval := v.values[0]
    for _, val := range v.values {
        if val > maxval {
            maxval = val
        }
    }
    return maxval, nil
}

// Min returns smallest value in vector
func (v *Vector) Min() (float64, error) {
    if v.Length == 0 {
        return 0, errors.New("matrices: can't get min value in empty vector")
    }
    minval := v.values[0]
    for _, val := range v.values {
        if val < minval {
            minval = val
        }
    }
    return minval, nil
}

func (v *Vector) String() (result string) {
    if v.Length == 0 {
        return "empty Vector"
    }
    maxval, _ := v.Max()
    flen := numberlen(maxval)
    floatfmt := fmt.Sprintf("%%%d.2f", flen + 6)
    for _, val := range v.values {
        result += fmt.Sprintf(floatfmt, val)
    }
    result = "[ " + result + " ]"
    return
}
