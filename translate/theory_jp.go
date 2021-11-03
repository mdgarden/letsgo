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
// case 1. if 使用
// Goでは条件文(condition)をチェックする前に変数を作ることが可能
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} 
	// switch分でもif分と同じく変数を作って使用できる
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
	// aの値：2ではなくaのメモリーアドレスをコピーする
	// &はアドレスを意味
	b := &a
	fmt.Println(&a, b)
	//*は観察（？）の意味
	fmt.Println(*b)
	// 出力値：2＝bはaのメモリアドレス値を持っていて、そのアドレス値を見せる

	*b = 20 // aの値を変更
}

func arrAndSlice() {
	names := [5]string{"john", "bob", "james"}
	names [3] = "billy"
	names [4] = "jane"
	names [6] = "emma" // エラー理由：namesという配列は[5]まで定義されているので、それ以上で入れるのはだめ

	// 長さ関係なく配列を使いたい時はslice使用 / length値を空にすればOK
	names_second := []string{"john", "bob", "james"}

	/* appendは二つの引数が必要。一つはappendするslice, もう一つはappendする値
	例）append(names, "harry")
	appendはnamesを修正するのではなく、新しい値が追加された新しいsliceをreturnする
	namesというslice自体を修正するためには以下の書き方で記述する。
	*/
	names_second = append(names_second, "jane")
	fmt.Println(names_second)
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

