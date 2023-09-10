# golang_projects

누가 고랭 좋다고 고랭?

- 23.08.05
  - 이야.. 이거 다만들어써야되네..
  - 코딩테스트의 메모내용
  - https://github.com/cogito1016/golang_projects/tree/main/02-coding-problem

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

- make를 사용하여 쉽게 배열을 생성할 수 있다.

  - make(타입, 길이, 용량)

  ```go
  arr := make([]int, 5, 10);
  fmt.Println(arr);//[0 0 0 0 0]
  fmt.Println(len(arr));//5
  fmt.Println(cap(arr));//10
  ```

  - 2차원배열 생성

  ```go
  arr := make([][]int,M);
  for i := 0; i < N; i++ {
  	arr[i] = make([]int,N);
  }
  ```

### 2.9.Map
- 맵에는 순서가 보장되지 않는다.

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

#### 2.9.1. 키의 존재유무 판단

#### 2.9.2. 맵변수와 nil의 관계

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

### 2.11 내장함수 len()

- 문자열의 크기는 len()으로 구할 수 있다.
- 문자열 뿐만아니라 배열, 슬라이스, 맵 등의 크기도 len()으로 구할 수 있다.

### 2.12 Type Method 타입메서드
- 타입메서드는 type키워드로 정의한 타입에 종속된 메서드를 의미한다.
- 아래와같이 정의 할 수 있다.
- 타입메서드를 통해 Go언어에서도 객체지향 프로그래밍이 가능하다.
```go
type ar2x2 [2][2]int
```

### 2.13 Interface 인터페이스
- 인터페이스는 구현해야할 메서드의 집합이다
- 어떤 데이터타입이 인터페이스를 만족하려면 해당 인터페이스가 요구하는 타입메서드를 구현해야한다
- sort.Sort는 sort.Interface라는 인터페이스를 만족하는 파라미터를 넘겨야하는데,
```go
func Sort(data Interface) {
	n := data.Len()
	if n <= 1 {
		return
	}
	limit := bits.Len(uint(n))
	pdqsort(data, 0, n, limit)
}
```
- sort.Interface를 예로들어보자면, 
  - Len, Less, Swap을 구현해야한다. 
  - 그렇지않으면서 sort.Sort()메서드를 사용하려고 한다면? 아래메세지를 받게된다.
```go
Cannot use '(data)' (type []S2) as the type Interface Type does not implement 'Interface' as some methods are missing: Len() int Less(i int, j int) bool Swap(i int, j int)
```
- (대충 인터페이스가 요구하는 것들 구현하라는 내용)
   
 
- 그 외에도, 빈 인터페이스라는 개념이 존재한다.
- 빈 인터페이스는 모든 타입이 구현하고있는 인터페이스이다.
- 빈 인터페이스를 만족하는 파라미터를 받는 함수가있다고했을 때, 이 함수의 내용은 '모든 타입을 만족하는' 구문이여야 오류가발생하지 않을것이다.
```go
func Print(a interface{}){
	fmt.Println(a)
}
```
- 따라서 경우에 따라, 매우매우 위험한 구문이 될 수 있다.

#### 2.13.1. 타입단언(type assertion)과 타입스위치(type switch)
- 빈 인터페이스를 사용할 때 매우 위험하다는 이야기를 위에서 했었다.
- 각 타입마다 동일한 로직을 사용하기 때문인데,
- 이는 타입단언과 타입스위치로 할 수 있다.
- 요약하자면, 빈 인터페이스로 넘겨온 데이터의 타입을 체크하여 타입 별 로직을 수행하는것이다. 

### 2.14 고루틴과 채널 Go-Rutine & Channel
#### 2.14.1. 고루틴원리
- OS의 프로세스 안에
- 스레드가있고
- 그 안에 고루틴이있다
- 고 루틴은 고 런타임에 돌아가는 고 스케줄러에 의해 관리된다
- 서로 다른 고 루틴들은 채널을 이용하여 데이터를 공유할 수 있다.
- 1번고루틴의 처리 -> 채널을통한 2번 고루틴으로 전달 -> 2번고루틴처리..등의 일련의 과정을 파이프라인이라고한다.
- 파이프라인은 정확히, 고루틴과 채널을 연결한 가상메서드이다(1번고루틴의 결과를 2번고루틴의 입력으로 만드는)

#### 2.14.2. 고루틴
- 고 루틴은 go 키워드를 사용하여 구현한다
```go
go func(x int){
fmt.Println(10)
}
```
- 고 프로그램은 고 루틴을 기다린 후 종료하지않는다. 대기하기위해선 아래의 구문이 필요하다
```go
var waitGroup sync.WaitGroup

waitGroup.Add(1)
go func(x int){
	defer waitGroup.Done()
	fmt.Println(10)
}
waitGroup.Wait()
```
#### 2.14.3. 채널
- 고루틴 사이의 데이터를 전달하기위해 채널을 사용한다.
- 채널은 타입을 갖고있다.
- 아래는 버퍼채널을 생성하는 방법이다.
- 버퍼채널은 버퍼가 채워지면 채널을 닫을 수 있고, 고 루틴은 다음 구문의 실행을 계속하고 반환할 수 있다.
- 버퍼를 사용하지 않는 채널은 누군가가 값을 가져갈때까지 실행을 멈춘다
```go
channel := make(chan int,1) //1의 크기를 갖는 버퍼채널
```
- 닫힌 채널에서 값을 가져가면 해당 채널 타입의 제로벨류를 반환한다.
- 이 뿐만아니라 다양한 문제를 방지하고자, 매개변수로 채널을 전달 할 때 읽기/쓰기 전용인지를 명시할 수 있다.
```go
func f2(out <-chan int, in chan<- int){
	x := <-out;
	in <- x
	return
}
```
#### 2.14.4. 경쟁상태 감지기
- go run -race 를 사용하면 경쟁상태 감지기가 작동한 결과가 출력된다.
- 아래는 예시인데, 채널을 닫은것과 썼을때의 순서를 모르기에 경쟁상태가 발생했다.
```go
==================
WARNING: DATA RACE
Write at 0x00c000104010 by main goroutine:
github.com/cogito1016/golang/00-grammar/go-rutine.RunChannelBasic()
/00-grammar/go-rutine/channel.go:39 +0x3ac

//39라인에 채널을 닫았다.

Previous read at 0x00c000104010 by goroutine 8:
github.com/cogito1016/golang/00-grammar/go-rutine.printer()
/00-grammar/go-rutine/channel.go:55 +0x34

//55라인에 채널에 쓰기를했다

Goroutine 8 (running) created at:
github.com/cogito1016/golang/00-grammar/go-rutine.RunChannelBasic()
/00-grammar/go-rutine/channel.go:27 +0x20c
```
- 데이터를 쓰는 작업을 여러 고루틴이아니라 하나의 고루틴으로 처리하고
- 그 고루틴 함수에서 채널을 닫는것으로 마무리한다면
- 한 고루틴에서 모든 데이터를 쓰고 채널까지닫았으니 
  - 모든일이 순차적으로 일어나므로
  - 경쟁상태도 발생하지않고 range또한 스스로 종료된다.


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

## 4. 디버그

- VScode를 이용한 디버그
	- 못해먹겠다 그냥 Jetbrains GO land살까

1. 왼쪽 Run and Debug 버튼 클릭
2. Add Configuration을 통해 launch.json 생성
3. Delve디버거를 설치(VS코드에서 go를 디버그할때 사용, 없으면 하단 오류 발생, 설치해주자)

```bash
Couldn't start dlv dap:
Error: Cannot find Delve debugger (dlv dap)
```

```bash
go get github.com/go-delve/delve/cmd/dlv
```

4.
