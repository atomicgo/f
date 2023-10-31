package f

import (
	"testing"

	"atomicgo.dev/assert"
)

type Person struct {
	Name string
	Age  int
}

func TestFormat(t *testing.T) {
	marvin := Person{Name: "Marvin", Age: 22}
	tests := []struct {
		template string
		data     any
		want     string
	}{
		{"Hello ${Name}", marvin, "Hello Marvin"},
		{"${Name} is ${Age} years old", marvin, "Marvin is 22 years old"},
		{"${Name} is ${Age} years old", map[string]any{"Name": "Marvin", "Age": 22}, "Marvin is 22 years old"},
		{"${Name} is 22 years old? ${Age == 22}", marvin, "Marvin is 22 years old? true"},
		{"${Name} is 22 years old? ${Age == 22 ? 'yes' : 'no'}", marvin, "Marvin is 22 years old? yes"},
		{"${Name} is 60 years old? ${Age == 60 ? 'yes' : 'no'}", marvin, "Marvin is 60 years old? no"},

		{"${1+2}", nil, "3"},
		{"${1+2*3}", nil, "7"},
	}

	for _, tt := range tests {
		t.Run(tt.template, func(t *testing.T) {
			got, err := FormatSafe(tt.template, tt.data)
			if err != nil {
				t.Errorf("Format() error = %v", err)
				return
			}

			if !assert.Equal(got, tt.want) {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
