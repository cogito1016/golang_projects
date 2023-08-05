package bufio_prac

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
