package main
import "fmt"
func hello(x int,y int,a int,b int) int{
	
	min:=x
	if y<min {
		min=y
	}
	if a<min {
		min=a
	}
	if b<min {
		min=b
	}
	fmt.Println(min)
	return min
}
func main() {
	hello(4,5,6,7)
	
	//fmt.Println(min)
}