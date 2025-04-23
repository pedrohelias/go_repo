package main

import "fmt"

func main(){
	nome := "Pedro Helias Carlos"

	fmt.Printf("meu nome é %v \n", nome)
	fmt.Printf("meu nome é %v \n", "Pedro Helias Carlos")

	//o %v se utiliza quando naõ se tem noção da tipagem, entao, para strings %s, inteiros %d, float %f, casas decimais .%.2f por exemplo.
	

}