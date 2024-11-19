package main

import "fmt"

func bitWiseOperations() {
	// Deslocamento à esquerda (<<)
	var num1 int = 1       // Representação binária: 0001
	shiftLeft := num1 << 2 // Desloca 2 bits para a esquerda: 0100 (equivale a 4)
	fmt.Printf("%d << 2 = %d (Binário: %08b)\n", num1, shiftLeft, shiftLeft)

	// Deslocamento à direita (>>)
	var num2 int = 8        // Representação binária: 1000
	shiftRight := num2 >> 2 // Desloca 2 bits para a direita: 0010 (equivale a 2)
	fmt.Printf("%d >> 2 = %d (Binário: %08b)\n", num2, shiftRight, shiftRight)

	// Exemplo com combinação de operações
	var num3 int = 15                // Representação binária: 1111
	leftAndRight := (num3 << 1) >> 3 // Desloca 1 para a esquerda e depois 3 para a direita
	fmt.Printf("(%d << 1) >> 3 = %d (Binário: %08b)\n", num3, leftAndRight, leftAndRight)

	// Usando deslocamento em flags (marcando bits específicos)
	var flag int = 1 << 3 // Ativando o bit 3 (00001000)
	fmt.Printf("Flag ativado no bit 3: %d (Binário: %08b)\n", flag, flag)

	// Multiplicação e divisão por potências de 2 usando deslocamento
	var num4 int = 5
	multiplyBy4 := num4 << 2 // Multiplica por 2^2 = 4
	divideBy2 := num4 >> 1   // Divide por 2^1 = 2
	fmt.Printf("%d * 4 = %d (Deslocamento << 2)\n", num4, multiplyBy4)
	fmt.Printf("%d / 2 = %d (Deslocamento >> 1)\n", num4, divideBy2)
}