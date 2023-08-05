# golang_projects

누가 고랭 좋다고 고랭?

- 23.08.05
  - 이야.. 이거 다만들어써야되네..

##### 참고

- 고랭 문법 튜토리얼 : https://go-tour-ko.appspot.com/flowcontrol/1
- 고랭 웹 튜토리얼 : https://gowebexamples.com/
- 고랭 모듈 튜토리얼 : https://go.dev/doc/tutorial/create-module

## 1.세팅

### 1.1.설치

- https://go.dev/learn/ 접속
- 최신버전 Go 컴파일러 설치

```bash
> go version
go version go1.20.6 darwin/arm64
```

### 1.2.VsCode Extension

- 익스텐션 : go

#### 1.2.1.GO익스텐션 종속성의 문제

- 이 확장 프로그램은 go, gopls, dlv 및 기타 선택적 도구에 따라 달라집니다. 종속성 중 하나라도 누락된 경우 ⚠️ 분석 도구 누락 경고가 표시됩니다. 경고를 클릭하여 종속성을 다운로드하세요.
- 확장프로그램 다운로드시 왼쪽하단에 GO아이콘과 버전이 표기되는데, 이걸 클릭하면 경고창을 확인 가능, 따라서 종속성다운로드가 가능하다.

```bash
Tools environment: GOPATH=~
Installing 1 tool at ~/bin in module mode.
  gopls

Installing golang.org/x/tools/gopls@latest (~/bin/gopls) SUCCEEDED

All tools successfully installed. You are ready to Go. :)
Installing golang.org/x/tools/gopls@latest (~/bin/gopls) SUCCEEDED

All tools successfully installed. You are ready to Go. :)
```

#### 1.2.2.gopls가 작업영역에서 모듈을 찾지못하는 문제

```bash
gopls was not able to find modules in your workspace.
When outside of GOPATH, gopls needs to know which modules you are working on.
You can fix this by opening your workspace to a folder inside a Go module, or
by using a go.work file to specify multiple modules.
See the documentation for more information on setting up your workspace:
https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.go list
```

- 오류는 위와같은데 저 링크들어가면 404뜸. 쫌마니당혹쓰

##### 1.2.2.1.원인

- 여러 디렉토리에 go module 존재하면 생기는 문제

##### 1.2.2.2.조치

- Settings > gopls > Edit in settings.json > gopls json 영역에 아래설정 추가

```json
"experimentalWorkspaceModule": true
```

##### 1.2.2.3.더 나아가서..

- 이 케이스 뿐만아니라, 여러 go파일을 포함하는 디렉토리를 작업영역으로 설정하면, gopls가 작업영역내의 모듈을 찾지못하는 문제가 발생한다.(오류 문구는 위와 동일)
- 이때는 작업영역내의 모듈을 찾을 수 있도록 go.mod 파일을 작성해주면 된다.
- go.mod 파일은 go mod init 명령어로 생성할 수 있다.
- go.mod 파일이 생성되면, gopls가 작업영역내의 모듈을 찾을 수 있게된다.
- go.mod 파일은 go.mod 파일이 있는 디렉토리를 기준으로 모듈을 찾는다.
- 따라서, go.mod 파일을 작성할 때는, 작업영역내의 모듈을 찾을 수 있는 디렉토리를 기준으로 작성해야한다.
- 예를들어, 작업영역이 ~/go/src/github.com/username/project 이고, 모듈이 ~/go/src/github.com/username/project/module 이라면, go.mod 파일은 ~/go/src/github.com/username/project 디렉토리에 생성해야한다.
- go.mod 파일을 생성하고 나면, gopls가 작업영역내의 모듈을 찾을 수 있게된다.
- go.mod 파일을 생성하고 나면, go.mod 파일이 있는 디렉토리를 기준으로 go build, go run 등의 명령어를 사용할 수 있다.

## 1.3.GOPATH

- GO를 설치하면 기본적으로 /Users/username/go에 설치된다.
- 기본적으로 이곳이 GOPATH이며, 이곳에만 코드를 작성할 수 있다.

### 1.3.1.GO의 프로젝트 구조

- GO에는 package.json같은 역할을 하는곳이 없다.
- /users/username/go에 /src/ 폴더를만들고 그 안에 다운로드 받은 코드들을 저장한다.
  - ex) ~/go/src/github.com/username/project 와 같이 말이다.
- 나는 별도의 workspace경로를 가지고있어, ~/go/src/github.com/cogito1016/golang에 지금까지의 프로젝트를 저장하고, 별도의 workspace경로를 심볼릭링크 처리했다.(그런데 VSCode IDE가 재기능을 하지못한다, 따라서 심볼릭 링크만 걸어주고 작업은 ~/go/src/github.com/cogito1016/golang에서 진행한다.)
- 이렇게하니까 기존에 나타나던 'gopls가 작업영역에서 모듈을 찾지못하는 문제'는 사라졌다. "experimentalWorkspaceModule": true 를 삭제했음에도 불구하고 말이다.
- 또한 go.mod를 다 지우고 아래처럼 동작이 가능하다.

```go
import (
	"fmt"

	"github.com/cogito1016/golang/01-module/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys");
    fmt.Println(message)
}
```

## 1.4.main

- go는 main 패키지를 가지고있어야한다.
- main 패키지는 main 함수를 가지고있어야한다.
- main 함수는 인자를 받지않고, 리턴값을 가지지않는다.
- 컴파일러는 main 패키지를 컴파일하고, main 함수를 실행한다.

---

## 2.문법

### 2.1.package 사용

- package는 import하여 가져온다.
- 대문자로 시작하는것들은 Export된다.
- 따라서 적재한 패키지의 함수를 사용하려면 대문자의것들을 참조해야한다.
- 같은 디렉토리 내의 go파일들은 같은 패키지를 사용해야한다. (아닐 시 아래 오류발생)

```bash
found packages print_problem (basic_print.go) and print_problemㅁ (boj_1002.go)
```

### 2.2. 함수

- func 함수명(인자명 인자타입) 리턴타입 { }
- 연속된 인자가 타입이 같은경우 마지막에만 표기할 수 있다. (x, y int)
- 한 함수는 복수개의 결과를 반환할 수 있다.
  - func add(x, y int) (int, int) { return x + y, x - y }
- 복수개의 결과를 가져올 때에는 변수를 선언하여 받아야한다.
  - sum, sub := add(42, 13)
  - 그러나 변수를 선언하지 않고 결과를 버리고 싶다면, \_ 를 사용한다.
    ```go
    sum, _ := add(42, 13)
    ```

### 2.3. 변수

#### 2.3.1. 타입의 종류

- bool
- string
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte // uint8의 다른 이름
- rune // int32의 다른 이름
- float32 float64
- complex64 complex128

#### 2.3.2. 변수선언

- var 변수명 변수타입
- var 변수명 = 값 (이때, 변수타입 생략가능 초깃값의 타입을 변수타입으로 사용)

```go
var a = "Hello?"
var b = 1
var c = 0.23
//or var a,b,c = "Hello?",1,0.23;
fmt.Printf("%T, %T, %T\n",a,b,c); //string, int, float64
```

- var과 타입을 생략한 := 를 사용하여 선언과 동시에 초기화 가능
  - 함수내에서만 사용가능
  - 함수 밖에서는 var로 선언해야함

```go
a := 1
fmt.Printf("%T, %d",a,a); //int, 1
```

- 초기값이 없이 선언한 변수는 Zero Value를 갖는다.
  - 숫자형 : 0
  - bool : false
  - string : ""

#### 2.3.3. 형변환

- 명시적으로 변환해야한다.

```go
var i int = 42;
fmt.Printf("%T, %d\n",i,i); //int, 42
var f float32 = float32(i);
fmt.Printf("%T, %f\n",f,f); //float32, 42.000000
var u uint = uint(f);
fmt.Printf("%T, %d\n",u,u); //uint, 42
```

### 2.4. 상수

- 상수는 const 키워드를 사용하여 선언한다.
- 상수는 문자, 문자열, boolean, 숫자타입 중 하나의 값만 가질 수 있다.
- 상수는 := 를 사용하여 선언할 수 없다.
- 상수는 대문자로 시작하여 Export할 수 있다.

```go
const PI = 3.14;
fmt.Printf("%T, %f\n",PI,PI);//float64, 3.140000
```

### 2.5. 반복문

```go

func forNumbers01(numbers ...int) int{
	for number := range numbers{
		fmt.Println(number);
	}

	return 0;
}

func forNumbers02(numbers ...int) int{
	for idx,value := range numbers{
		fmt.Println(idx,value);
	}
	return 0;
}

func forNumbers03(numbers ...int) int{
	for i := 0; i< len(numbers); i++{
		fmt.Println(i,numbers[i]);
	}

	return 0;
}
```

### 2.6. 조건문

```go
func canDrink(age int) bool{
	if age > 18{
		return true;
	}

	return false;
}

func canDrink2(age int) bool{
	if koreanAge := age+2; koreanAge>18{
		return true;
	}

	return false;
}
```

### 2.7.포인터

- C언어의 \*와 &와 같은 역할을 한다.
- &는 주소참조연산자, \*는 포인터연산자이다.

```go

	a = 10;
	c := &a; // b = &a 할 시, cannot use &a (value of type *int) as int value in assignment
	fmt.Println(a,c); //10 0x140000140b8
	*c = 20;
	fmt.Println(a,c); //20 0x140000140b8
	fmt.Println(&a, &b); //0x140000140b8 0x140000140c0
	fmt.Println(a, b); //20 10
	fmt.Println(a, b); //20 10

	//a의 type을 알고 싶을때
	fmt.Printf("%T\n",a); //int
	fmt.Printf("%T\n",&a); //*int

```

### 2.8.배열

- basic array

```go
names := []string{"kim","lee"};
// names[2] = "park"; 오류발생
names = append(names,"park"); //append는 원본을변경하지않고 새로운배열을 반환한다.

for idx,name := range names{
	fmt.Println(idx,name);
}
```

- slice

```go
names := []string{"kim","lee"};
// names[2] = "park"; 오류발생
names = append(names,"park"); //append는 원본을변경하지않고 새로운배열을 반환한다.

for idx,name := range names{
	fmt.Println(idx,name);
}
```

- array에는 :(콜론)개념이 있는데,이는 배열의 범위를 지정해주는 역할을 한다

```go
for i := 0; i < 10; i++ {
	arr = append(arr,i+1);
}

fmt.Println(arr);//[1 2 3 4 5 6 7 8 9 10]
fmt.Println(arr[0:5]);//[1 2 3 4 5]
fmt.Println(arr[:3]);//[1 2 3]
fmt.Println(arr[2:3]);//[3]
fmt.Println(arr[3:]);//[4 5 6 7 8 9 10]
```

- append는 중요한개념이다.

  - append(대상배열, 추가할값1, 추가할값2..)
  - :(콜론)을 조합하여서 쓸 수 있다

  ```go
  arr = append(arr[:2],arr[3:]...);
  ```

### 2.9.Map

```go
jayden := map[string]string{"name":"jayden","age":"18"}
fmt.Println(jayden);//map[age:18 name:jayden]

for key,value := range jayden{
fmt.Println(key, value);
/**
name jayden
age 18
**/
}
```

### 2.10 Struct

```go
type person struct {//구조체
	//생성자는없음
	name string
	age int
	favFood []string
}

func Run(){
	fmt.Println("Struct Test");
	favFood := []string{"apple"};
	jayden := person{name:"jayden",age:18,favFood:favFood}
	//jayden := person{"jayden",18,favFood} 이렇게 사용해도 됨
	fmt.Println(jayden);//{jayden 18 [apple]}
	fmt.Println(jayden.favFood);//[apple]
}
```

## 3.모듈

### 3.1.에러

#### 3.1.1. Broken Import 에러

##### 3.1.1.1. 배경

- greetings/ 에 go.mod와 greetings.go가 존재
- hello/ 에 go.mod와 hello.go가 존재
- hello.go에서 greetings를 import하려고할때 아래의 에러가 발생한다.

```bash
could not import example.com/greetings (cannot find package "example.com/greetings" in GOROOT or GOPATH)compilerBrokenImport
```

- 그러나 정상동작은 함 (..)

##### 3.1.1.1. 조치

- 일단 기본적으로 위의 구조가 잘못되었었다.
- go.mod는 디렉토리구조의 최상단에 위치되어야하며, 위의 케이스처럼 greetings와 hello에 각각 위치되면안된다.
- 따라서 아래와 같이 구조를 변경해주었다.

```bash
golang
├── go.mod
├── greetings
│   └── greetings.go
└── hello
	└── hello.go
```

## 4. 웹
