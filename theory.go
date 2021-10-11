package main

import (
	"fmt"
	"strings"
)

/* Functions */

//naked return : return값이 되는 변수를 굳이 명시하지 않아도 됨
//어떤 타입의 변수를 return 할건지는 여기서 정의를 했기 때문↓↓↓↓↓↓
func lenAndUpper(name string) (length int, uppercase string) {
    length = len(name)
	uppercase = strings.ToUpper(name)
	return
	/* 따라서 여기서 return length, uppercase 라고 작성하지 않아도 
	알아서 length와 uppercase가 업데이트 된다는 뜻
	물론 무엇인가는 return 되어야함 */
}

//defer
func lenAndUpperDefer(name string) (length int, uppercase string) {
	//defer라고 적으면 return이 된 후에(function이 끝나고 나서) 실행됨
	defer fmt.Println("I'm done")
	//활용용도:특정 API를 불러온다거나, 이미지를 불러온다거나, 창을 닫는다던가...
    length = len(name)
	uppercase = strings.ToUpper(name)
	return
}


/* for, range, ...args */

//Go에서 반복문은 for만 이용가능함
//forEach, for in 그런거 없음
func superAdd(numbers ...int) int {
	total := 0
	/* range : array에 loop를 적용할 수 있도록 해줌
	range는 index를 주기 때문에 0부터 시작함!!
	index가 필요 없는 경우는 다음과 같이 언더바로 기술 */

	// for _, number := range ...
	for index, number := range numbers {
		fmt.Println(index, number)
		//0, 1
		//1, 2
		//2, 3 ...
		total += number
	}
	// 요렇게도 가능함
	for  i:=0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
		//1,2,3,4...
	}
	return total
}

/* if문 */


func canIdrink(age int) bool {
// case 1. if 사용
// Go에서는 조건(condition)을 체크하기 전에 변수를 만들 수 있음
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} 
	// switch문에서도 if문과 같이 변수를 만들고 사용할 수 있다
	switch koreanAge := age + 2; koreanAge {
		case 10:
			return false
		case 18:
			return true
	}

	return true
}