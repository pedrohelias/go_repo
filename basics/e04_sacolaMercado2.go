package main

import "fmt"


func listaProdutos(produtos map[string]float64){
	counter := 0
	fmt.Println("Tabela de Produtos")
	for chave, valor := range produtos {
		counter++
		fmt.Printf("%d - %s, preço: R$ %.2f\n", counter, chave, valor)
	}
	fmt.Println("______________________")
}

func addProdutos(produto string, preco float64, produtos map[string]float64) map[string]float64{
	produtos[produto] = preco
	return produtos
}

func main(){

	produtos := make(map[string]float64)
	fruta:= "pera"
	preco := 2.99

	produtos = map[string]float64{
		"uva":10.50,
		"maça":20.80,
		"laranja":5.56,
	}

	listaProdutos(produtos)
	produtos = addProdutos(fruta, preco, produtos)
	listaProdutos(produtos)


	fmt.Println("teste")
}