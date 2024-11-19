package main

import "fmt"

 // DEFININDO UM TIPO COMPOSTO (STRUCT)
 type Coordinate struct {
	x int
	y int
}

func typesInGo() {
	/**
	*	TIPOS PRIMITIVOS DA LINGUAGEM
	*	EXEMPĹO ABAIXO USA O OPERADOR CURTO DE ATRIBUIÇÃO := OU GHOPER
	*	DECLARANDO E ATRIBUINDO E ELE FAZ INFERÊNCIA DE TIPO
	*
	 */
	x := 16
	y := "texto"
	z := true
	fmt.Println(x, y, z)

	/**
	*	TIPOS
	*	TIPOS SÃO ESTÁTICOS -> DECLARANDO COMO INT SERÁ SEMPRE INT
	*	USANDO O GHOPER É REALIZADO INFERÊNCIA
	 */
	var typeVar int = 10
	fmt.Println(typeVar)

	var typeVar2 int = 42
	fmt.Println(typeVar2)

	/**
	*	TIPOS
	*	TIPOS COMPOSTOS  = COMPOSTOS DE TIPOS PRIMITIVOS (TAMBÉM PODEM SER DEFINIDOS PELO USUÁRIO)
	*	SLICE, ARRAY, STRUCT E MAP
	*	O ATO DE DEFINIR, CRIAR, ESTRUTURAR TIPOS COMPOSTOS CHAMA-SE COMPOSIÇÃO.
	 */
	// 1. ARRAY: TAMANHO FIXO E TIPOS HOMOGÊNEOS
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println("Array:", numeros)

	// 2. SLICE: TAMANHO VARIÁVEL, BASEADO EM ARRAYS
	slice := []string{"Maçã", "Banana", "Laranja"}
	slice = append(slice, "Uva")
	fmt.Println("Slice:", slice)

	// 4. MAP: ESTRUTURA CHAVE-VALOR
	mapa := map[string]int{
		"Maçã":    5,
		"Banana":  10,
		"Laranja": 7,
	}
	mapa["Uva"] = 15
	fmt.Println("Map:", mapa)

	//IMPRIMINDO COORDENADAS
	printCoordinate()
	printNCoordinates()

	/**
	*	VALOR DEFAULT DE TIPOS
	*	VALOR DEFAULT DE INICIALIZAÇÃO AO CRIAR UMA VARIÁVEL E NÃO ATRIBUIR VALOR.
	*	PARA POINTERS, FUNCTIONS, INTERFACES, SLICES, CHANNELS, MAPS?: nil
	 */
	var inteiro int
	var flutuante float64
	var booleano bool
	var texto string
	var coordenada Coordinate
	fmt.Println("Valores padrão de inicialização:")
	fmt.Printf("int: %d\n", inteiro)
	fmt.Printf("float64: %f\n", flutuante)
	fmt.Printf("bool: %t\n", booleano)
	fmt.Printf("string: '%s'\n", texto)
	fmt.Printf("Coordinate: %+v\n", coordenada) //    %v: Exibe apenas os valores da struct. e %+v: Exibe os nomes dos campos e seus respectivos valores.

	/**
	*	STRINGS EM GO
	*	EXISTEM 2 TIPOS DE STRINGS: INTERPRETED STRING LITERALS E RAW STRING LITERALS
	*	RUNE LITERALS -> CADA CARACTERE É UMA RUNA, ELA PODE TER CHARSET, ENCODING, E QUANTIDADE DE BYTES VARIÁVEL
	*	UM LITERAL É UMA NOTAÇÃO PÁRA REPRESENTAR UM VALOR FIXO NO CÓDIGO FONTE.
	*	UMA STRING INTERPRETADA É A MAIS COMUM, E OS CARACTERES DE ESCAPE SÃO INTERPRETADOS
	*	UMA RAW STRING É DEFINIDA USANDO CRASES NO INÍCIO E FIM E OS CARACTERES DE ESCAPE NÃO SÃO INTERPRETADOS DIRETAMENTE.
	*  https://go.dev/blog/strings
	*/
	animal := "gato\nfrajola"
	rawAnimal := `"gato\nfrajola"`
	rawAnimal2 := `"gato\nfrajola"
					miou,
					miando`
	fmt.Println(animal)
	fmt.Println(rawAnimal)
	fmt.Println(rawAnimal2)

	/**
	*	TIPO BOOL EM GO
	*	TIPO BINÁRIO QE PODE TER 2 VALORES COMO EM QUALQUER OUTRA LINGUAGEM
	*
	 */
	var tipoBool bool
	fmt.Println(tipoBool)
	fmt.Println(true && false)

	/**
	*	TIPO NUMÉRICO EM GO
	*	HÁ 2 TIPOS NUMÉRICOS PRINCIPAIS: INT E FLOAT
	*	HÁ VARIAÇÕES DESSES TIPOS:
	*		INT  ->	`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`	
	*		FLOAT -> `float32`, `float64`	
	*
	*		O PADRÃO DA MÁQUINA DETERMINA O PADRÃO DO INT UTILIZADO (ARQUITETURA = 64 LOGO O INT É INT64 POR DEFAULT)
	*		OBS.: `byte` É UM ALIAS PARA UINT8 E `rune` É UM ALIAS PARA INT32
	*		OBS.: STRING SÃO FEITAS DE RUNES E CADA RUNE É UM CARACTERE QUE PODE TER 1, 2, 3 OU 4 BYTES
	*		OBS.: USE SEMPRE FLOAT64
	*/
	// Declarando tipos numéricos MAIS UTILIZADOS	
   var numeroInt int = 42\
   var numeroFloat float64 = 3.14159
   fmt.Printf("Valor inteiro (int): %d\n", numeroInt)
   fmt.Printf("Valor flutuante (float64): %f\n", numeroFloat)

   /**
   *	RUNES, BYTES, STRINGS EM BYTES
   *	OBS.: COMO VISTO ANTERIORMENTE -> STRING SÃO FEITAS DE RUNES E CADA RUNE É UM CARACTERE QUE PODE TER 1, 2, 3 OU 4 BYTES
   */
   churros := "e"
   pipoca := "é"
   brigadeiro := "🐹"
   fmt.Printf("%v, %v, %v\n", churros, pipoca, brigadeiro)

   sliceChurros := []byte(churros) 		//observa-se que usa 1 byte
   slicePipoca := []byte(pipoca)   		//observa-se que usa 2 bytes
   sliceBrigadeiro := []byte(brigadeiro)	//observa-se que usa 4 bytes
   fmt.Printf("%v, %v, %v", sliceChurros, slicePipoca, sliceBrigadeiro)


	/**
	*	OVERFLOW
	*	É O MESMO TENTAR USAR NÚMERO LONG EM UMA VARIÁVEL INT EM C, ELE VAI ESTOURAR O LIMITE DO TIPO.
	*/
	//cannot use 99999999999999999999999999999 (untyped int constant) as int32 value in variable declaration (overflows)
	//var numeroInt int32 = 99999999999999999999999999999
	//fmt.Println(numeroInt)

	/**
	*	CRIANDO O PRÓPRIO TIPO
	*	O TIPO ABAIXO SE BASEIA NA STRING
	*	PARA TIPO COMPOSTO E PERSONALIZADO TEMOS AQUI O COORDINATE
	 */
	type cat string
	var gato cat = "milzinho"
	fmt.Println(gato)

	/**
	*	CONVERSÃO DE TIPOS
	*	https://go.dev/ref/spec#Conversions
	 */
	// Definindo as variáveis
	inteiro = 42
	flutuante = 3.14159
	booleano = true
	texto = "123"

	// 1. Conversão de int para float64
	flutuanteConvertido := float64(inteiro)
	fmt.Printf("Inteiro: %d, Convertido para Flutuante: %f\n", inteiro, flutuanteConvertido)

	// 2. Conversão de float64 para int
	inteiroConvertido := int(flutuante)
	fmt.Printf("Flutuante: %f, Convertido para Inteiro: %d\n", flutuante, inteiroConvertido)

	// 3. Conversão de bool para string
	textoConvertido := strconv.FormatBool(booleano)
	fmt.Printf("Booleano: %t, Convertido para Texto: %s\n", booleano, textoConvertido)

	// 4. Conversão de string para int (usando strconv.Atoi)
	numeroConvertido, err := strconv.Atoi(texto)
	if err != nil {
		fmt.Println("Erro ao converter texto para inteiro:", err)
	} else {
		fmt.Printf("Texto: %s, Convertido para Inteiro: %d\n", texto, numeroConvertido)
	}

	// 5. Conversão de string para float64 (usando strconv.ParseFloat)
	flutuanteConvertido2, err := strconv.ParseFloat(texto, 64)
	if err != nil {
		fmt.Println("Erro ao converter texto para flutuante:", err)
	} else {
		fmt.Printf("Texto: %s, Convertido para Flutuante: %f\n", texto, flutuanteConvertido2)
	}
}

func printCoordinate() {
	//3. STRUCT: DEFININDO TIPOS PERSONALIZADOS
	coord := Coordinate{
		x: 42,
		y: 42,
	}
	fmt.Printf("Struct: %+v\n", coord)
}

func printNCoordinates() {
	// COMPOSIÇÃO DE TIPOS: EXEMPLO COM STRUCT E SLICE
	var coordinates []Coordinate
	coordinates = append(coordinates, Coordinate{x: 10, y: 20})
	coordinates = append(coordinates, Coordinate{x: 30, y: 40})
	fmt.Println("Slice de Structs (Coordinates):", coordinates)

	// ITERANDO SOBRE AS COORDENADAS PARA MOSTRAR VALORES INDIVIDUAIS
	for _, coord := range coordinates {
		fmt.Printf("Coordinate: x=%d, y=%d\n", coord.x, coord.y)
	}

}
