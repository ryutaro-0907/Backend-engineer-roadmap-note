package main 

import "fmt"

func main()  {
	var a = "initial"
    fmt.Println(a)

	var b, c = 2, 3
}

func sum(num1, num2) {
	return num1 * num2
	// it should be +
	// logical error
}

func isSame(str1, str2){
	return str1 = str2
	// it should be ==
	// semantic error
}

