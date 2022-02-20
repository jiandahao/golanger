package main

import (
	"fmt"

	"github.com/jiandahao/golanger/pkg/generator/jsongen"
)

func main() {
	res, err := jsongen.GenerateFromString(`{
		"city":["beijing", "shanghai"],
		"cotunry":"china",
		"countryCode":"CN",
		"age":12,
		"users": [{"name":"jian", "age":24}],
		"jobs":{
			"name":"jobname"
		}
	}`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	res, err = jsongen.GenerateFromString(`[{"name":"jian", "age":24}]`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
