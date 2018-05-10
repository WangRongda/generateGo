package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"unicode"
)

// type configer struct {
//     Name
// }

func main() {
	fr, err := os.Open("test.json")
	if nil != err {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(fr)
	if nil != err {
		log.Fatal(err)
	}
	fr.Close()
	var config map[string]interface{}
	if err := json.Unmarshal(bytes, &config); nil != err {
		log.Fatal(err)
	}

	fw, err := os.OpenFile("test.kit", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if nil != err {
		log.Fatal(err)
	}

	funcComment := "\n/* ------------\n * " + config["name"].(string)
	if nil != config["nameEn"] {
		funcComment = funcComment + "\n" + config["nameEn"].(string)
	}
	funcComment = funcComment + "\n * ------------*/\n"
	funcStart := "func " + config["funcName"].(string) + "(w http.ResponseWriter, r *http.Request) {\n"
	funcEnd := "\n}\n"
	getReqData(config["resData"].(map[string]interface{}))
	// reqData :=
	codeText := funcComment + funcStart + funcEnd
	if _, err = fw.WriteString(codeText); err != nil {
		panic(err)
	}
}

func getReqData(reqData map[string]interface{}) (strStruct string) {
	strStruct = "var reqData struct {\n"
	for key, value := range reqData {
		// fmt.Println(i)
		// fmt.Println(v)
		fmt.Println(key, ": ", reflect.TypeOf(value))
		switch value.(type) {
		case map[string]interface{}:
			getReqData(value.(map[string]interface{}))
		case string:
			r := []rune(key)
			r[0] = unicode.ToUpper(r[0])
			strStruct = strStruct + "\t" + string(r) + " " + value.(string) + "\n"
		default:
			// log.Fatal("JSON file error")
		}
	}
	strStruct = strStruct + "}\n"
	fmt.Println(strStruct)
	return ""
}
