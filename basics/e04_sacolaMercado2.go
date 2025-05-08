package main

import "fmt"

type product struct { //para fazer um map mais complexo
	qtd int
	code string
	isAlive bool
	valor float64
}

func listaProdutos(produtos map[string]float64){
	counter := 0
	fmt.Println("Tabela de Produtos")
	for chave, valor := range produtos {
		counter++
		fmt.Printf("%d - %s, preço: R$ %.2f\n", counter, chave, valor)
	}
	fmt.Println("______________________")
}

func listaProdutosComplex(produtos map[string]product){
	counter := 0
	fmt.Println("Tabela de Produtos")
	for chave, valor := range produtos {
		counter++
		fmt.Printf("%d - %s, preço: R$ %.2f, quantidade: %d, válido? %v\n", counter, chave, valor.valor, valor.qtd, valor.isAlive)
	}
	fmt.Println("______________________")
}

func addProdutos(produto string, preco float64, produtos map[string]float64) map[string]float64{
	produtos[produto] = preco
	return produtos
}

func removerProdutos(produtos map[string]float64, chaveRem string) map[string]float64{
	for chave:= range produtos {
		if (chave == chaveRem){
			delete(produtos, chaveRem)
		}

	}
	return produtos
}

func atualizaValor(produtos map[string]float64, chaveA string, valorA float64) map[string]float64{
	for chave := range produtos{
		if chave == chaveA {
			produtos[chave] = valorA
		}
	}
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

	produtosComplexos := map[string]product{ //aqui vai permitir que tenhamos objetos e construções mais complexas
		"uva":{
			qtd: 3,
			code: "12345",
			isAlive: true,
			valor: 14.90,
		},
		"maca":{
			qtd: 4,
			code: "54321",
			isAlive: false,
			valor: 12.00,
		},
	}

	listaProdutos(produtos)
	produtos = addProdutos(fruta, preco, produtos)
	listaProdutos(produtos)
	produtos = removerProdutos(produtos, "uva")
	listaProdutos(produtos)
	produtos = atualizaValor(produtos, "laranja", 4.60)
	listaProdutos(produtos)
	listaProdutosComplex(produtosComplexos)



	fmt.Println("teste")
}