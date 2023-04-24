package main
import "fmt"
		
func main() {
	var X uint8 = 225
	fmt.Println(X, X-3)
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)

	a := 20.45
    b := 34.89
    c := b-a
    fmt.Printf("Result is: %f", c)
    fmt.Printf("\nThe type of c is : %T", c) 

	var m complex128 = complex(6, 2)
	var n complex64 = complex(9, 2)
	fmt.Println("\n",m)
	fmt.Println(n)
	fmt.Printf("The type of a is %T and "+"the type of b is %T", m, n)

	str1 := "GeeksforGeeks"
	str2:= "geeksForgeeks"
	str3:= "GeeksforGeeks"
	result1:= str1 == str2
	result2:= str1 == str3
	fmt.Println("\n", result1)
	fmt.Println( result2)
	fmt.Printf("The type of result1 is %T and "+"the type of result2 is %T",result1, result2)
	fmt.Printf("Length of the string is:%d",len(str1))
	fmt.Printf("\nType of str is: %T", str1)
}
