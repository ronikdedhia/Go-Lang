package main
import "fmt"
func main() {
	p:= 34
	q:= 20
	result1:= p + q
	fmt.Printf("Result of p + q = %d", result1)
	result2:= p - q
	fmt.Printf("\nResult of p - q = %d", result2)
	result3:= p * q
	fmt.Printf("\nResult of p * q = %d", result3)
	result4:= p / q
	fmt.Printf("\nResult of p / q = %d", result4)
	result5:= p % q
	fmt.Printf("\nResult of p %% q = %d", result5)

	result1:= p == q
	fmt.Println(result1)
	result2:= p != q
	fmt.Println(result2)
	result3:= p < q
	fmt.Println(result3)
	result4:= p > q
	fmt.Println(result4)
	result5:= p >= q
	fmt.Println(result5)
	result6:= p <= q
	fmt.Println(result6)

	if(p!=q && p<=q){
		fmt.Println("True")
	}
	if(p!=q || p<=q){
		fmt.Println("True")
	}
	if(!(p==q)){
		fmt.Println("True")
	}

   result1:= p & q
   fmt.Printf("Result of p & q = %d", result1)     
   result2:= p | q
   fmt.Printf("\nResult of p | q = %d", result2)     
   result3:= p ^ q
   fmt.Printf("\nResult of p ^ q = %d", result3)
   result4:= p << 1
   fmt.Printf("\nResult of p << 1 = %d", result4)
   result5:= p >> 1
   fmt.Printf("\nResult of p >> 1 = %d", result5)
   result6:= p &^ q
   fmt.Printf("\nResult of p &^ q = %d", result6)
     
}
