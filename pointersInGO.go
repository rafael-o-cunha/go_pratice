/**
*	PONTEIROS EM GO
*		Em Go, ponteiros são usados para referenciar a memória de uma variável, permitindo que você
*	altere o valor original da variável dentro de uma função ou compartilhe dados de maneira eficiente,
*	sem a necessidade de fazer cópias. Um ponteiro em Go é declarado com o operador * e um endereço de
*	memória é armazenado utilizando o operador &.
*
*	SEMELHANTE EM C
*		Usado em muitos casos em aplicações reais para operar dados por referência e consumir a alterar contextos
* 	dessa forma GO que é pass by value pode operar by Ref
*/

package main

import "fmt"

func pointersInGO() {
	a := 10
	b := &a
    fmt.Println("Antes:", a)

	alterarValor(&a)

    fmt.Println("Depois:", a)

	//Imprime: int 	 *int
	fmt.Printf("%T \t %T\n", a, b)
}

func alterarValor(p *int) {
    *p = 20 // Modifica o valor da variável que o ponteiro aponta
}