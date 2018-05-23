package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

type coder struct {
	reqStruct, resStruct, deferRes, parseBizSign, checkPremission string
}

type configer struct {
	config map[string]interface{}
	index  int
	code   coder
	fw     map[string]*os.File
}

func main() {
	configFile := flag.String("c", "./feedback.json", "config file")
	flag.Parse()
	config := getConfig(*configFile)
	defer func() {
		for _, f := range config.fw {
			end := "\n\n// <<<<<<<<<<<<<<<<<<<<<<<<<<<< generation@wangrongda " + time.Now().String() + "\n\n"
			if _, err := f.WriteString(end); err != nil {
				panic(err)
			}
			f.Close()
		}
	}()

	config.generateAll()
}

func (c configer) appendToFile(filePath string, codeText string) {
	fullPath := c.Get("projectDir").(string) + filePath
	_, exist := c.fw[filePath]
	if !exist {
		var err error
		c.fw[filePath], err = os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if nil != err {
			log.Fatal(err)
		}
		start := "\n\n// >>>>>>>>>>>>>>>>>>>>>>>>>>>> generation@wangrongda " + time.Now().String() + "\n\n"
		if _, err := c.fw[filePath].WriteString(start); err != nil {
			panic(err)
		}
	}
	if _, err := c.fw[filePath].WriteString(codeText); err != nil {
		panic(err)
	}
}

func json2struct(key string, value interface{}, strStruct *string, strIndent string, bJsonTag bool) {
	var jsonTag string
	if bJsonTag && "" != key {
		jsonTag = " `json:\"" + key + "\"`"
	}
	jsonTag += "\n"
	switch value.(type) {
	case string:
		*strStruct = *strStruct + value.(string) + jsonTag
	case map[string]interface{}:
		*strStruct = *strStruct + "struct {\n"
		for k, v := range value.(map[string]interface{}) {
			*strStruct = *strStruct + strIndent + "\t" + key2field(k) + " "
			json2struct(k, v, strStruct, strIndent+"\t", bJsonTag)
		}
		*strStruct = *strStruct + strIndent + "}" + jsonTag
	case []interface{}:
		*strStruct = *strStruct + "[]"
		if len(value.([]interface{})) == 0 {
			value.([]interface{})[0] = "interface{}"
		}
		json2struct(key, value.([]interface{})[0], strStruct, strIndent, bJsonTag)
	case nil:
		*strStruct = *strStruct + "\n"
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

func (c configer) generateReqStruct(UT bool) (strStruct string) {
	if nil != c.Get("reqData") {
		if UT {
			strStruct = "\ttype reqData "
		} else {
			strStruct = "\tvar reqData "
		}
		var bJsonTag bool
		if "get" == c.Get("method") {
			bJsonTag = false
		} else {
			bJsonTag = true
		}
		json2struct("", c.Get("reqData"), &strStruct, "\t", bJsonTag)
	}
	c.code.reqStruct = strStruct
	return
}
func (c configer) generateResStruct(UT bool) (strStruct string) {
	if UT {
		strStruct = "\ttype resData "
	} else {
		strStruct = "\tvar resData "
	}
	if nil != c.Get("resData") {
		switch c.Get("resData").(type) {
		case map[string]interface{}:
			_, exist := c.Get("resData").(map[string]interface{})["result"]
			if exist {
				c.Get("resData").(map[string]interface{})["KitResData"] = nil
				json2struct("", c.Get("resData"), &strStruct, "\t", true)
			} else {
				tmp := map[string]interface{}{}
				tmp["KitResData"] = nil
				tmp["Result"] = c.Get("resData")
				json2struct("", tmp, &strStruct, "\t", true)
			}
		default:
			tmp := map[string]interface{}{}
			tmp["KitResData"] = nil
			tmp["result"] = c.Get("resData")
			json2struct("", tmp, &strStruct, "\t", true)
		}
	} else {
		strStruct += "KitResData\n"
	}
	c.code.resStruct = strStruct
	return
}

func (c configer) generateHandleFunc() (codeText string) {
	funcComment := "\n/** " + c.Get("name").(string) + "\n * [" + strings.ToUpper(c.Get("method").(string)) + "] " + c.Get("url").(string) + " */\n"
	funcStart := "func " + c.Get("funcName").(string) + "(w http.ResponseWriter, r *http.Request) {\n"
	funcEnd := "}" + "/** " + c.Get("name").(string) + " " + c.Get("funcName").(string) + "() " + " **/\n"

	codeText = funcComment + funcStart + c.generateReqStruct(false) + c.generateResStruct(false) +
		deferRes + c.generateReqCode() + c.generateCheckPremission() + funcEnd
	return codeText
}

func (c configer) registeRouter() {
	code := "mux.HandlerFunc(\"" + c.Get("method").(string) + "\", \"" + c.Get("url").(string) + "\", " + c.Get("funcName").(string) + ")\n"
	c.appendToFile(c.Get("routerFile").(string)+".go", code)
}

func (c configer) Get(key string) interface{} {
	raw := c.config["raw"].([]interface{})[c.index].(map[string]interface{})
	switch key {
	case "name", "url", "method", "funcName", "fileName":
		if nil == raw[key] {
			log.Fatal("\"", key, "\" is required")
		}
	case "checkPremission":
		if nil == raw[key] || false == raw[key] {
			return -1
		}
	case "parseBizSign", "kitLog":
		if nil == raw[key] {
			return false
		}
	case "routerFile", "projectDir":
		if nil == c.config[key] {
			log.Fatal("\"", key, "\", is required")
		}
		return c.config[key]
	}

	return raw[key]
}

func (c *configer) Next() bool {
	if c.index+1 == len(c.config["raw"].([]interface{})) {
		return false
	} else {
		c.index++
		return true
	}
}

const deferRes string = `
	defer func() {
		resData.Status = resData.Code
		if "" == resData.Msg {
			resData.Msg = codemsg[resData.Code]
		}
		resBody, _ := json.Marshal(&resData)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resBody)
	}()
`

const parseBizSign = `
	reqBizSign, err := parseBizSign(r.Header.Get("bizsign"))
	if err != nil {
		Warn("parseBizSign:", err)
		resData.Code = DecBizSign
		return
	}	
`

//不需要BizSign
const parseRequestParameters1 = `
	if _, code := parseRequestParameters(r, &reqData, false); Success != code {
		resData.Code = code
		return
	}
`

//需要BizSign
const parseRequestParameters2 = `
	reqBizSign, code := parseRequestParameters(r, &reqData, true)
	if Success != code {
		resData.Code = code
		return
	}
`

const checkPremission string = `
	if checkPremission(%d, reqBizSign, %s) {
		Warn("checkPremission:")
		resData.Code = ProNotInGroup
		return
	}
`

func (c configer) generateCheckPremission() string {
	v := c.Get("checkPremission").(float64)
	switch v {
	case 0:
		return fmt.Sprintf(checkPremission, v, "0")
	case 2:
		return fmt.Sprintf(checkPremission, v, "reqData.pid")
	default:
		return ""
	}
}

//获取请求参数（到结构体reqData）
//获取bizSign(如果需要)
func (c configer) generateReqCode() string {
	if c.bNeedBizSign() {
		if nil != c.Get("reqData") {
			return parseRequestParameters2
		} else {

			return parseBizSign
		}
	} else {
		if nil != c.Get("reqData") {
			return parseRequestParameters1
		} else {
			return ""
		}
	}
}

func (c configer) bNeedBizSign() bool {
	if false == c.Get("parseBizSign") && false == c.Get("checkPremission") && false == c.Get("kitLog") {
		return false
	} else {
		return true
	}
}

func (c configer) generateWiKi() string {
	title := "## " + c.Get("name").(string) + "\n\n"
	url := "[" + strings.ToUpper(c.Get("method").(string)) + "] " + c.Get("url").(string) + "\n\n"
	req := ""
	if "get" == c.Get("method") {
		req = "* 请求参数\n\n"
		req += "key | value | 说明 \n----|-----|---- \n"
		for key, value := range c.Get("reqData").(map[string]interface{}) {
			req += key + " | " + value.(string) + " | .\n"
		}
	} else {
		req = "* 请求示例\n\n``` JSON\n"
		reqJSON, err := json.MarshalIndent(c.Get("reqData"), "", "    ")
		if nil != err {
			log.Fatal()
		}
		req += string(reqJSON) + "\n```\n"
	}
	res := "* 响应示例\n\n``` JSON\n"
	resJSON, err := json.MarshalIndent(c.Get("resData"), "", "    ")
	if nil != err {
		log.Fatal()
	}
	res += string(resJSON) + "\n```\n"
	return title + url + req + res
}

func (c configer) generateAll() (code string) {
	for c.Next() {
		code = c.generateHandleFunc()
		utCode := c.generateUTfunc()
		wiki := c.generateWiKi()
		c.registeRouter()
		c.appendToFile(c.Get("fileName").(string)+".go", code)
		c.appendToFile(c.Get("fileName").(string)+"_test.go", utCode)
		c.appendToFile(c.Get("fileName").(string)+".md", wiki)
	}
	return
}

// func (c configer) generateKitlog() (code string) {

// }

func getConfig(filePath string) configer {
	fr, err := os.Open(filePath)
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
	return configer{
		config,
		-1,
		coder{},
		map[string]*os.File{},
	}
}

func (c configer) generateUTfunc() string {
	// funcComment := "\n/** UT " + c.Get("name").(string) + "\n * [" + strings.ToUpper(c.Get("method").(string)) + "] " + c.Get("url").(string) + " */\n"
	funcStart := "\nfunc " + "Test" + c.Get("funcName").(string) + "(t *testing.T) {\n"
	funcEnd := "}" + "// " + "Test" + c.Get("funcName").(string) + "(t *testing.T) " + "\n"

	s1 := `mock, db := KitMock.NewDB()
	oldDB := dbpool.db
	defer func() {
		dbpool.db = oldDB
	}()
	dbpool.db = db
	reqBizSign, bizSignString := KitMock.BizSign()
	`

	s2 := `
	testcases := []struct {
		input reqData 
		expect 
	}{

	}

	for i := range testcases {
		input := testcases[i].input
		reqBody, _ := json.Marshal(input)

		req := httptest.NewRequest(
			"%s",
			"%s",
			bytes.NewReader(reqBody),
		)
		w := httptest.NewRecorder()

		req.Header.Add("bizsign", bizSignString)

		%s(w, req)

		result := w.Result()
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			t.Error(err)
		}
		if result.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, %d", result.StatusCode)
		}

	} // for case
`

	codeText := funcStart + c.generateReqStruct(true) + c.generateResStruct(true) + s1 +
		fmt.Sprintf(s2, strings.ToUpper(c.Get("method").(string)), c.Get("url").(string), c.Get("funcName").(string)) +
		funcEnd

	return codeText
}
