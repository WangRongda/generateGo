package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

func main() {
	config := getConfig("./test.json")
	fw, err := os.OpenFile("dist/output.go", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if nil != err {
		log.Fatal(err)
	}

	var codeText string
	for config.Next() {
		funcComment := "\n/** " + config.Get("name").(string) + " **/\n"
		funcStart := "func " + config.Get("funcName").(string) + "(w http.ResponseWriter, r *http.Request) {\n"
		funcEnd := "}" + "/** " + config.Get("name").(string) + " " + config.Get("funcName").(string) + "() " + " **/\n"
		var strStruct string
		if nil != config.Get("reqData") {
			strStruct = strStruct + "\tvar reqData "
			json2struct("", config.Get("reqData"), &strStruct, "\t")
		}
		strStruct = strStruct + "\tvar resData "
		if nil != config.Get("resData") {
			json2struct("", config.Get("resData"), &strStruct, "\t")
		} else {
			strStruct = strStruct + "KitResData\n"
		}

		// fmt.Println(strStruct)
		codeText = codeText + funcComment + funcStart + strStruct + deferRes + getReq + funcEnd
	}
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
	case nil:
		*strStruct = *strStruct + "interface{}"
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

func getConfig(filePath string) configer {
	fr, err := os.Open("test.json")
	if nil != err {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(fr)
	if nil != err {
		log.Fatal(err)
	}
	fr.Close()
	var configs []map[string]interface{}
	if err := json.Unmarshal(bytes, &configs); nil != err {
		log.Fatal(err)
	}
	return configer{
		configs,
		nil,
		-1,
		len(configs),
	}
}

type configer struct {
	configs []map[string]interface{}
	config  map[string]interface{}
	index   int
	length  int
}

func (c configer) Get(key string) interface{} {
	switch key {
	case "name", "url", "method", "funcName":
		if nil == c.config[key] {
			log.Fatal("\"", key, "\" is required")
		}
		// case "resData":
		// 	if nil == c.config[key] {
		// 		return map[string]interface{}{}
		// 	}
	}

	return c.config[key]
}

func (c *configer) Next() bool {
	if c.index+1 == c.length {
		return false
	} else {
		c.index++
		c.config = c.configs[c.index]
		return true
	}
}

const deferRes string = `
	defer func() {
		if "" == resData.Msg {
			resData.Msg = codemsg[resData.Code]
		}
		resBody, _ := json.Marshal(resData)
		w.Head().Set("Content-Type", "application/json")
		w.Write(resBod)
	}()`

const getReq string = `
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resData.Code = RequestParse
		return
	}
	if err := json.Unmarshal(reqBody, &reqData); nil != err {
		resData.Code = RequestParse
		return
	}
`

const parseBizSign string = `
	if reqBizSign, err := ParseBizsign()
	`

const checkPremission string = `
`
