package main

import "github.com/jiandahao/golanger/pkg/generator/dbgen"

func main() {

	if err := dbgen.GenerateFromDDL("./user.sql", "test_project", "./model", false); err != nil {
		panic(err)
	}

	// if err := dbgen.GenerateFromDDL("./user.sql", "test_project", "./model", true); err != nil {
	// 	panic(err)
	// }
}
