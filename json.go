// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type courses struct {
// 	Name     string   `json:"coursename"`
// 	Fees     int64    `json:"price"`
// 	Platform string   `json:"palce"`
// 	Password string   `json:"-"`
// 	Tags     []string `"password:"tags,imiteempty"`
// }

// func main() {
// 	fmt.Println("About json(encoding) file")
// 	encodeJson()
// 	decodeJson()
// }

// func encodeJson() {
// 	padmashreeCourse := []courses{
// 		{"BiT Hons", 6000000, "PiC", "bit@123", []string{"bit", "hons"}},
// 		{"BHM Hons", 6000000, "PiC", "rit@123", []string{"bit", "hons"}},
// 		{"BCA Hons", 6000000, "PiC", "mit@123", []string{"bit", "hons"}},
// 	}

// 	finalJson, err := json.MarshalIndent(padmashreeCourse, "", "\t")

// 	if err != nil {
// 		panic(err)

// 	}
// 	fmt.Println(string(finalJson))
// }

// func decodeJson() {
// 	jsonData := []byte(`
// 	{
// 		"coursename": "BiT Hons",
// 		"price": 6000000,
// 		"palce": "PiC",
// 		"Tags": ["bit","hons"]
// 	}
// 	`)

// 	var padmashreeCoursee courses

// 	checkValid := json.Valid(jsonData)

// 	if checkValid {
// 		fmt.Println("json is valid")
// 		json.Unmarshal(jsonData, &padmashreeCoursee)
// 		fmt.Printf("%#v\n", padmashreeCoursee)

// 	} else {
// 		fmt.Println("The json is not working")
// 	}

// 	var mymapdata map[string]interface{}
// 	json.Unmarshal(jsonData, &mymapdata)
// 	fmt.Println(mymapdata)

// 	for k, v := range mymapdata {
// 		fmt.Println(k, " :", v)
// 	}
// }
