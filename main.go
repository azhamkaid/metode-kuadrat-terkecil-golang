package main

import "fmt"

//program berikut adalah program untuk menghitung regresi linear menggunakan metode kuadrat terkecil
func main() {
	var x = [10]float32{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	var y = [10]float32{101, 109, 302, 308, 552, 601, 703, 709, 808, 1003}
	var x2 [10]float32
	var y2 [10]float32
	var xy [10]float32
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

	b := ((SumX * SumY) - (SumXY * 10)) / ((SumX * SumX) - (SumX2 * 2))

	fmt.Println(b)
	a := (SumY - SumX*b) / 10
	fmt.Println(a)

	fmt.Printf("Y=(%vX)+(%v)", a, b)

}

/*
10*a+sumX*b=sumY
sumX*a+Sumx2*b=sumXY

sumX*10*a+sumX*sumX*b=sumY*sumX
sumX*10*a+Sumx2*10*b=sumXY*10

b=((sumX*sumY)-(sumXY*10))/((sumX*SumX)-(sumX2*2))


*/
