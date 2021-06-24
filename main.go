package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return length, uppercase
}


func main()	{
	totalLength, upperName := lenAndUpper("lemma")
	fmt.Println(totalLength, upperName)
}
