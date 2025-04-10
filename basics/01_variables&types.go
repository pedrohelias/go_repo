// o operador de declaração é o gopher := . Você declara variaveis com o gopher
//a tipagem é automatica. O Operador curto, o gopher, atribui um tipo para a variavel

//o gopher funciona apenas dentro dos codeblocks, ou seja, espaços onde o codigo realmente é feito. No nosso caso, apenas dentro da funcao main -> CONTINUA

package main

import "fmt"

//se eu faço igual o trecho abaixo, a variavel estará presente e disponivel em todo o espaço do arquivo, proximo de ser global, podendo ser utilizdo o =. Portanto, não dará erro

var z = 10 //aqui

//e por ser algo fora do codeblocks, não funciona o gopher. E ai necessita usar o var

func main(){
	x := 10
	y := "Ola, bom dia"
	fmt.Println(x, "e", y)


	//é possivel tmb usar o Printf

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	//podemos utilizar o =, mas ele só é possível quando se deseja atribuir um valor novo a uma variavel ja existente. por exemplo:

	x = 20
	fmt.Printf("x: %v, %T\n", x, x)

	//caso eu tente usar o gopher para atribuir um novo valor, ele dará erro x := 20


}
