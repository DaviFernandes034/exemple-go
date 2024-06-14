package main


import "fmt"



type person struct{

	name string
	age int
	count int
}

func (p *person) increment(){

	p.count ++

}


func NewPerson(name string, age int) person{


	p1:= person{}
	pointer:= &p1

	pointer.name = name
	pointer.age = age
	p1.increment()

	return p1


	
}

func main(){

	p1:= NewPerson("rada", 21)
	p2:= NewPerson("davi", 14)
	

	fmt.Println(p1, p2)
}