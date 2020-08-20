package utils

import "fmt"

// PrtVar print var
func PrtVar(){

	var a = 1
	var b , c int = 3 ,2 
	fmt.Println(a);
	for i := 0; i< 2; i++ {
		fmt.Println("loo")
	}  
	fmt.Println(b,c);
}