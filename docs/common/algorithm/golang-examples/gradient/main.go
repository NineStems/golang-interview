package main

import (
	"fmt"
	"math"
)

// Функция, которую мы хотим минимизировать
func f(x float64) float64 {
	return math.Pow(x, 2)
}

// Производная функции f
func df(x float64) float64 {
	return 2 * x
}

// Градиентный спуск
func gradientDescent(x float64, alpha float64, iterations int) float64 {
	for i := 0; i < iterations; i++ {
		x = x - alpha*df(x)
	}
	return x
}

func main() {
	// Начальное значение x
	x := 10.0

	// Скорость обучения (шаг градиентного спуска)
	alpha := 0.1

	// Количество итераций
	iterations := 100

	// Запускаем градиентный спуск
	result := gradientDescent(x, alpha, iterations)

	// Выводим результат
	fmt.Printf("Minimum point of f(x) is: %.2f\n", result)
}
