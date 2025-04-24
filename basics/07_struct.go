package main

import "fmt"

type car struct {
	Make string
	Model string 
	Height int 
	Width int 
}

//é possivel trabalhar com estruturas mais complexas:
//chama a struct rodas dentro da struct principal

type car2 struct {
	Make string
	Model string 
	Height int 
	Width int
	Frontwheel Wheel
	Backwheel Wheel
}

type Wheel struct {
	Radius int
	Material string
}

//herannça

type car3 struct { 
	Make string 
	Model string 
}

type truck struct {
	car3  //isso sinaliza que vamos herdar a estutura de car no nosso modelo truck
	badSize int 
}


//metodos com struct

//como chamamos:

func (c car) printCar() string {
	make := c.Make 
	model := c.Model
	result := "carro da marca: " + make + "e do modelo: " + model 
	return result
}


func main(){
	fmt.Println("teste")
	newCar := car {Make: "Teste", Model: "Teste2", Height: 1, Width: 10} //assim que se inicializa uma variavel do tipo "car" segundo a struct que construimos
	fmt.Println(newCar.Height)

	//assim que se inicializa uma variavel do tipo car2 com structs mais complexas
	newCar2 := car2 {Make: "Teste", Model: "Teste2", Height: 1, Width: 10, Frontwheel: Wheel{Radius: 10, Material: "trste"}, Backwheel: Wheel{Radius: 5, Material: "teste"}}
	fmt.Println(newCar2.Frontwheel.Material)

	newTruck := truck {badSize: 10, car3: car3{Make: "Teste", Model: "Teste2"}} //assim se iniciliza structs que herdam caracteristicas. Os dados podem ser acessados com ose fossem pertecnetes a ela

	fmt.Println(newTruck.Make)

	result := newCar.printCar() //assim que se printa uma funcao associada ao objeto
	fmt.Println(result)


}

