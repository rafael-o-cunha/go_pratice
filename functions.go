/**
*	FUNÇÕES EM GO
*
*/

package main

import "fmt"

func functions() {


}

//  O protótipo de uma função em GO possui 
//			func somar(a int, b int) int {}
//
//func		->  palavra chave
//somar 	->	numeDaFuncao
//params	->	parâmestros que a funação recebe seguidos de seus tipos
//int		->	o tipo de retorno da função (pode haver mais de um)

// FUNÇÃO SEM PARAM E SEM RETORNO
func digaOla() {
    fmt.Println("Olá, mundo!")
}

//FUNÇÃO COM PARAMS E RETORNOS
//Definindo tipo dos params uma vez já que são iguais
//Realizando 2 retornos diferentes sendo o segundo escolhido para refletir erro
func divisao(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("não é possível dividir por zero")
    }
    return a / b, nil
}

// Função que determina o tipo de triângulo
//Realizando 2 retornos diferentes sendo o segundo escolhido para refletir erro
func tipoTriangulo(a, b, c float64) (string, error) {
	// Verifica se os lados formam um triângulo válido
	if a <= 0 || b <= 0 || c <= 0 {
		return "", fmt.Errorf("Lados inválidos. Os lados devem ser maiores que zero.")
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return "", fmt.Errorf("Não é um triângulo válido.")
	}

	// Determina o tipo de triângulo
	if a == b && b == c {
		return "Triângulo Equilátero", nil
	} else if a == b || a == c || b == c {
		return "Triângulo Isósceles", nil
	} else {
		return "Triângulo Escaleno", nil
	}
}

//FUNÇÃO VARIÁDICA
//Com uso do operador ... ela pode receber N argumentos.
//Mesmo quando a função recebe vários params, esse param pode ser usado, desde que seja o último param.
func soma(numeros ...int) int {
    total := 0
    for _, numero := range numeros {
        total += numero
    }
    return total
}

//DEFER STATEMENT
//empilha funções para serem executadas depois.
func printForaOrdem() {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	//saída -> 1, 3, 2
}

//MÉTODO  = FUNÇÃO ANEXADA A UM TIPO
// Definindo uma struct chamada "Retangulo"
type Retangulo struct {
	Largura, Altura float64
}

// Método para calcular a área do Retângulo
// O receptor (r Retangulo) significa que o método pertence à struct Retangulo
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

// Método para calcular o perímetro do Retângulo
func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

// Método com receptor de ponteiro para alterar o retângulo
func (r *Retangulo) Redimensionar(novaLargura, novaAltura float64) {
	r.Largura = novaLargura
	r.Altura = novaAltura
}

//INTERFACES E POLIMORFISMO
//Um pouco de POO para mostrar uso de métodos em Go
// Gato implementa a interface Som automaticamente porque possui o método EmitirSom().
// Não há necessidade de "registrar" ou "declarar" explicitamente que Gato usa a interface.
type Som interface {
    emitirSom()
}
type Animal struct {
	nome string
}

type Cachorro struct {
	Animal
	espécie string
}

type Gato struct {
	Animal
	espécie string
}

func (a Gato) emitirSom() {
	fmt.Println(a.Animal.nome," miando")
}

func (a Cachorro) emitirSom() {
	fmt.Println(a.Animal.nome," latindo")
}

//outra maneira é injetar a inteface e não outros tipos de variável comuns... logo o tipo que implementar a interface pode ser passado como argumento
// Função que aceita a interface Animal
func FazerSom(a Animal) {
    a.EmitirSom()
}
/*
func main() {
    cachorro := Cachorro{nome: "Rex"}
    gato := Gato{nome: "Felix"}

    // Injetando instâncias que implementam a interface Animal
    FazerSom(cachorro) // "Rex latindo"
    FazerSom(gato)     // "Felix miando"
}
*/

//FUNÇÕES ANÔNIMAS EM GO
func useAnonimousFuncs() {
    // Função anônima atribuída a uma variável (função como expressão)
    saudacao := func(nome string) {
        fmt.Println("Olá,", nome)
    }
    
    // Chamando a função anônima
    saudacao("João")
    
    // Passando uma função anônima diretamente como argumento
    resultado := operacao(5, 3, func(a, b int) int {
        return a + b
    })
    fmt.Println("Resultado da soma:", resultado)
    
    // Função anônima com retorno
    dobro := func(x int) int {
        return x * 2
    }
    fmt.Println("Dobro de 4:", dobro(4))

}

// Função que recebe uma função como parâmetro
func operacao(a, b int, f func(int, int) int) int {
    return f(a, b)
}

//função anônima autoinvocável
xzinho := 10
func(x int) {
	fmt.printf("xzinho x 3: %v\n", xzinho * 3)
}(xzinho)

//função como retorno
zzinho := func returnFunc() func(int) int {
	return func(i int) int {
		return x * 10
	}
}

//Função como expressão sem variável e retornando outra função
//antes colocou zzinho recebendo o retorno da returnFunc
//depois contando com isso recebe xzinho como argumento da função interna.
fmt.Println(zzinho()(xzinho))



//FUNÇÃO DE CALLBACK
// Função de callback que realiza uma operação matemática
func operacao(a, b int, callback func(int, int) int) int {
    return callback(a, b)
}

somar := func(a, b int) int {
	return a + b
}

multiplicar := func(a, b int) int {
	return a * b
}

// Chamando a função operacao passando a função somar como callback
resultadoSoma := operacao(5, 3, somar)
fmt.Println("Resultado da soma:", resultadoSoma)

// Chamando a função operacao passando a função multiplicar como callback
resultadoMultiplicacao := operacao(5, 3, multiplicar)
fmt.Println("Resultado da multiplicação:", resultadoMultiplicacao)


//CLOUSURE
// contator está no escopo de "criarContador" mesmo que a função interna o utilize e retorne seu estado mais atual, ele é da função externa.
// Função que retorna uma closure
func criarContador() func() int {
    contador := 0
    return func() int {
        contador++
        return contador
    }
}
/**
*		A função criarContador cria uma variável contador e retorna uma função anônima que, quando chamada, incrementa e retorna o valor de contador.
*		A função anônima tem acesso à variável contador, mesmo depois que criarContador já tenha retornado.
*		Isso é o que caracteriza uma closure: a função "lembra" do ambiente em que foi criada e consegue acessar as variáveis externas a ela.
* 
*		Cada chamada para criarContador cria uma nova closure, com seu próprio contador. Isso significa que, mesmo que criarContador tenha sido 
*  chamada várias vezes, cada closure "lembra" do valor de contador associado a ela, e eles não compartilham o mesmo valor.
*
*		Closures são muito úteis quando se deseja criar um contador ou gerador de números que mantém o estado entre as chamadas, sem precisar de uma
*	variável global. (Como visto abaixo)
*		Closures são frequentemente usadas em funções de callback para capturar variáveis locais e fornecer comportamento personalizado dentro de
*	funções externas. (Como visto anteriormente)
*
*/
func usingClousureConcecpt() {
    // Criando um contador
    contador1 := criarContador() 
    contador2 := criarContador() a

    // Usando o contador1
    fmt.Println("Contador 1:", contador1()) // 1
    fmt.Println("Contador 1:", contador1()) // 2

    // Usando o contador2
    fmt.Println("Contador 2:", contador2()) // 1
    fmt.Println("Contador 2:", contador2()) // 2

    // Mostrando que cada contador tem seu próprio estado
    fmt.Println("Contador 1:", contador1()) // 3
    fmt.Println("Contador 2:", contador2()) // 3
}


//RECURSIVIDADE
// Função recursiva para calcular o fatorial de um número
func fatorial(n int) int {
    if n == 0 {
        return 1 // Caso base: fatorial de 0 é 1
    }
    return n * fatorial(n-1) // Chamada recursiva
}

func usingFactorialFunc() {
	numero := 5
    fmt.Printf("O fatorial de %d é %d\n", numero, fatorial(numero))	
}