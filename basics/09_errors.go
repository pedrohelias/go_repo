package main

import (
	"fmt"
	"strconv"
)

func main(){
	//fmt.Println("teste")

	//o erro no go é tratado de forma semelhante ao try catch, mas como variavel:
	//ex:
	// user, err := getUser()
	// if err != nil {
	// 	 fmt.Println(err)
	//   return 
	// }

	/*
	func getUser() (User, error){
		funções ja indicando a possivel saida de erro se ocorrer algum problema
	}
	

	*/

	i, err := strconv.Atoi("42b")
	if err != nil {
		fmt.Println("deu algum problema: ", err) //resposta deu algum problema:  strconv.Atoi: parsing "42b": invalid syntax
		return 
	}

	fmt.Println(i) //se nao tiver erro, vai imprimir aqui. Mas como forçamos o erro, noa vai passar.
	
}	