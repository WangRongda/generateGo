package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	strStruct := "var reqData struct {\n"
	getReqData(config["resData"].(map[string]interface{}), &strStruct)
	strStruct = strStruct + "}\n"
	fmt.Println(strStruct)
	// reqData :=
	codeText := funcComment + funcStart + funcEnd
	if _, err = fw.WriteString(codeText); err != nil {
		panic(err)
	}
}

func getReqData(key string, value interface{}, strStruct *string) {
	r := []rune(key)
	r[0] = unicode.ToUpper(r[0])
	structKey := "\t" + string(r) + " "
	switch value.(type) {
	case string:
		*strStruct = *strStruct + structKey + value.(string) + "`json:\"" + key + "\"`\n"
	case map[string]interface{}:
		*strStruct = *strStruct + structKey + "struct {\n"
		for k, v := range value.(map[string]interface{}) {
			getReqData(k, v, strStruct)
		}
		*strStruct = *strStruct + "} `json:\"" + key + "\"`\n"
	case []interface{}:
		var typ string
		if len(value.([]interface{})) == 0 {
			typ = "interface{}"
		} else {
			switch value.([]interface{})[0].(type) {
			case string:
				*strStruct = *strStruct + structKey + "[]" + value.(string) + "`json:\"" + key + "\"`\n"
			case map[string]interface{}:
				*strStruct = *strStruct + structKey + "[]struct {\n"
				for k, v := range value.(map[string]interface{}) {
					getReqData(k, v, strStruct)
				}
				*strStruct = *strStruct + "} `json:\"" + key + "\"`\n"
			case []interface{}:
				*strStruct = *strStruct + structKey + "[]"
				getReqData()
			}
		}
		*strStruct = *strStruct + structKey + "[]" + typ
	default:
		log.Fatal("JSON file error")
	}
}
