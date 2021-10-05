package main

import "fmt"

func main(){

	var ordinates []float64
	ordinates = append(ordinates, 32, 34, 56, 77, 12, 5, 88)
	fmt.Println(ordinates)
	magn := Transform(ordinates, 3.5, magnify)
	fmt.Println(magn)

	magn = Transform(ordinates, 0.19, rotate)
	fmt.Println(magn)

	magn = Transform(ordinates, 6.43, reflect)
	fmt.Println(magn)

	fmt.Println(ordinates)
}


func Transform(list []float64, factor float64, transFunction func(float64,  float64) float64) []float64{
	var ordinates []float64
	for index, _ := range list{
		ordinates = append(ordinates, transFunction( list[index], factor))
	}
	return ordinates
}


func TransformPtr(list *[]float64, factor float64, transFunction func(float64,  float64) float64){
	for index, _ := range *list{
		(*list)[index] = transFunction( (*list)[index], factor)
	}
}

