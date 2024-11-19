/**
*	GO NÃO TEM OPERADOR TERNÁRIO
* COMO DECISÃO DE MANTER UM ESTILO DE CÓDIGO MAIS LIMPO
*
*	GO NÃO TEM WHILE E DO-WHILE
*	GO TEM FOR E POSSUI CONTROLE OR CONTINUE E BREAK 
*	
*	GO POSSUI GOTO MAS NÃO É BOA PRÁTICA UTILIZAR
*
*	OPERADORES LÓGICOS EM GO -> && E ||
*/

package main

import "fmt"

func controlFlow() {
	//Não é boa prática utilizar
	fmt.Println("Início")
    
    goto PularParaAqui // Desvia para o rótulo "PularParaAqui"

    fmt.Println("Este código será ignorado")

PularParaAqui:
    fmt.Println("Agora estamos no rótulo 'PularParaAqui'")


	//Condicionais padrão Obs.: podem ter ()
	idade := 20
    if idade >= 18 {
        fmt.Println("Você é maior de idade.")
    } else {
        fmt.Println("Você é menor de idade.")
    }

	nome := "foo"
	if !(nome == "bar") {
		fmt.Println("nomes diferentes")
	} else {
		fmt.Println("nomes iguais")
	}

	dia := "terça-feira"
    switch dia {
    case "segunda-feira":
        fmt.Println("Hoje é segunda-feira.")
    case "terça-feira":
        fmt.Println("Hoje é terça-feira.")
    case "quarta-feira":
        fmt.Println("Hoje é quarta-feira.")
    default:
        fmt.Println("Dia inválido.")
    }

    // Simulando o comportamento do 'while'
	i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }

    // Simulando o comportamento do 'do-while'
	i = 0
    for {
        fmt.Println(i)
        i++
        if i >= 5 {
            break
        }
    }

	// 'for' tradicional com inicialização, condição e incremento
    for j := 0; j < 5; j++ {
        fmt.Println(j)
    }

	//é possível gerar um loop infinito sem condições como em um servidor http
	for {
		fmt.Println("infinito")
	}

    // Definindo um slice de inteiros
    numeros := []int{10, 20, 30, 40, 50}
    
    // Apenas os valores (ignorando o índice)
     for _, valor := range numeros {
        fmt.Printf("Valor: %d\n", valor)
    }

    // Apenas os índices (ignorando os valores)
    for indice := range numeros {
        fmt.Printf("Índice: %d\n", indice)
    }

    //usando comma ok idiom
    // Criando um map de ferramentas e suas descrições
    ferramentas := map[string]string{
        "Martelo":      "Usado para pregar ou quebrar objetos.",
        "Alicate":      "Usado para segurar ou cortar fios.",
        "Chave de Fenda": "Usada para apertar ou soltar parafusos.",
    }

    // Ferramenta que queremos buscar
    ferramenta := "Serrote"

    // Usando o idioma comma ok para verificar a presença da chave
    if descricao, ok := ferramentas[ferramenta]; ok {
        fmt.Printf("Ferramenta encontrada: %s - %s\n", ferramenta, descricao)
    } else {
        fmt.Printf("Ferramenta '%s' não encontrada no mapa.\n", ferramenta)
    }
}