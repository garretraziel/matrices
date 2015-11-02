package matrices

func numberlen(f float64) (res int) {
    for f >= 10 {
        f /= 10
        res++
    }
    return
}
