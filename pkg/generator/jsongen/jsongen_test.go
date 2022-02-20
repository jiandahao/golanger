package jsongen

import (
	"fmt"
	"go/format"
	"strings"
	"testing"
)

func TestGenerateFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "basic type",
			args: args{s: `{
				"float_val": 1.0,
				"bool_val": false,
				"string_val":"string"
			}`},
			want: fmt.Sprintf(
				`type Generated struct {
					BoolVal bool %s
					FloatVal float64  %s
					StringVal string %s 
				}`,
				"`json:\"bool_val,omitempty\"`",
				"`json:\"float_val,omitempty\"`",
				"`json:\"string_val,omitempty\"`",
			),
			wantErr: false,
		},
		{
			name: "with embed object",
			args: args{s: `{
				"user": {
					"name":"jian",
					"age": "12"
				}
			}`},
			want: fmt.Sprintf(
				`type Generated struct {
					User User %s
				}
				type User struct {
					Age string %s
					Name string %s
				}`,
				"`json:\"user,omitempty\"`",
				"`json:\"age,omitempty\"`",
				"`json:\"name,omitempty\"`",
			),
			wantErr: false,
		},
		{
			name: "array with object elem",
			args: args{s: `{
				"user": [{
					"name":"jian",
					"age": "12"
				}]
			}`},
			want: fmt.Sprintf(
				`type Generated struct {
					User []User %s
				}
				type User struct {
					Age string %s
					Name string %s
					
				}`,
				"`json:\"user,omitempty\"`",
				"`json:\"age,omitempty\"`",
				"`json:\"name,omitempty\"`",
			),
			wantErr: false,
		},
		{
			name: "array with object elem",
			args: args{s: `[{"name":"jian"}]`},
			want: fmt.Sprintf(
				`type GeneratedArray []Generated
				 type Generated struct {
					Name string %s
				 }`,
				"`json:\"name,omitempty\"`",
			),
			wantErr: false,
		},
		{
			name:    "array with built-in type elem",
			args:    args{s: `["name","test"]`},
			want:    `type GeneratedArray []string`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			res, err := format.Source([]byte(tt.want))
			if err != nil {
				t.Errorf("format.Source() error = %v", err)
				return
			}

			want := string(res)

			replacer := strings.NewReplacer(" ", "", "\t", "", "\n", "")

			if replacer.Replace(got) != replacer.Replace(want) {
				t.Errorf("\nGenerateFromString() = \n%v, \nwant =\n%v", got, string(res))
			}
		})
	}
}
