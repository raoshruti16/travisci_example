package main
import "fmt"

func main(){
fmt.Println("hello")
fmt.Println(Sq(4))
}

func Sq(x int) int{
	z := 0
	z = x * x
	return z

}