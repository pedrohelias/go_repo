package main

import "fmt"


func soma(numeros ...int)int{
	total := 0
	//para iterar sob coleções, slices,maps, utiliza-se for range
	for _, n := range numeros{ //O primeiro valor (_) é o índice (posição), O segundo valor (n) é o valor da posição atual.
		total += n
	}

	return total
}


func main(){
	fmt.Println("teste")

	//arrays em Go pode ser declarado assim
	//var myInts [10]int
	//ou
	primes := [6]int{2,3,5,7,11,13} //aqui o valor entre chaves é o tamanho do array, após isso a tipagem dos dados, e depois o conteúdo

	for i := 0; i < len(primes);i++{
		fmt.Printf("valor %v, com indice %v no array\n", primes[i], i )
	}

	//make

	//mySlace := make([]int, 5, 10) criando um array de inteiros, com 5 elementos zerados, com capacidade de 10 antes de realocar
	//Por que usar make?
	//Quando você sabe que o slice vai crescer, e quer evitar várias realocações (por performance).
	//Quando você precisa inicializar maps ou channels, pois var deixa eles como nil.
	//ao estourar o slice, o go vai realocar na memoria um novo slice que comporte o slice antigo + o que extravazou. É bom informar que será uma nova referência, portanto precisa ser novamente referenciado
	//len -> retorna o tamanho
	//cap -> retorna a capacidade



	//quando uma funcao recebe argumetnos assim: func soma(numeros ...int), quer dizer que ela recebe número variável de argumentos inteiros, ou seja, ela recebe quantos argumentos for. Nessa função soma podemos fazer assim: soma(1,2,1), soma(1,1), soma(1,2,3,4,5,6) que naõ vai dar problema. O Go automaticamente cria um slice quand ovocê informa o ...alguma coisa. Se eu coloco soma(1,2,1), o Go cria um slice contendo 1,2,1

	somaTotal := soma(1,2,4,5,6,7)
	fmt.Println(somaTotal)
	//ou podemos chamar assim:
	numeros := []int{1,2,3,4,5,6,7,8,9,10}
	somaTotal2 := soma(numeros...) //se chama assim um slice dentro de uma funcao que aceita varios elementos -> esse ... é chamado de spread operator
	fmt.Println(somaTotal2)



	//slice de slice -> ou matriz de matriz
	
}