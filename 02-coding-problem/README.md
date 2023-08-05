# GO언어 백준 풀이 팁

## 1. 패키지

- 패키지는 main
- 함수도 main

## 2. array를 다룰 때

### 2.1. 출력

- fmt.println시에는 양옆에 []가 붙는다.
- 따라서 반복문을 사용하여 print함수를 이용하여 출력한다.

```go
for i := 0; i < N; i++ {
    fmt.Print(arr[i]," ");
}
```

### 2.2. array에 값 유무체크 (contains, find ...)

- 없다..는데.. 진짜인가?
- 그냥 반복문을 돌려서 체크한다.

```go

```
