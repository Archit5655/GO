package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int    `json:"price"`
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("Jsonnnnnn")
	// encoedeJSON()
	DecodingJSON()

}

func encoedeJSON() {
	LCOcourses := []course{
		{"GOLANG KA COURSE", 2345, "Youtube", "Archit123", []string{"webdev", "GO", "LANG"}},
		{"NEXTJS", 23456, "Youtube", "Archit123", []string{"webdev", "tailwind", "js"}},
		{"REact", 23457, "Youtube", "Archit123", nil},
	}

	//  package this data ads json data

	finaljson, err := json.MarshalIndent(LCOcourses, "", "\t")
	if err != nil {
		panic(err)

	}

	fmt.Println(string(finaljson))
}

func DecodingJSON() {
	jsondata := []byte(`
	{
		"coursename": "GOLANG KA COURSE",
		"price": 2345,
		"Platform": "Youtube",
		"tags": [
				"webdev",
				"GO",
				"LANG"
		]
}
	`)

	var gocourse course
	isVAlid := json.Valid(jsondata)
	if isVAlid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsondata, &gocourse)
		fmt.Printf("%#v\n", gocourse)
	} else {
		fmt.Println("Json was not valid")
	}

	//  some cases where you  just want ti add data to your key value
	var onlinedata map[string]interface{}
	json.Unmarshal(jsondata, &onlinedata)

	fmt.Printf("%#v\n", onlinedata)

}
