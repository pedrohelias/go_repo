package main

import "fmt"


func main(){

	//maps são similiares a objetos no javascript, dicionários em python... Sempre vamos associar chaves e valores. O valor zero de um mapa é nil
	//construção

	mapa := make(map[string]int) //isso quer dizer que a chave vai ser uma string, e o valor um int
	
	mapa["Pedro"] = 29
	mapa["Anne"] = 27
	mapa["teste"] = 10


	mapa = map[string]int{ //isso reestrutura todo um mapa caso ele ja existe. Entao todos os argumentos precisam ser repassados. Caso o mapa nao exista, esse comendo serve para adicionar argumentos.
		"teste2": 40,
	}

	//ou

	fmt.Println(len(mapa)) //retorna a quantidade de pares chave-valor tem no mapa
	fmt.Println(mapa) //printa todo o mapa
	

}