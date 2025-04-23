package main

import "fmt"

func main(){
	//condicionais similares ao c mas naõ tanto rebuscada

	height := 4

	if height > 5 {
		fmt.Println("é isso")
	}else{
		fmt.Println("Não é isso")
	}


	//os operadores de cmoparação sao iguais aos do c
	//é possivel usar statement condition
	//ao inves disso:
	/*

	altura := 10
	if altura > 5 {
		fmt.Println("muito alto")
	}else{
		fmt.Println("ate que esta ok")
		}
	*/

	//escreva assim

	if altura := 10; altura >5 {  //if INITIAL STATEMENT; CONDITION
		fmt.Println("muito alto")
	}else{
		fmt.Println("ate que esta ok")
	}

	//isso simplifica o codigo
}