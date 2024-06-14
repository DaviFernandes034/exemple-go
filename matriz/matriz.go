package main


import "fmt"



func main(){


	var a [2]string
		
	a[0] = "erhhj"
	a[1] = "refd"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes:= [6]int{12,241,1,42,13,34}
	fmt.Println(primes)


	var s[]int = primes[1:4]

    s = append(s, 6)

	fmt.Println(s)

}