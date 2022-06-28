package filter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type (
	OpNameType    string
	OperatorType  string
	FieldNameType string

	Operator struct {
		shortName OpNameType
		operator  OperatorType
	}

	Conditions struct {
		FieldName FieldNameType
		Operator  OperatorType
		Value     interface{}
	}
)

var allRegisteredOpsM = make(map[OpNameType]Operator)

func registerOperator(name string, op string) Operator {
	o := Operator{
		shortName: OpNameType(name),
		operator:  OperatorType(op),
	}
	allRegisteredOpsM[o.shortName] = o

	return o
}

// all permitted operators
var (
	Equal            = registerOperator("eq", "=")
	LessThanEqual    = registerOperator("lte", "<=")
	LessThan         = registerOperator("lt", "<")
	GreaterThanEqual = registerOperator("gte", ">=")
	GreaterThan      = registerOperator("gt", ">")
	Contains         = registerOperator("contains", "LIKE")
	In               = registerOperator("in", "IN")
)

type Parser struct {
	supportedFields map[FieldNameType][]Operator
}

// NewParser returns a new filter parser
func NewParser(supportedFields map[FieldNameType][]Operator) *Parser {
	return &Parser{
		supportedFields: supportedFields,
	}
}

// Parse parse JSON-Formatted filters. The db query sql and args are returned.
//
// for example:
/*
	{
		"name": {
			"eq": "jack"
		},
		"create_time": {
			"lt": 1234,
			"gt": 123
		}
	}
*/
func (f *Parser) Parse(filters string) (string, []interface{}, error) {
	if filters == "" {
		return "", nil, nil
	}

	filterM := make(map[FieldNameType]map[OpNameType]interface{})

	if err := json.Unmarshal([]byte(filters), &filterM); err != nil {
		return "", nil, err
	}

	var whereStatement []*Conditions
	// validate filter fields
	for fieldName, conditions := range filterM {
		if _, ok := f.supportedFields[fieldName]; !ok {
			continue
		}

		allSupportedOps := f.supportedFields[fieldName]
		var allSupportedOpsM = make(map[OpNameType]struct{})

		for _, item := range allSupportedOps {
			allSupportedOpsM[item.shortName] = struct{}{}
		}

		for opName, opValue := range conditions {
			if _, ok := allSupportedOpsM[opName]; !ok {
				return "", nil, fmt.Errorf("unsupported operator (%s) for field (%s)", opName, fieldName)
			}

			if _, ok := allRegisteredOpsM[opName]; !ok {
				return "", nil, fmt.Errorf("unsupported operator (%s) for field (%s)", opName, fieldName)
			}

			if err := f.validate(opName, opValue); err != nil {
				return "", nil, err
			}

			opValue = fixValueType(opValue)

			switch opName {
			case Contains.shortName:
				opValue = "%" + fmt.Sprint(opValue) + "%"
			default:
			}

			whereStatement = append(whereStatement, &Conditions{
				FieldName: fieldName,
				Operator:  allRegisteredOpsM[opName].operator,
				Value:     opValue,
			})
		}
	}

	sort.SliceStable(whereStatement, func(i, j int) bool {
		return whereStatement[i].FieldName < whereStatement[j].FieldName
	})

	var conds []string
	var args []interface{}

	for _, item := range whereStatement {
		conds = append(conds, fmt.Sprintf("%s %s ?", item.FieldName, item.Operator))
		args = append(args, item.Value)
	}

	return strings.Join(conds, " AND "), args, nil
}

func (f *Parser) validate(opName OpNameType, opValue interface{}) error {
	operator := allRegisteredOpsM[opName]
	switch operator {
	case In:
		switch reflect.ValueOf(opValue).Kind() {
		case reflect.Slice, reflect.Array:
		default:
			return fmt.Errorf("In operator only support value with array type")
		}
	case Equal, LessThan, LessThanEqual, GreaterThan, GreaterThanEqual, Contains:
		switch reflect.ValueOf(opValue).Kind() {
		case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
			return fmt.Errorf("invalid filter value. operators, like eq, lt, lte, gt, only support basic data type (e.g: int, string)")
		}
	default:
		return fmt.Errorf("unknown operator: %v", operator)
	}

	return nil
}

func fixValueType(val interface{}) interface{} {
	switch v := val.(type) {
	case float64:
		if float64(int64(v)) == v {
			return int64(v)
		}
	case []interface{}:
		var elems []interface{}
		for _, item := range v {
			elems = append(elems, fixValueType(item))
		}
		return elems
	}

	return val
}

func Where(fieldName FieldNameType, op Operator, val interface{}) *Filter {
	filter := &Filter{
		fields: make(map[FieldNameType]map[OpNameType]interface{}),
	}

	return filter.Where(fieldName, op, val)
}

type Filter struct {
	fields map[FieldNameType]map[OpNameType]interface{}
}

func NewFilter(filter string) (*Filter, error) {
	filterM := make(map[FieldNameType]map[OpNameType]interface{})
	if filter == "" {
		return &Filter{fields: filterM}, nil
	}

	if err := json.Unmarshal([]byte(filter), &filterM); err != nil {
		return nil, err
	}

	return &Filter{fields: filterM}, nil
}

// Get get filter condition value
func (f *Filter) Get(fieldName FieldNameType, op Operator) (interface{}, bool) {
	if f.fields == nil {
		return nil, false
	}

	fieldCond, ok := f.fields[fieldName]
	if !ok {
		return nil, false
	}

	val, ok := fieldCond[op.shortName]
	if !ok {
		return nil, false
	}

	return fixValueType(val), true
}

// Where adds filter condition
func (f *Filter) Where(fieldName FieldNameType, op Operator, val interface{}) *Filter {
	if f.fields == nil {
		return f
	}

	if f.fields[fieldName] == nil {
		f.fields[fieldName] = make(map[OpNameType]interface{})
	}

	f.fields[fieldName][op.shortName] = val

	return f
}

func (f *Filter) HasField(fieldName FieldNameType) bool {
	if f.fields == nil {
		return false
	}

	_, ok := f.fields[fieldName]
	return ok
}

func (f *Filter) String() string {
	data, _ := json.Marshal(f.fields)
	return string(data)
}
