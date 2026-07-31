package main

import (
	"fmt"
	"encoding/json"
)

func convertingToJson() {

	var name,address string

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("enter your address: ")
	fmt.Scan(&address)

	result := map[string]string {
		"name" : name,
		"address": address,
	}

	jsonData,err:=json.Marshal(result)

	if err != nil {
		fmt.Println("An Error occurred : while displaying json data",err)
		return
	}

	fmt.Println(string(jsonData))

}

