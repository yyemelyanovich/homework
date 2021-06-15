package fibonacci

import (
    "errors"
)

func Fibonacci(n interface{}) interface{} {
    switch n.(type) {
        case int:
            return fibonacciInt(n.(int))
        case uint:
            return fibonacciUint(n.(uint))
        case float64:
            return fibonacciFloat(n.(float64))
        default:
            return errors.New("input is not a number")
    }
}

func fibonacciUint(n uint) uint {
    first := uint(0)
    second := uint(1)

    for i := uint(0); i < n; i++ {
        first, second = second, first + second
    }
    return first
}

func fibonacciInt(n int) int {
    first := 0
    second := 1

    for i := 0; i < n; i++ {
        first, second = second, first + second
    }
    return first
}

func fibonacciFloat(n float64) float64 {
    first := float64(0)
    second := float64(1)

    for i := float64(0); i < n; i++ {
        first, second = second, first + second
    }
    return first
}
