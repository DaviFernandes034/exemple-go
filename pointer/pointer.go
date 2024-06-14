package main


import "fmt"


func main(){


	i, j := 42, 2701

	p:= &i // ponteiro para i
	fmt.Println(&p) //ler o valor de i atraves do ponteiro p

	*p = 21 //definindo um nono valor para i atraves do ponteiro i

	fmt.Println(&i) //novo valor

	p = &j //ponteiro para J

	*p = *p/37 //dividndo o valor de J por 37, usando o ponteiro

	fmt.Println(j) //novo valor de j






}