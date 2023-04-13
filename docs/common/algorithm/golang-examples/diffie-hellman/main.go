package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Генерируем случайные простые числа p и g
	p, _ := rand.Prime(rand.Reader, 128)
	g := big.NewInt(2)

	// Генерируем случайное число a для первого участника
	a, _ := rand.Int(rand.Reader, p)

	// Вычисляем A = g^a mod p
	A := new(big.Int).Exp(g, a, p)

	// Генерируем случайное число b для второго участника
	b, _ := rand.Int(rand.Reader, p)

	// Вычисляем B = g^b mod p
	B := new(big.Int).Exp(g, b, p)

	// Обменяемся значениями A и B

	// Вычисляем общий секретный ключ K = B^a mod p
	K1 := new(big.Int).Exp(B, a, p)

	// Вычисляем общий секретный ключ K = A^b mod p
	K2 := new(big.Int).Exp(A, b, p)

	// Проверяем, что оба вычисления дают одинаковый результат
	if K1.Cmp(K2) == 0 {
		fmt.Println("Секретный ключ успешно обменен!")
		fmt.Println("K =", K1)
	} else {
		fmt.Println("Ошибка: секретные ключи не совпадают!")
	}
}
