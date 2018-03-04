package main
import "fmt"
func main ()  {
	for i := 0; i <= 1000; i++ {
		go fmt.Println(i)
	}
	fmt.Println("Hello World " )
}