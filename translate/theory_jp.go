package translate

import (
	"fmt"
	"strings"
)

/* Functions */

//naked return : returnする変数を表記しなくてもいい
//どういうタイプの変数をreturnするかはすでに定義されている=(length int, uppercase string)
func lenAndUpper(name string) (length int, uppercase string) {
    length = len(name)
	uppercase = strings.ToUpper(name)
	return // ここで return length, uppercaseと書かなくてもいい
	/* 変数のタイプ表記を省略しただけ。return値は必須。*/
}

//defer
func lenAndUpperDefer(name string) (length int, uppercase string) {
	//deferはreturnの後に(functionが終わってから)実行される
	defer fmt.Println("I'm done")
	//活用例：特定APIを呼ぶ、ブラウザーを閉じるなどなど
    length = len(name)
	uppercase = strings.ToUpper(name)
	return
}


/* for, range, ...args */

// Goでループはforだけ
// forEach, for in などはない
func superAdd(numbers ...int) int {
	total := 0

	/* range : arrayへループを適用できるようにする
	rangeはindexをくれるので0から始まる
	indexがいらない場合、次のように_で記述
	*/

	// for _, number := range ...
	for index, number := range numbers {
		fmt.Println(index, number)
		//0, 1
		//1, 2
		//2, 3 ...
		total += number
	}
	// この方法でも記述可能
	for  i:=0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
		//1,2,3,4...
	}
	return total
}

/* if文 */


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
	names [4] = "lifeline"
	/* 여기서 에러 발생 : names라는 배열은 5칸으로 정의되어있는데 6번째 자리에
	더 넣을려고 하기 때문에 
	names [6] = "pathfinder"
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
	wraith := map[string]string{"name":"Renee Blasey", "age":"32"}
	fmt.Println(wraith)
	for key, value := range wraith {
		fmt.Println(key)
		fmt.Println(value)
	}
	// TODO: map에 요소 추가하는 법
}

/*
Structs : Object와 비슷하면서 map보다 유연한 것
JavaScript의 오브젝트 같은 느낌으로 쓰고 싶을 때
원래는 파일 가장 위-import 아래에 작성하는 것이 일반적
Go에는 constructor method가 없음
파이썬의 __init__이나 JavaScript의 constructor()같은거
*/

type legend struct {
	name string
	age int
	affiliations []string
}

// 실제 사용
func inputStruct(){
	// case 1 : struct 순서대로 입력 = key/value를 동시에 확인하기가 어렵고 코드가 깔끔해보이지않음
	affiliations := []string{"IMC", "The Syndicate"}
	wraith := legend{"Renee Blasey", 32,  affiliations}
	fmt.Println(wraith)

	// case 2
	wattson := legend{name:"Natalie Paquette", age:22, affiliations: affiliations}
	fmt.Println(wattson.name)

	// field:value로 작성 시작했으면 그 내용은 전부 field:value로 맞춰야함
}

