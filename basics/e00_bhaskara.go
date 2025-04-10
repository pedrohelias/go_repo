package main

import(
	"fmt"
	"math"
)

func main(){

	a := 1.0
	b := -2.0
	c := -3.0
	exp := (b*b)-4*a*c
	root1 := ((-b)+ math.Sqrt(exp))/(2*a)
	root2 := ((-b)- math.Sqrt(exp))/(2*a)

	fmt.Println(exp)
	fmt.Println(root1)
	fmt.Println(root2)



}