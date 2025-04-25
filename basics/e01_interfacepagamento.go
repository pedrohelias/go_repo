package main

import "fmt"

type Pagamento interface{
	pagar(msg string)
}

//criado cartao
type cartao struct {
	nome string
	numero int
	codigo string
}

func (c cartao) pagar(msg string){
	fmt.Printf("O Cartão com o nome %s, numero %v e codigo %s acabou de realizar o pagamento. %s \n", c.nome, c.numero, c.codigo, msg)
}


//criando boleto
type boleto struct {

}

func (b boleto) pagar(msg string){
	fmt.Printf("O pagamento foi realizado utilizando boleto. %s \n", msg)
}


//criando pix
type pix struct {
	chave string
}

func (p pix) pagar(msg string){
	fmt.Printf("O pagamento foi realizado utilizando pix, com a chave: %s. %s \n", p.chave, msg)
}

func payment(p Pagamento){
	p.pagar("Pagamento realizado")
}




func main(){
	cartao := cartao{nome: "Pedro Helias Carlos", numero: 123456, codigo: "123"}
	boleto := boleto{}
	pix := pix{chave: "12345"}

	payment(cartao)
	payment(boleto)
	payment(pix)
}