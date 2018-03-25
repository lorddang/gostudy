package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id       string `json:"id"`
	NickName string `json:"nickname"`
	Age      int8   `json:"age"`
	Sex      string `json:"sex"`
}

func testStruct() {
	user := User{
		Id:       "110",
		NickName: "hh",
		Age:      18,
		Sex:      "boy",
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json encoding error: ", err)
		return
	}
	str := string(userJson)
	fmt.Println(str)
	user2 := User{}
	json.Unmarshal(userJson, &user2)
	fmt.Println(user2)
}

func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "alex"
	m["sex"] = "male"
	m["age"] = 19
	mapJson, _ := json.Marshal(m)
	fmt.Println(string(mapJson))
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	json.Unmarshal(mapJson, &m1)
	fmt.Println(m1)
}

func testSlice() {
	var s []map[string] interface{}
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "alex"
	m["sex"] = "male"
	m["age"] = 19
	mapJson, _ := json.Marshal(m)
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	json.Unmarshal(mapJson, &m1)
	m1["age"] = 100
	s = append(s, m)
	s = append(s, m1)
	sJson, _ := json.Marshal(s)
	fmt.Println(string(sJson))

}
func main() {

	testStruct()
	testMap()
	testSlice()
}
