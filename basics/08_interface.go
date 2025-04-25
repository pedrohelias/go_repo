package main

//Interface é um tipo que define um conjunto de métodos. Qualquer tipo que implementar todos os métodos da interface automaticamente "satisfaz" essa interface — sem precisar declarar isso explicitamente.

//vantagem de usar. Suponto que voce tenha 2 fluxos de serviço -> Enviar dados via fax e via correios. Você escreve uma função para fax. e para correios. Ou vocẽ apenas alteraria a função de fax e duplicaria. Em larga escala isso pode ser um problema. A interface veio para simplificar a reutilização das estruturas codificadas.



import "fmt"

// Interface
type Flow interface {
	send(msg string)
}

// Implementação fax
type fax struct{}

func (f fax) send(msg string) {
	fmt.Println("mandando via fax:", msg)
}

// Implementação correios
type correios struct{}

func (c correios) send(msg string) {
	fmt.Println("mandando via correio:", msg)
}

// Função que usa a interface
func realizarEnvio(fl Flow) {
	fl.send("Olá, mundo!")
}

func main() {
	fluxo := fax{}
	realizarEnvio(fluxo)
}
//	A interface não diz como a função funciona. Ela só diz:
// "qualquer tipo que quiser ser um Flow precisa ter um método send(string). Por isso as outras func tem o send, pois elas o definem"	
//	Quem define o que essa função faz é o tipo que implementa a interface — fax e correios.}