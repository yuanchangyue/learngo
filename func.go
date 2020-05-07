package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported arg : %s", op)
	}
}

// 复合函数 简化 eval (函数编程)
func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	optName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Call funcation %s with args (%d,%d)", optName, a, b)
	return op(a, b)
}

//可变参数列表
func sum(values ...int) int {
	var sum int
	for i := range values {
		sum += values[i]
	}
	return sum
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(eval(1, 3, "2"))
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(sum(1, 2, 3, 4))
}
