package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jiandahao/goutils/convjson"
)

var swaggerFile = flag.String("f", "", "specify the swagger file you wanna covert")
var outputFile = flag.String("o", "./gen.swagger.json", "specify the output file")

func main() {
	flag.Parse()

	if *swaggerFile == "" {
		panic("invalid file")
	}

	data, err := ioutil.ReadFile(*swaggerFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	mdata := make(map[string]interface{})
	if err := json.Unmarshal(data, &mdata); err != nil {
		fmt.Println(err)
		return
	}

	val := convjson.NewValue(mdata)

	paths, err := val.Get("paths")
	if err != nil {
		fmt.Println(err)
		return
	}

	pathKeys := paths.RawValue().MapKeys()

	for _, key := range pathKeys {
		for _, method := range []string{"get", "post", "put", "delete", "patch", "head", "options", "connect", "trace"} {
			ref, err := val.Get(fmt.Sprintf("paths.%s.%s.responses.200.schema.$ref", key, method))
			if err != nil {
				continue
			}

			refStr := ref.MustString()
			defName := strings.TrimPrefix(refStr, "#/definitions/")
			//  wrap orignal response as code-message format
			val.Set(fmt.Sprintf("definitions.genSwagger%s", defName), map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"code": map[string]string{
						"type": "integer",
					},
					"msg": map[string]string{
						"type": "string",
					},
					"data": map[string]interface{}{
						"$ref": refStr,
					},
				},
			})
			val.Set(fmt.Sprintf("paths.%s.%s.responses.200.schema.$ref", key, method), fmt.Sprintf("#/definitions/genSwagger%s", defName))
		}
	}

	// setting unexpected error response.
	val.Set("definitions.rpcStatus", map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"code": map[string]interface{}{
				"type":   "integer",
				"format": "int32",
			},
			"msg": map[string]interface{}{
				"type": "string",
			},
		},
	})

	res, err := json.MarshalIndent(val.RawValue().Interface(), "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	dir := filepath.Dir(*outputFile)
	os.MkdirAll(dir, os.ModePerm)

	if err := ioutil.WriteFile(*outputFile, res, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}
}
