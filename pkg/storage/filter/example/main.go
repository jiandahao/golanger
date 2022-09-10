package main

import (
	"fmt"
	"time"

	"github.com/jiandahao/golanger/pkg/storage/filter"
)

func main() {
	filterParser := filter.NewParser(map[filter.FieldNameType][]filter.Operator{
		"name":        {filter.Equal, filter.Contains, filter.In},
		"create_time": {filter.GreaterThan, filter.LessThan},
		"email":       {filter.Contains},
	})

	conds, args, err := filterParser.Parse(`{"name":{"eq": 123, "contains":"234", "in": [1]}, "create_time":{"lt": 1234, "gt": 123}}`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(conds, args)

	filters := filter.Where("name", filter.Equal, "jian").
		Where("create_time", filter.LessThan, time.Now().Add(2*time.Hour).Unix()).
		Where("create_time", filter.GreaterThan, time.Now().Unix()).
		Where("email", filter.Contains, "jack")

	fmt.Println(filters.String())

	conds, args, err = filterParser.Parse(filters.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(conds, args)
}
