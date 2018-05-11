package main

import (
	"encoding/json"
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

	fw, err := os.OpenFile("output.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	strStruct := "\tvar resData "
	json2struct("", config["resData"].(map[string]interface{}), &strStruct, "\t")
	// fmt.Println(strStruct)
	codeText := funcComment + funcStart + strStruct + funcEnd
	if _, err = fw.WriteString(codeText); err != nil {
		panic(err)
	}
}

func json2struct(key string, value interface{}, strStruct *string, strIndent string) {
	var jsonTag string
	if "" != key {
		jsonTag = " `json:\"" + key + "\"`\n"
	} else {
		jsonTag = "\n"
	}
	switch value.(type) {
	case string:
		*strStruct = *strStruct + value.(string) + jsonTag
	case map[string]interface{}:
		*strStruct = *strStruct + "struct {\n"
		for k, v := range value.(map[string]interface{}) {
			*strStruct = *strStruct + strIndent + "\t" + key2field(k) + " "
			json2struct(k, v, strStruct, strIndent+"\t")
		}
		*strStruct = *strStruct + strIndent + "}" + jsonTag
	case []interface{}:
		*strStruct = *strStruct + "[]"
		if len(value.([]interface{})) == 0 {
			value.([]interface{})[0] = "interface{}"
		}
		json2struct(key, value.([]interface{})[0], strStruct, strIndent)
	default:
		log.Fatal("JSON file error")
	}
}

func key2field(key string) string {
	if "id" == key {
		return "ID"
	}
	r := []rune(key)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
