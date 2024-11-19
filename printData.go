package main

import "fmt"

func printData() {
	x := 16
	y := "texto"
	/*
	*	PRINT FORMATADO
	*	PARA IMPRIMIR FORMATADO É BEM PRÓXIMO DE COMO É FEITO EM C
	*	NO CASO ABAIXO ESTÁ SENDO IMPRESSO O VALOR E O TIPO USANDO O PRINT FORMATADO
	*	https://pkg.go.dev/fmt#hdr-Printing
	**/
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	/**
	*	USANDO OUTROS PRINTS
	*	Sprint -> FAZ UM REDIRECT PARA UMA SPRINT (PRINT STRING)
	*	Fprint -> FAZ PRINT PARA UM FILE OU COMO DEVE SER CHAMADO: PARA UM FILE DESCRIPTOR (SERÁ VISTO MAIS A FRENTE SEU USO DURANTE PRÁTICA COM ARQUIVOS)
	 */
	msg := fmt.Sprintf("Este é um exemplo de %s com valor formatado: %d", "fmt.Sprintf", 42)
	fmt.Println(msg)
	msg = fmt.Sprint("Este é um exemplo de", " ", "join de strings", " ", "usando o sprint")
	fmt.Println(msg)

   /**
	*	STRING É UM SLICE OF BYTES 
	*  %v:  Este campo imprime o valor numérico bruto do byte (no caso, a tabela ASCII).
	*  %T:  Este campo imprime o tipo da variável.
	*  %#U: Este campo imprime o ponto de código Unicode em formato legível:
	*  %#x: Este campo imprime o valor hexadecimal do byte.
	*
	*	OBS.: SE FIZER RANGE EM UMA STRING IRÁ OBTER CARACTERE POR CARACTERE, OU RUNE
	*		  SE FIZER FOR VAI TER BYTE POR BYTE E DEVIDO A ISSO CARACTERES COM MÚLTIPLOS BYTES TERÃO APENAS O PRIMEIRO BYTE RETORNADO.
	* https://go.dev/blog/strings
	*/
	sliceY := []byte(y)
	for _, v := range sliceY {
	   fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}
   
	fmt.Println(" ")

   for i := 0; i < len(sliceY); i++ {
	   v := sliceY[i]
	   fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
   }


}