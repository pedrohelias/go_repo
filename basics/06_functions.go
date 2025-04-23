package main

import (
	"errors"
	"fmt"
)

//funcoes podem ou nao ter retorno. Caso tenha, precisa de uma tipagem:

/*
funcoes com retorno (o int apos a declaracao das variaveis representam a tipagem):

*/

func somaDoisNumeros(a int, b int) int {
	return a + b
}

// pode ter 2 ou mais tipos de retorno

func retornaNomeSobrenome(a string, b string) (string, string) {
	return a, b //vai retornar duas, dai é necessario ter duas variaveis para obter os dados
}

//quando mais de um argumento é do mesmo tipo
func mesmotipo(a, b, c int) int { 
	return a + b + c
}

//um exemplo mais complexo

func divisao(divisor, dividendo int) (int, error){
	if divisor == 0 {
		return 0, errors.New("Não é possível dividir por zero")
	}

	operation := dividendo / divisor
	return operation, nil //esse nil é para retornar 2 argumentos
}



func main(){
	a := 0
	b := 5
	nome1 := "Pedro"
	nome2 := "Helias"
	soma2return := somaDoisNumeros(a,b)

	retNome1, retNome2 := retornaNomeSobrenome(nome1, nome2)

	fmt.Println(soma2return)
	fmt.Println(retNome1, retNome2)

	//eh possivel ignorar valores de retorno
	//utilizando o retornanomesobreome

	retNome1New, _:= retornaNomeSobrenome(nome1, nome2)
	fmt.Println(retNome1New)

	divisao1, erro := divisao(a, b)

	if(erro == nil){
		fmt.Printf("o resultado da divisao é: %d\n", divisao1)
	}else{
		fmt.Printf("o resultado da divisão é: %d, com o erro: %v\n", divisao1, erro)
	}


	


}