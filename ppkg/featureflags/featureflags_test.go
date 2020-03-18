package featureflags

import (
	"net/http"
	"testing"
)

func TestIsTestUser(t *testing.T) {
	tests := []struct {
		in   http.Cookie // input
		want bool        // expected result
	}{
		{http.Cookie{Name: "__test_user", Value: "false"}, false},
		{http.Cookie{Name: "__test_user", Value: "true"}, true},
		{http.Cookie{Name: "", Value: ""}, false},
	}

	for _, tt := range tests {
		req, err := http.NewRequest("GET", "/api/feature-flags", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&tt.in)
		if isTestUser(req) != tt.want {
			t.Errorf("isTestUser should return %t, got %t", tt.want, isTestUser(req))
		}
	}
}
