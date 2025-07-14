package auth

import (
	"net/http"
	"strings"
	"testing"
)

type testCase struct {
	name           string
	headers        http.Header
	expectedAPIKey string
	expectedErr    string
}

func TestGetAPIKey(t *testing.T) {
	tests := []testCase{
		{
			name:           "valid authorization header",
			headers:        http.Header{"Authorization": {"ApiKey abcde"}},
			expectedAPIKey: "abcde",
		},
		{
			name:           "no authorization header",
			headers:        http.Header{"Content-Type": {"application/json"}},
			expectedAPIKey: "",
			expectedErr:    "no authorization header included",
		},
		{
			name:           "malformed authorization header",
			headers:        http.Header{"Authorization": {"ApiKey"}},
			expectedAPIKey: "",
			expectedErr:    "malformed authorization header",
		},
	}

	for _, test := range tests {
		apiKey, err := GetAPIKey(test.headers)
		if apiKey != test.expectedAPIKey {
			t.Errorf("Unexpected APIKey returned from test '%s'", test.name)
		}

		if err != nil {
			if strings.Contains(err.Error(), test.expectedErr) {
				return
			}
			t.Errorf("Unexpected Error returned from test '%s'", test.name)
			return
		}
	}
}
