/**
*variáveis, valores e tipos, fluxo de controle
 */

 package main

 /**
 *	IMPORT SIMPLES EM APENAS UMA LINHA
  */
//import "fmt"


 /**
 *	IMPORT DE MÚLTIPLOS ARTEFATOS		OBS: se tiver import simples e import múltiplo e usar 
 *										um auto formatter, o simples sofrerá replace para dentro do múltiplo
  */
import (
	"fmt"
	"strconv"
 )
 
 /**
 *	COMO NO C, A FUNÇÃO MAIN É O ENTRYPOINT DA APLICAÇÃO
  */
func main() {
	 returnOfFunctions()
	 typesInGo()
	 printData()
	 operators()
	 variablesInGo()
	 bitWiseOperations()
	 controlFlow()
	 dataCollections()
	 genericsInGo()
	 functions()
	 pointersInGO()
}
