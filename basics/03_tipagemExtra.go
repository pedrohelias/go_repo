package main

import (
"fmt"
"reflect"
)

func main(){
	fmt.Println("teste!") //formatadores: %v para uma forma genérica (nao muito recomendado, mas funciona), %s string, %d decimal inteiro, %q entre aspas

	//tipagens:
	//var nome_variavel tipo_variavel (pode inicializar sem um valor, entao sua saida sera o zero - OU iniciar com um valor)
	//ex:
	/*

	STRINGS
	var exemplo string = "string"
	var exemploStringvarzia string 
	string := "String"
	stringConc := string + " " + exemplo
	
	FLOAT 

	var exemploFloat float32 = 3.14
	float := 3.14 -> esse float força o float 64
	Para operações matemáticas entre floats, precisam ter o mesmo tipo. Se necessario, alterar o tipo do valor

	BOOLEAN

	var ligado bool = false
	se coloar ! ao trabalhar com essa variavel, ela inverte o tipo

	ARRAYS

	var arrayExemplo [tamanho] tipo
	var frutas [3] string -> 3 indices, apenas

	frutas[0] = "uva"
	frutas[1] = "banana"
	frutas[2] = "melao"

	primeiraFruta := frutas[0]
	fmt.printLn(primeiraFruta)
	fmt.printLn(frutas[1])

	OU

	numeros := [5]int{10,20,30,40,50}
	bools := [3]bool{true, false, false}

	podemos ter arrays de arrays:
	array := [3][3]string { {"Vermelho", "azul", "verde"}, {"outrroVermelho", "outracor", "cor"}, {"CorCor", "corzao", "corzinha"} }
	printando:
	fmt.printLn(array[0][1]) //procura o array de indice 0, e o objeto no arry de indice 1, no caso "azul"



	
	*/

	//array := [3][3]string { {"Vermelho", "azul", "verde"}, {"outrroVermelho", "outracor", "cor"}, {"CorCor", "corzao", "corzinha"} }
	numeros := [5]int{10,20,30,40,50}


	for i := 0; i < len(numeros); i++{
		fmt.Println(numeros[i])
	
		var stringTeste string = "string"
		fmt.Println(stringTeste)
	}


	//conversao de valores:
	//utilizAar o valor em parentes para isso ->float64(temperatura)

	temperatura := 10
	fmt.Println(reflect.TypeOf(temperatura)) //tipo int

	temperaturaConvert := float64(temperatura)
	fmt.Println(reflect.TypeOf(temperaturaConvert)) //tipo float64


	//tipos sugeridos: bool, string, int, uint32, byte, rune, float64, complex 128

	//constantes: não utilizam :=, precisam ser declaradas como variaveis normais, mas usam o nome const -> const nome = "Pedro Helias Carlos". E como o proprio nome diz, elas são imutaveis, voce nao altera os dados internos de uma const
	
}