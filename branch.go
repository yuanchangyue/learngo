package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score : %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "B"
	case score < 90:
		g = "A"
	}
	return g
}

func readFile() {
	const fileName = "abc.txt"
	if content, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s", content)
	}
}

func main() {
	readFile()
	fmt.Println(grade(12))
}
