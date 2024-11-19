package main

import "fmt"

 /**
 *	VARIÁVEL COMO PACKAGE LEVEL SCOPE
 */
var variavel = 10

func variablesInGo() {
	/*
	*	DECLARAÇÃO DE VARIÁVEL
	*	PALAVRA CHAVE VAR
	*	É DECLARADA EM UM CODE BLOCK E UNDEFINED EM OUTRA
	*	OU SEJA É COMO O LET DO JS
	*	MAS PODE SER PACKAGE LEVEL SCOPE
	*   USANDO VAR É POSSÍVEL DECLARAR E ATRIBUIR VALOR DEPOIS
	*	SE DECLARAR E NÃO ATRIBUIR SENDO EM PACKAGE LEVEL SCOPE, A ATRIBUIÇÃO SÓ PODERÁ SER FEITA
	*	DENTRO DE UM CODE BLOCK
	**/
	newVar := 32
	soma(variavel, newVar)


	/**
	*	CONSTANTES EM GO
	*	CONSTANTES NÃO TIPADAS SÓ TERÃO ALGUM TIPO DE VALOR QUANDO FOREM USADAS
	*/
   const pi = 3.14159
   const greeting = "Hello, World!"
   fmt.Println("Constante pi:", pi)
   fmt.Println("Constante greeting:", greeting)

   //Declarando e definindo tipo ao utilizar.
   const x = 10

   fmt.Println("Constante x:", x)
   fmt.Println("Constante y:", y)


   /**
	*	iota
	*	SUCESSIVAS CONSTANTES NÃO TIPADAS INTEIRAS -> NÃO TIPADAS POIS O TIPO SERÁ DEFINIDO NESSE CASO QUANDO USAR.
	*	OBS.: PODE SER USADO QUANDO VOCÊ NÃO LIGA PARA QUAL VALOR SERÁ ATRIBUÍDO, SÓ QUE QUE SEJA DIFERENTE DOS OUTROS E NÃO CONTROLAR NA MÃO
	*/
   const (
	   _ = iota //aqui descartei o zero pois não vou usar
	   b = iota
	   c = iota
	   d = iota
   ) 
   fmt.Println(b, c, d)
}

func soma(x int, y int) {
	//funciona
	fmt.Println(variavel)
	//erro
	//fmt.Println(newVar)

	//funciona
	fmt.Println(x + y)
}