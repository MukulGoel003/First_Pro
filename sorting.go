package main

import (
	"sort"
	"fmt"
)

type people []string
func (p people) Swap(i,j int){p[i],p[j]=p[j],p[i]}
func (p people) Len() int{return len(p)}
func (p people) Less(t,j int) bool{return p[t]<p[j]}

func main() {

	arr:=people{"john","ron","harry","alex","james"}
	a:=[]string{"abc","bhd","aaa","ujhg"}
	sort.Sort(arr)
	sort.Strings(a)
	fmt.Println(arr)
fmt.Println(a)
	sl:=[]int{3,5,2,7,6,90}
	sort.Ints(sl)
	fmt.Println(sl)
	sort.Sort(sort.Reverse(arr))
	fmt.Println(sort.Reverse(arr))
	fmt.Println(arr)
}
