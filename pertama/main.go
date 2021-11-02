package main

import (
	"fmt"
	"pertama/calculation"
	"pertama/quiz"
)

func main() {
	fmt.Println("Halo, Belajar Golang")
	result := calculation.Add(1, 2)
	resultPerkalian := calculation.Kali(1, 2)
	resultPembagian := quiz.Bagi(100, 2)

	fmt.Println(result)
	fmt.Println(resultPerkalian)
	fmt.Println(resultPembagian)
}
