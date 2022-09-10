package filter

import "testing"

func TestFilter_Parse(t *testing.T) {
	type args struct {
		filters string
	}

	supportedFields := map[FieldNameType][]Operator{
		"name":        {Equal},
		"create_time": {LessThan, GreaterThan},
		"email":       {Equal, Contains},
	}

	testCases := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid case",
			args: args{
				filters: `{"name":{"eq": 123}, "create_time":{"lt": 1234, "gt": 123}}`,
			},
			want:    `create_time < ? AND create_time > ? AND name = ?`,
			wantErr: false,
		},
		{
			name: "valid case (contains)",
			args: args{
				filters: `{"email":{"contains": "123"}}`,
			},
			want:    `email LIKE ?`,
			wantErr: false,
		},
		{
			name: "invalid operator",
			args: args{
				filters: `{"name":{"lt": 123}}`,
			},
			wantErr: true,
		},
		{
			name: "invalid field",
			args: args{
				filters: `{"unavailable_field":{"lt": 123}}`,
			},
			wantErr: false, // skip unavailable field
		},
		{
			name: "invalid filter json",
			args: args{
				filters: `{"name":{"lt": 123}`,
			},
			wantErr: true,
		},
	}

	f := &Parser{
		supportedFields: supportedFields,
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			got, _, err := f.Parse(tt.args.filters)
			if (err != nil) != tt.wantErr {
				t.Errorf("Filter.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Filter.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
