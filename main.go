package main

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
	superAdd(1,2,3,4,5,6)
}
