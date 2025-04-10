package main // dizendo que é o pacote principal

import "fmt" //assim qeu se importa apenas uma lib
/*
assim que se importa varias:

import {
	"lib1"
	"lib2"
	...
}


*/


func main(){ //porta de entrada e saide de qualquer programa em go
	fmt.Println("Hello world!")
	fmt.Println("Oi", "teste", 456) //da pra imprimr assim tmb, e numeros
}