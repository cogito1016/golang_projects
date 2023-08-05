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