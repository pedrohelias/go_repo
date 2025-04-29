package main

//adicionar frutas
//retirar frutas

import (
	"fmt"
	//"slices"
)

func addFruta (lista []string, fruta string) ([]string) {
	novaLista := append(lista, fruta)
	return novaLista
}

func removerFruta (lista []string, indice int) ([]string){
	return append(lista[:indice], lista[indice+1:]...)
}

func mostrarLista (lista []string){
	for i := 0; i < len(lista);i++{
		fmt.Printf("%v - %v\n", i, lista[i] )
	}
	fmt.Println("____________________________________")
}
func main(){

	listaFrutas := []string{"limão", "melão", "banana", "uva", "maça"}
	mostrarLista(listaFrutas)
	novaLista := addFruta(listaFrutas, "goiaba")
	mostrarLista(novaLista)
	novaLista = removerFruta(novaLista, 2)
	mostrarLista(novaLista)



	
}