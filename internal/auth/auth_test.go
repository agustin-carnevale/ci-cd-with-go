package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}

	tests := []test{
		{input: http.Header{
			"Authorization": []string{"ApiKey this-is-my-api-key"},
		}, want: "this-is-my-api-key"},
		{input: http.Header{
			"Authorization": []string{"Bearer this-is-my-api-key"},
		}, want: ""},
		{input: http.Header{
			"Auth": []string{"ApiKey this-is-my-api-key"},
		}, want: ""},
		{input: http.Header{}, want: ""},
	}

	for _, tc := range tests {
		got, _ := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("\nExpected: %v\nGot: %v\n", tc.want, got)
		}
	}
}
