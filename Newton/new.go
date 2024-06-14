package main



import(

	"fmt"
	"math"
)


func sqrt(x float64) float64{

	z:= 1.0
	for i:= 0; i< 10; i++{
		z = z - ( z* z - x) / (2 * z) //metado de Newton

		fmt.Printf("interações: %d: z = %f\n", i+1, z)
	}

	return z
}


func main(){


	
		aprox:= sqrt(121.0)
		esp:= math.Sqrt(121.0)
		fmt.Printf("aproximadamente: %f\n", aprox)
		fmt.Printf("esperado: %f \n", esp)

		fmt.Printf("diferença: %f \n", math.Abs(aprox - esp))
	
}


