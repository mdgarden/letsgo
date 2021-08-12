package main

import "fmt"

func canIDrink(age int) bool {

	switch koreanAge := age + 2; koreanAge {
		case 10:
			return false
		case 18:
			return true
	}

	return true
}


func main()	{
	fmt.Println(canIDrink(16))
}
