package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 转化二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 按照一行一行的读文件
func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(6),
	)
	printFile("abc. txt")
}
