package kata

func PowersOfTwo(n int) []uint64 {
    result := make([]uint64, 0, n+1)

    for i := 0; i <= n; i++ {
        result = append(result, uint64(1)<<uint(i))
    }

    return result
}
