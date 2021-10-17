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

func pointers()	{
	a := 2
	//a의 값인 2가 아닌 a의 메모리 주소를 복사하겠다는 뜻
	//&는 주소
	b := &a
	fmt.Println(&a, b)
	//*는 살펴본다는 뜻
	fmt.Println(*b)
	// 출력값:2 = b는 a의 메모리주소 값을 가지고 있고 그 주소값의 할당치를 보여줌
	*b = 20
	// a의 값을 변경시킴
}

func arrAndSlice() {
	names := [5]string{"wraith", "wattson", "crypto"}
	names [3] = "pathfinder"
	names [4] = "pathfinder"
	/* 여기서 에러 발생 : names라는 배열은 5칸으로 정의되어있는데 6번째 자리에
	더 넣을려고 하기 때문에 
	names [5] = "pathfinder"
	길이가 없는 배열을 쓰고 싶을 때 slice사용 / length만 비워두면 됨 */
	legends := []string{"wraith", "wattson", "crypto"}

	/* 근데 여기서 93번째 줄같이 쓰면 에러남.
	추가하려면 아래와 같이 작성
	append는 두개의 인자를 요구함. 하나는 slice, 두번째는 값
	append(legends, "pathfinder")
	append는 legends 슬라이스 자체를 수정해주는 것이 아니라 새 값이 추가된 새 슬라이스를 return함
	따라서 legends라는 slice 자체를 수정하려면 다음과 같이 작성*/

	legends = append(legends, "pathfinder")
	fmt.Println(legends)
}

func Maps() {
	// map에는 key와 value가 있음
	wraith := map[string]string{"name":"Renee", "age":"32"}
	fmt.Println(wraith)
	for key, value := range wraith {
		fmt.Println(key)
		fmt.Println(value)
	}
	// TODO: map에 요소 추가하는 법
}

// Structs : Object와 비슷하면서 map보다 유연한 것이 특징