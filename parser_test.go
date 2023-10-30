package f

import (
	"reflect"
	"testing"
)

func Test_splitIntoParts(t *testing.T) {
	tests := []struct {
		template string
		want     []Part
	}{
		{"Hello ${Name}", []Part{{Value: "Hello ", Parsed: true}, {Value: "Name", Parsed: false}}},
		{"${Name} is ${Age} years old", []Part{{Value: "Name", Parsed: false}, {Value: " is ", Parsed: true}, {Value: "Age", Parsed: false}, {Value: " years old", Parsed: true}}},
		{"${Name} is ${Age} years old ${a == b}", []Part{{Value: "Name", Parsed: false}, {Value: " is ", Parsed: true}, {Value: "Age", Parsed: false}, {Value: " years old ", Parsed: true}, {Value: "a == b", Parsed: false}}},
	}

	for _, tt := range tests {
		t.Run(tt.template, func(t *testing.T) {
			if got := splitIntoParts(tt.template); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
