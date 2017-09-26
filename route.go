package main

import (
	"time"
	"fmt"
	"math"
)

func main() {

	route1:=[][]float64{
		{0,0},
		{185,78},
		{212,99},
		{265,146},
		{273,188},
		{279,56},
		{292,243},
		{367,69},
		{392,321},
		{426,324},
		{434,115},
		{437,132},
		{477,364},
		{491,225},
		{499,121},
	}

	//route2:=[][]float64{
	//	{0,0},
	//	{123,35},
	//	{156,63},
	//	{190,355},
	//	{223,255},
	//	{235,144},
	//	{278,35},
	//	{302,366},
	//	{306,465},
	//	{387,234},
	//	{404,134},
	//	{458,234},
	//	{480,134},
	//	{499,211},
	//}
	//
	//route3:=[][]float64{
	//	{0,0},
	//	{142,24},
	//	{156,57},
	//	{178,167},
	//	{234,388},
	//	{278,299},
	//	{333,120},
	//	{387,268},
	//	{390,444},
	//	{397,222},
	//	{406,100},
	//	{453,98},
	//	{487,273},
	//	{499,110},
	//}

	out1:=ret1(route1)
	//out2:=ret2(route2)
	//out3:=ret3(route3)
	v,d:=cal1(out1)
	//cal2(out2)
	//cal3(out3)
	//for v:=range out1{
	//	fmt.Println(v)
	//}

fmt.Println(v,d)
}

func ret1(route1 [][]float64) chan []float64{
	out:=make(chan []float64)
	go func() {
		for _,v:=range route1{
			out<-v
			time.Sleep(time.Millisecond)
		}
		close(out)
	}()

	return out
}

func cal1(in1 chan []float64) (float64,float64){
	var av float64
	sum:=0.0
	go func() {
		old:=<-in1
		new:=<-in1
		var d float64
		counter:=1.0
		vel:=0.0
		for v:=range in1{
			d = math.Sqrt(math.Pow(old[0]-new[0], 2) + math.Pow(old[1]-new[1], 2))
			fmt.Println(d)
			sum+= math.Sqrt(math.Pow(old[0]-new[0], 2) + math.Pow(old[1]-new[1], 2))
			old=new
			new=v
			vel += d/2  //here 2 is the interval after which the coordinates are retreived
			counter++
		}
		av=vel/counter
	}()
	return av,sum

}

//func ret2(route2 [][]float64) chan []float64{
//	out:=make(chan []float64)
//	go func() {
//		for _,v:=range route2{
//			out<-v
//			time.Sleep(time.Millisecond)
//		}
//		close(out)
//	}()
//
//	return out
//}
//
//func ret3(route3 [][]float64) chan []float64{
//	out:=make(chan []float64)
//	go func() {
//		for _,v:=range route3{
//			out<-v
//			time.Sleep(time.Millisecond)
//		}
//		close(out)
//	}()
//
//	return out
//}