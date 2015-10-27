package matrices

// Field represents generic type for describing matrices or vectors
type Field interface {
    Shape() []int
}

func numberlen(f float64) (res int) {
    for f >= 10 {
        f /= 10
        res++
    }
    return
}
