package json

import (
	"encoding/json"
	"fmt"
	"log"
)

func Encode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Decode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}

type Person struct {
	Name string
	Age  string
	Address string
}

func EncodeDecodeUsage(){
	j := make(map[string]string)
	j["name"] = "shane"
	j["age"] = "18"
	j["address"] = "shanghai"
	encode, _ := Encode(j)
	fmt.Println(encode)

	bytes := []byte(encode)
	var obj Person
	res := Decode(bytes, &obj)
	if res != nil {
		log.Fatalln(res)
	}
	fmt.Println(obj)
}


type Mobile struct {
	Resultcode string `json:"resultcode"`
	Reason     string `json:"reason"`
	Result     struct {
		Province string `json:"province"`
		City     string `json:"city"`
		Areacode string `json:"areacode"`
		Zip      string `json:"zip"`
		Company  string `json:"company"`
		Card     string `json:"card"`
	} `json:"result"`
}

func DecodeJson() {
	// 解析json数据 1. 先将json转成struct; 2. 然后使用json.Unmarshal()即可。
	// json转struct，在线工具：https://mholt.github.io/json-to-go/
	json_str := `
{
			"resultcode": "200",
			"reason": "Return Successd!",
			"result": {
				"province": "浙江",
				"city": "杭州",
				"areacode": "0571",
				"zip": "310000",
				"company": "中国移动",
				"card": ""
			}
		}`
	var a Mobile
	err := json.Unmarshal([]byte(json_str), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.Result)
}


