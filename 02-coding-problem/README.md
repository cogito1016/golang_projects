# GO언어 백준 풀이 팁

## 1. 패키지

- 패키지는 main
- 함수도 main

## 2. array출력 시

- fmt.println시에는 양옆에 []가 붙는다.
- 따라서 반복문을 사용하여 print함수를 이용하여 출력한다.

```go
for i := 0; i < N; i++ {
    fmt.Print(arr[i]," ");
}
```
