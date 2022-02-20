package main

import "github.com/jiandahao/golanger/pkg/generator/dbgen"

func main() {

	if err := dbgen.GenerateFromDDL("./user.sql", "./model"); err != nil {
		panic(err)
	}
}
