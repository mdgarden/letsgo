# 쉽고 빠른 Go 시작하기

## 1장 THEORY

### 1.1 Packages and Imports

- Node.js나 Python과는 달리 모듈(function)을 export 하려면 별도로 작성을 해주는 것이 아니라, function이 대문자로 시작하면 됨
- fmt.Println이 대문자인 이유인 fmt라는 폴더에서 Println이라는 함수를 불러오고 있기 때문.

### 1.2 Variables and Constants

- var name string = "lemma" 이거랑
- name := "lemma" 이건 같은거 / 이렇게 코드를 축약해서 사용하면 첫번째 값을 기준으로 타입을 알아서 찾아줌, 단 타입을 임의로 변경할 수 없음.
- 축약코드는 func밖에서 작동하지 않음, 변수에만 적용 가능함.