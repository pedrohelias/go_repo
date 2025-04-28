package main 

import "fmt"

func mostraParcelamento (nParcelas int, valor float64){
	valorParceladoTotal := 0.0
	for i := 1; i <= nParcelas; i++ {
		valorParceladoTotal = valorParceladoTotal + valor
		fmt.Printf("No mês %v, o parcelamento fica %v, com o total arrecadado de: %.2f\n", i, valor, valorParceladoTotal)
	}
}
//exemplo while no go
func contagem (valorInicial, fim int )(iteracoes int){
	for valorInicial < fim{
		iteracoes++
		valorInicial++
		fmt.Printf("numero de iteracoes: %v, valor atual: %v\n", iteracoes, valorInicial)
	}

	return iteracoes
}


func main(){
	
	//loop for -> for inicial; condicao; depois

	for i:=1; i <= 10; i++{
		fmt.Printf("repetição número: %v - teste\n", i)
	}
	//n := 12
	//valor := 45.60
	//mostraParcelamento(n, valor)

	//nao existe while loop em go, para suprimir isso, utiliza-se o for. Entao um while em Go é um for com apenas uma condições


	valorInicial := 10
	limite := 20
	contagem := contagem(valorInicial, limite)
	fmt.Printf("a quantiade de iterações foi: %v\n", contagem)

	//continue e break no Go. Diferença primoridal: Enquanto o continue paraliza uma iteração para seguir em dirençaõ a outra, o break encerra a iteração e sai do laço


}

