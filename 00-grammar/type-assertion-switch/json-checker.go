package type_assertion_switch

import (
	"encoding/json"
	"fmt"
)

func RunJsonChecker() {
	JSONmap := make(map[string]interface{})
	json.Unmarshal([]byte(JSONrecord), &JSONmap)
	typeChecker(JSONmap)
}

var JSONrecord = `{
"Flag" : true,
"String" : "Hello World!",
"Array" : ["a","b","c"],
"Int" : 7,
"Entity" : {
"innerInt" : 1,
"innerString" : "Bye World!",
"innerEntity" : {
  "innerinnerInt" : 1
}
} 
}`

func typeChecker(m map[string]interface{}) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("it is string", k, c)
		case int:
			fmt.Println("it is int", k, c)
		case map[string]interface{}:
			fmt.Println("it is map", k, c)
			typeChecker(v.(map[string]interface{}))
		default:
			fmt.Printf("It is what..? : %v %T \n", k, c)
		}
	}
}
