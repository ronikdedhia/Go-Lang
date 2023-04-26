package main
import "fmt"

func main() {
	for i := 0; i < 4; i++{
	fmt.Printf("Welcome!\n")
	}
	
	rvariable:= []string{"A", "B", "C"}
    for i, j:= range rvariable {
       fmt.Println(i, j)
	}
}
