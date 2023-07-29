# golang_projects

누가 고랭 좋다고 고랭?

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

### 1.3.gopls가 작업영역에서 모듈을 찾지못하는 문제

```bash
gopls was not able to find modules in your workspace.
When outside of GOPATH, gopls needs to know which modules you are working on.
You can fix this by opening your workspace to a folder inside a Go module, or
by using a go.work file to specify multiple modules.
See the documentation for more information on setting up your workspace:
https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.go list
```

- 오류는 위와같은데 저 링크들어가면 404뜸. 쫌마니당혹쓰

#### 1.3.1.원인

- 여러 디렉토리에 go module 존재하면 생기는 문제

#### 1.3.2.조치

- Settings > gopls > Edit in settings.json > gopls json 영역에 아래설정 추가

```json
"experimentalWorkspaceModule": true
```

---

## 2.문법

### 2.1.package 사용

- package는 import하여 가져온다.
- 대문자로 시작하는것들은 Export된다.
- 따라서 적재한 패키지의 함수를 사용하려면 대문자의것들을 참조해야한다.

### 2.2. function

- func 함수명(인자명 인자타입) 리턴타입 { }
- 연속된 인자가 타입이 같은경우 마지막에만 표기할 수 있다. (x, y int)
- 한 함수는 복수개의 결과를 반환할 수 있다.
  - func add(x, y int) (int, int) { return x + y, x - y }

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

#### 2.3.2. 변수할당

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
