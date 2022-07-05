package usepackage

import (
	"encoding/json"
	"fmt"
)

/**
 * 基本数据类型
 */
func testBasic() {
	num := 1.23
	marshal, err := json.Marshal(num)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
	}
	fmt.Println("after marshal: ", string(marshal))
	var num2 float32
	json.Unmarshal(marshal, &num2)
	fmt.Printf("after unmarshal: %f", num2)
}

type School struct {
	Address string `json:"address"`
}

/**
 * Student xe
 */
type Student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	School School `json:"school"`
}

/**
 * 结构体
 */
func testStruct() {
	student := Student{
		Name: "Peng",
		Age:  23,
		School: School{
			Address: "GuangZhou, GuangDong Province",
		},
	}
	marshal, err := json.Marshal(student)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("after marshal: ", string(marshal))
	var stu2 Student
	json.Unmarshal(marshal, &stu2)
	fmt.Printf("after unmarshal, %v", stu2)
}

/**
 * Map
 */
func testMap() {
	var stu = make(map[string]interface{})
	stu["name"] = "peng"
	stu["age"] = 23
	stu["school"] = map[string]interface{}{"address": "GuangZhou"}
	marshal, err := json.Marshal(stu)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("after marshal: ", string(marshal))
	// var stu2 map[string]interface{}
	var stu2 Student // 反序列化为结构体
	json.Unmarshal(marshal, &stu2)
	fmt.Printf("after unmarshal, %v", stu2)
}

/**
 * 切片
 */
func testSlice() {
	var slice []map[string]interface{}
	var stu = make(map[string]interface{})
	stu["name"] = "peng"
	stu["age"] = 23
	stu["school"] = map[string]interface{}{"address": "GuangZhou"}
	slice = append(slice, stu)
	marshal, err := json.Marshal(slice)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("after marshal: ", string(marshal))
	var stu2 []map[string]interface{}
	// var stu2 []Student // 反序列化为结构体
	json.Unmarshal(marshal, &stu2)
	fmt.Printf("after unmarshal, %v", stu2)
}
