/**
*	GO POSSUI ARRAYS, SLICES(CONSTRUÍDO SOBRE ARRAY), MAPS, STRUCTS
*	ENTENDER STRUCT NESTE CASO COMO AGRUPAMENTO DE DADO EM CONTEXTO
*
*	ARRAY TEM UM TAMANHO DEFINIDO E SLICE NÃO.
*	DEVE-SE TER CUIDADO COM RESLICE, POIS ELE SOBRESCREVE MEMÓRIA (OVERLAP)
*/

package main

import "fmt"

/**
* PODE-SE TER STRUCTS DENTRO DE STRUCTS, EM OUTRAS PALAVRAS: COMPOSIÇÃO
*/
type Endereco struct {
	Rua       string
	Numero    int
	Bairro    string
	Cidade    string
	Estado    string
	CEP       string
}

type Pessoa struct {
    Name string
    Age  int
	Adress Endereco
}




func dataCollections() {
	// 1. ARRAY: TAMANHO FIXO E TIPOS HOMOGÊNEOS
	//https://go.dev/doc/effective_go#arrays
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println("Array:", numeros)


	// 2. SLICE: TAMANHO VARIÁVEL, BASEADO EM ARRAYS
	//https://go.dev/doc/effective_go#slices
	slice := []string{"Maçã", "Banana", "Laranja"}
	slice = append(slice, "Uva")
	fmt.Println("Slice:", slice)


	// 2.1. ACESSANDO/MODIFICANDO APENAS UM SUBCONJUNTO DE UM SLICE
    frutas := []string{"Maçã", "Banana", "Cereja", "Damasco", "Figo"}

    // Acessando uma fatia do slice (elementos 1 a 3, exclusivo do índice final)
    fatia := frutas[1:4]

    fmt.Println("Slice original:", frutas)
    fmt.Println("Fatia do slice:", fatia)

    // Modificando a fatia e observando o impacto no slice original
    fatia[0] = "Manga"
    fmt.Println("Após modificar a fatia:")
    fmt.Println("Slice original:", frutas)
    fmt.Println("Fatia do slice:", fatia)

	// 2.2. APPEND E ... PARA EXPANSÃO E COMBINAÇÃO DE SLICES
    nums := []int{1, 2, 3}
    fmt.Println("Slice inicial:", nums)

    // Adicionando um único elemento
    nums = append(nums, 4)
    fmt.Println("Após append de um elemento:", nums)

    // Adicionando vários elementos
    nums = append(nums, 5, 6, 7)
    fmt.Println("Após append de múltiplos elementos:", nums)

    // Criando outro slice para combinar
    maisNumeros := []int{8, 9, 10}

    // Usando append com o operador ... para expandir o slice
    nums = append(nums, maisNumeros...)
    fmt.Println("Após append com operador ...:", nums)

	// 2.3. USANDO MAKE COM SLICES (SLICE TEM LEN E CAPACITY) uso do make é eficiente para evitar overhead de alocação quando sabemos aproximadamente o tamanho do slice a ser usado.
	// Criando um slice de strings com capacidade inicial usando make
    ferramentas := make([]string, 0, 5) // Slice vazio com capacidade para 5 elementos

    fmt.Println("Slice inicial:", ferramentas)
    fmt.Printf("Tamanho: %d, Capacidade: %d\n", len(ferramentas), cap(ferramentas))

    // Adicionando nomes de ferramentas ao slice
    ferramentas = append(ferramentas, "Martelo", "Chave de Fenda", "Alicate")
    fmt.Println("Após adicionar ferramentas:", ferramentas)
    fmt.Printf("Tamanho: %d, Capacidade: %d\n", len(ferramentas), cap(ferramentas))

    // Adicionando mais ferramentas ao slice
    ferramentas = append(ferramentas, "Serrote", "Furadeira", "Parafusadeira")
    fmt.Println("Após adicionar mais ferramentas:", ferramentas)
    fmt.Printf("Tamanho: %d, Capacidade: %d\n", len(ferramentas), cap(ferramentas))

	// 2.4. OVERLAPPING - DEVE-SE TER CUIDADO AO REALIZAR APPEND/RESLICE POIS PODE CAUSAR OVERLAP 
	//uma forma de evitar é inicializar o segundo slice com base em um slice vazio como base para depois apendar
	fmt.Println("Slice original:", ferramentas)

    // Criando um novo slice a partir do original (refatiamento)
    subFerramentas := ferramentas[1:4] // Inclui índices 1, 2 e 3
    fmt.Println("Subslice (refatiamento):", subFerramentas)

    // Modificando o subslice
    subFerramentas[0] = "Chave Inglesa"
    fmt.Println("Subslice modificado:", subFerramentas)

    // Verificando o impacto no slice original
    fmt.Println("Slice original após modificação:", ferramentas)

	
	// 3. MAP: ESTRUTURA CHAVE-VALOR
	//https://go.dev/doc/effective_go#maps
	//caso tente acessar um valor que não existe será retornado 0
	//maps não tem ordem
	mapa := map[string]int{
		"Maçã":    5,
		"Banana":  10,
		"Laranja": 7,
	}
	mapa["Uva"] = 15
	delete(mapa, "Banana")
	fmt.Println("Map:", mapa)

	// 4. STRUCT: TIPO COMPOSTO PARA AGRUPAR VALORES HETEROGÊNEOS
	pessoa := Pessoa{
        Name: "João Silva",
        Age:  35,
        Adress: Endereco{
            Rua:     "Rua das Acácias",
            Numero:  456,
            Bairro:  "Centro",
            Cidade:  "Rio de Janeiro",
            Estado:  "RJ",
            CEP:     "22000-000",
        },
    }
	fmt.Println("Informações da Pessoa:")
	fmt.Println(pessoa)
    fmt.Println("------------")
	fmt.Println("Nome:", pessoa.Name)
	fmt.Println("Endereço:")
    fmt.Println("Rua:", pessoa.Adress.Rua)
    fmt.Println("Número:", pessoa.Adress.Numero)

	//4.1 Criando sem explicitar os nomes dos atributos
	pessoa2 := Pessoa{
        "João Silva",           
        35,                     
        Endereco{               
            "Rua das Acácias",  
            456,                
            "Centro",           
            "Rio de Janeiro",   
            "RJ",               
            "22000-000",         
        },
    }

	//4.2 campos podem ser promovidos se não houver conflito entre a struct externa e a interna
	fmt.Println("Rua:", pessoa.Rua)

	//4.3 - Struct que é criado e utilizado uma vez, é descartável.
	 // Criando e instanciando uma struct anônima para representar uma conta bancária
	 contaBancaria := struct {
        NumeroDaConta string
        Titular       string
        Saldo         float64
    }{
        NumeroDaConta: "12345-6",
        Titular:       "Carlos Pereira",
        Saldo:         1500.75,
    }

    // Imprimindo as informações da conta bancária
    fmt.Println("Informações da Conta Bancária:")
    fmt.Println("Número da Conta:", contaBancaria.NumeroDaConta)
    fmt.Println("Titular:", contaBancaria.Titular)
    fmt.Println("Saldo:", contaBancaria.Saldo)
}