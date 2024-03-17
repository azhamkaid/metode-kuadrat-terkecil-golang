package main

import "fmt"

//program berikut adalah program untuk menghitung regresi linear menggunakan metode kuadrat terkecil
func main() {
	//insert value of x from 0 to 100 and y from 0 to 1000 for 10 times
	var x = [10]float32{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	var y = [10]float32{71, 119, 282, 328, 552, 631, 713, 832, 898, 1063}
	var x2 [10]float32
	var y2 [10]float32
	var xy [10]float32
	//initiate new value for x2,y2,and xy
	for i := 0; i < 10; i++ {
		x2[i] = x[i] * x[i]
		y2[i] = y[i] * y[i]
		xy[i] = x[i] * y[i]
	}
	fmt.Println("x^2", x2, "\n y^2", y2, "\n x*y", xy)
	var SumX float32
	var SumY float32
	var SumX2 float32
	var SumY2 float32
	var SumXY float32
	SumX = 0
	SumY = 0
	SumX2 = 0
	SumY2 = 0
	SumXY = 0
	//search value for sum x,sum y,sumx2, and sumxy
	for i := 0; i < 10; i++ {
		SumX += x[i]
		SumY += y[i]
		SumX2 += x2[i]
		SumY2 += y2[i]
		SumXY += xy[i]
	}

	MeanX := SumX / float32(len(x))
	MeanY := SumY / float32(len(y))

	fmt.Println("SumX", SumX)
	fmt.Println("SumY", SumY)
	fmt.Println("SumX2", SumX2)
	fmt.Println("SumY2", SumY2)
	fmt.Println("MeanX", MeanX)
	fmt.Println("MeanY", MeanY)
	//calculate a value of b
	b := ((SumX * SumY) - (SumXY * 10)) / ((SumX * SumX) - (SumX2 * 10))

	fmt.Println(b)
	//calculate value of a
	a := (SumY - SumX*b) / 10
	fmt.Println(a)

	fmt.Printf("Y=(%v)X+(%v)", a, b)

}

/*
equation to calculate linear euations using Least squares approximation metods
n*a+sumX*b=sumY
sumX*a+Sumx2*b=sumXY

sumX*10*a+sumX*sumX*b=sumY*sumX
sumX*10*a+Sumx2*n*b=sumXY*n

b=((sumX*sumY)-(sumXY*n))/((sumX*SumX)-(sumX2*n))
a=
*/
