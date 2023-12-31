# 패키지

## 1. bufio 입출력 패키지

- bufio를 사용하면 입출력 속도가 빨라진다.

```go

import (
"bufio"
"fmt"
"os"
)

func Run(){
//bufio패키지를 사용한 입력 예시

    //1. bufio.NewReader(os.Stdin)을 통해 입력을 받을 준비를 한다.
    //2. 준비된 입력을 통해 입력을 받는다.
    //3. 입력받은 내용을 출력한다.
    reader := bufio.NewReader(os.Stdin);
    fmt.Println("Enter your name: ");
    name, _ := reader.ReadString('\n');
    fmt.Println("Hello, ", name);

	//Scanner를 사용한 입력 예시
	//1. bufio.NewScanner(os.Stdin)을 통해 입력을 받을 준비를 한다.
	//2. 준비된 입력을 통해 입력을 받는다.
	//3. 입력받은 내용을 출력한다.
	scanner := bufio.NewScanner(os.Stdin);
	fmt.Println("Enter your name: ");
	scanner.Scan();
	name2 := scanner.Text();
	fmt.Println("Hello, ", name2);

	//Reader와 Scanner의 차이점
	//1. Reader는 줄바꿈 문자를 포함한 한 줄을 읽어들인다.
	//2. Scanner는 줄바꿈 문자를 포함하지 않은 한 줄을 읽어들인다.

    //bufio패키지를 사용한 출력 예시
    //1. bufio.NewWriter(os.Stdout)을 통해 출력을 받을 준비를 한다.
    //2. 버퍼에 출력할 내용을 쌓는다.
    //3.Flush()를 통해 버퍼에 내용을 출력한다.
    writer := bufio.NewWriter(os.Stdout);
    fmt.Fprint(writer, "Hello, ");
    writer.Flush();
}
```

- 입력받은 데이터를 숫자로 처리할 때,

    - itoa를 사용하기보다는 byte[] (scanner.text())를 사용해도 좋다.

- "0 0 0" , "abc def gii" 등과같이 split이 필요하다면
    - split말고 strings.Fields()를 사용하면 배열이된다

## 2. 형변환을 할 때 strconv 패키지

- strconv 패키지를 사용한다.

```go
package strconv_prac

import "strconv"

func Run(){
	//strconv.Itoa()
	//정수 -> 문자열
	//문자열로 변환된 값 반환
	strconv.Itoa(12345) //"12345"
	strconv.Itoa(-12345) //"-12345"

	//strconv.Atoi()
	//문자열 -> 정수
	//문자열로 변환된 값 반환
	strconv.Atoi("12345") //12345
	strconv.Atoi("-12345") //-12345

	//strconv.FormatFloat()
	//실수 -> 문자열
	//문자열로 변환된 값 반환
	strconv.FormatFloat(1.2345, 'f', 2, 64) //"1.23"
	strconv.FormatFloat(1.2345, 'e', 2, 64) //"1.23e+00"

	//strconv.ParseFloat()
	//문자열 -> 실수
	//문자열로 변환된 값 반환
	strconv.ParseFloat("1.2345", 64) //1.2345
	strconv.ParseFloat("1.2345e+00", 64) //1.2345

	//strconv.FormatBool()
	//불리언 -> 문자열
	//문자열로 변환된 값 반환
	strconv.FormatBool(true) //"true"
	strconv.FormatBool(false) //"false"

	//strconv.ParseBool()
	//문자열 -> 불리언
	//문자열로 변환된 값 반환
	strconv.ParseBool("true") //true
	strconv.ParseBool("false") //false

	//strconv.FormatInt()
	//정수 -> 문자열
	//문자열로 변환된 값 반환
	strconv.FormatInt(12345, 10) //"12345"
	strconv.FormatInt(-12345, 10) //"-12345"

	//strconv.ParseInt()
	//문자열 -> 정수
	//문자열로 변환된 값 반환
	strconv.ParseInt("12345", 10, 64) //12345
	strconv.ParseInt("-12345", 10, 64) //-12345

	//strconv.FormatUint()
	//정수 -> 문자열
	//문자열로 변환된 값 반환
	strconv.FormatUint(12345, 10) //"12345"

	//strconv.ParseUint()
	//문자열 -> 정수
	//문자열로 변환된 값 반환
	strconv.ParseUint("12345", 10, 64) //12345
}
```

- 형변환 오류
    - any로 받는것들은 {any type value}.(type)으로 변환을 시켜줘야한다.

```bash
cannot use value (variable of type any) as int value in argument to strconv.Itoa: need type assertion
```

## 3. 날짜를 다룰 때 time 패키지

```go
package print_problem

import (
	"fmt"
	"time"
)

func RunBoj10699(){
	t := time.Now().UTC().Format("2006-01-02"); //이때 2006-01-02는 고정된 값이다.
	fmt.Println(t);
}
```

### 3.1. 2006-01-02란?

- 2006-01-02 15:04:05 는 날짜를 나타내는 것이 아닌 yyyy-mm-dd hh:mm:ss 를 의미합니다. 해당 날짜는 아무 무작위로 한 것처럼 보이지만 뜻이 숨어있음
- 2006년 1월 2일 3시 4분 5초는 time 패키지에서 제공하는 기본값이다.

```go
//로케이션설정
loc,_ := time.LoadLocation("Asia/Seoul");
t := time.Now();
t = t.In(loc);

//포멧설정
t := time.Now().Format("2006-01-02");
```

## 4. 수학을 다룰 때 math 패키지

- math 패키지를 사용한다.
- math.Abs()는 절대값을 구하는 함수이다.
    - math.Abs(float64)를 사용한다.
    - 따라서 int형을 넣어주기 위해서는 math.Abs(float64(int형))을 사용한다.

```go
N = int(math.Abs(float64(N)));
```

### 4.1 2의 n승을 구할때는 비트연산자를 활용해도 ㅗ딘다.
- math패키지를 이용하면 math.Pow(2,n)을 사용할 수 있다.
- 비트연산자를 이용하여 처리하면 1 << uint(n)을 사용할 수 있다.

```go
size := 1 << uint(n-1)
```

## 5. unicode 패키지

- unicode 패키지는 유니코드를 다루는 패키지이다.
- 유니코드란?
    - 유니코드는 전 세계의 모든 문자를 컴퓨터에서 일관되게 표현하고 다룰 수 있도록 설계된 산업 표준이다.
- 어느때 사용하냐면..
    - 문자열의 길이를 구할 때 사용한다.
        - 한글의 경우에는 3바이트로 인식하기 때문에 len()함수를 사용하면 3이 아닌 9가 출력된다.
        - 따라서 한글의 경우에는 unicode 패키지를 사용하여서 길이를 구해야한다.
    - 대,소문자를 구분할 때 사용한다.
      ```go
      if unicode.IsUpper(char){
          fmt.Printf("%c", unicode.ToLower(char));
          continue;
      }
      ```

## 6. postmanerator 패키지
- postman의 컬렉션을 Export하면 Json이 생성되는데, 이 Json을 기반으로 html문서를 만들어주는 패키지
- 주의할 점
  - postman export 시 2.1 포멧으로 출력해야 함

### 6.1. 설치
- https://github.com/aubm/postmanerator
- mac용 설치파일이 없는걸로 보인다.
- 따라서, Mac의 경우
  1. git clone 하여 레포를 받는다
  2. ```go get github.com/aubm/postmanerator``` 를 통해 postmanerator 패키지를 받는다 
  3. ```go install```를 통해 바이너리를 생성한다
  4. clone 받은 레포 접근 > go run ./main.go 를 실행한다.
     5. 필수 파라미터의 경우 아래 커맨드를 참조한다(실행파일의 경우 사용하는 커맨드이다)
        6. ```postmanerator -output=./doc.html -collection=$YOUR_PROJECT/postman/collection.json```
     7. 필자의 경우, 파라미터 전달 이슈가 있어서 아래의 방식으로 해결했다.
        8. configuration/configuration.go 접근
        9. func parseCommandFlags()함수 내 collection 및 output의 value값을 하드코딩하여 경로조정
