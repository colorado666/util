package amath

func IntToUint(v int) uint {
    if v < 0 {
        return 0
    }
    return uint(v)
}

func Int64ToUint64(v int64) uint64 {
    if v < 0 {
        return 0
    }
    return uint64(v)
}

func IntAbs(v int) int {
    if v < 0 {
        return -v
    }
    return v
}

func Int64Abs(v int64) int64 {
    if v < 0 {
        return -v
    }
    return v
}
