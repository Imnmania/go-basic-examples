package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tag      []string `json:"tag,omitempty"`
}

func EncodeJson() {
	lcoCourses := []Course{
		{"ReactJS Bootcamp", 100, "learncodeonline.in", "pass123", []string{"web-dev", "mobile-dev"}},
		{"Flutter Bootcamp", 199, "learncodeonline.in", "pass456", []string{"web-dev", "mobile-dev"}},
		{"Angular Bootcamp", 50, "learncodeonline.in", "pass789", nil},
	}

	//? package this as json data
	// finalJson, error := json.Marshal(lcoCourses)
	finalJson, error := json.MarshalIndent(lcoCourses, "", "\t")
	if error != nil {
		panic(error)
	}
	fmt.Println(string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "Flutter Bootcamp",
		"price": 199,
		"platform": "learncodeonline.in",
		"tag": [
				"web-dev",
				"mobile-dev"
		]
	}
	`)

	var lcoCourse Course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	/// some cases where you just want to add data to key-value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key => %v || Value => %v || Type => %T \n", k, v, v)
	}
}

func main() {
	fmt.Println("Welcome to json handling in golang!")
	// EncodeJson()
	DecodeJson()
}
