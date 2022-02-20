package jsongen

import (
	"bytes"
	"fmt"
	"go/format"
	"sort"
	"strings"
	"text/template"

	"github.com/jiandahao/goutils/convjson"
)

var structTempl *template.Template

func init() {
	structTempl, _ = template.New("struct").Parse(
		`type {{.typeName}} struct {
			{{range .fields}}{{.}}{{end}}
		}`,
	)
}

// GenerateFromString parses json string and converts it into golang structure definition.
func GenerateFromString(s string) (string, error) {
	v, err := convjson.NewValueFromJSONString(s)
	if err != nil {
		return "", err
	}

	var codeString string

	if v.Type() == convjson.TypeArray {
		// TODO: support type: [][]Genrated
		codeString = codeString + `type GeneratedArray []`
		v.Range(func(key string, value *convjson.Value) bool {
			v = value
			return false
		})

		if v.Type() != convjson.TypeMap {
			codeString = codeString + getTypeName(v) + "\n"
			return codeString, nil
		}

		codeString = codeString + "Generated\n"
	}

	var objects map[string]*convjson.Value = map[string]*convjson.Value{"Generated": v}

	for len(objects) > 0 {
		var newObjects = make(map[string]*convjson.Value)

		for name, o := range objects {
			var fields []string
			// TODO: handle map[string]interface{} type
			o.Range(func(key string, value *convjson.Value) bool {
				if value.Type() == convjson.TypeMap {
					newObjects[toCamel(key)] = value
				}

				var typeName string = getTypeName(value)
				if typeName == "" {
					switch value.Type() {
					case convjson.TypeMap:
						typeName = toCamel(key)
					case convjson.TypeArray:
						var basicTypeName string
						var elem *convjson.Value
						value.Range(func(key string, value *convjson.Value) bool {
							basicTypeName = getTypeName(value)
							elem = value
							return false
						})

						if basicTypeName == "" {
							basicTypeName = strings.TrimSuffix(toCamel(key), "s")
							newObjects[basicTypeName] = elem
						}

						typeName = fmt.Sprintf("[]%s", basicTypeName)
					default:
						panic(fmt.Errorf("invalid type: %s", typeName))
					}
				}

				fields = append(fields, fmt.Sprintf("%s %s `json:\"%s,omitempty\"`\n", toCamel(key), typeName, key))
				return true
			})

			sort.Sort(sort.StringSlice(fields))

			var buf bytes.Buffer
			if err := structTempl.Execute(&buf, map[string]interface{}{
				"typeName": name,
				"fields":   fields,
			}); err != nil {
				return "", err
			}

			codeString = codeString + buf.String() + "\n"
		}

		objects = newObjects
	}

	res, err := format.Source([]byte(codeString))
	if err != nil {
		fmt.Println(codeString)
		return "", err
	}

	return string(res), nil
}

// getTypeName returns the type name if v has a built-in type (bool, int, float, string).
func getTypeName(v *convjson.Value) string {
	var typeName string
	switch v.Type() {
	case convjson.TypeBool:
		typeName = "bool"
	case convjson.TypeInt:
		typeName = "int64"
	case convjson.TypeFloat:
		typeName = "float64"
	case convjson.TypeString:
		typeName = "string"
	}

	return typeName
}

func toCamel(s string) string {
	list := strings.Split(s, "_")
	var target []string
	for _, item := range list {
		if !(len(item) == 0 || strings.TrimSpace(item) == "") {
			item = strings.Title(item)
		}
		target = append(target, item)
	}
	return strings.Join(target, "")
}
