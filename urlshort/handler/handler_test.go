package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/phss/gophercises/urlshort/handler"
)

func TestMapHandler(t *testing.T) {
	pathsToUrls := map[string]string{
		"/in-map": "http://example.com/in-map",
	}

	tt := []struct {
		name             string
		requestPath      string
		expectedRedirect string
	}{
		{name: "path in map", requestPath: "/in-map", expectedRedirect: "http://example.com/in-map"},
		{name: "fallback", requestPath: "/not-in-map", expectedRedirect: "http://example.com/redirect-through-fallback"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request, _ := http.NewRequest("GET", tc.requestPath, nil)
			response := httptest.NewRecorder()

			handler := handler.MapHandler(pathsToUrls, fallbackHandler())
			handler.ServeHTTP(response, request)

			actualRedirect := response.Header().Get("Location")
			if actualRedirect != tc.expectedRedirect {
				t.Fatalf("redirect expected %s, got %s", tc.expectedRedirect, actualRedirect)
			}
		})
	}
}

func TestYAMLHandler(t *testing.T) {
	yaml := `
- path: /in-yaml
  url: http://example.com/in-yaml
`

	tt := []struct {
		name             string
		requestPath      string
		expectedRedirect string
	}{
		{name: "path in map", requestPath: "/in-yaml", expectedRedirect: "http://example.com/in-yaml"},
		{name: "fallback", requestPath: "/not-in-yaml", expectedRedirect: "http://example.com/redirect-through-fallback"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request, _ := http.NewRequest("GET", tc.requestPath, nil)
			response := httptest.NewRecorder()

			handler, _ := handler.YAMLHandler([]byte(yaml), fallbackHandler())
			handler.ServeHTTP(response, request)

			actualRedirect := response.Header().Get("Location")
			if actualRedirect != tc.expectedRedirect {
				t.Fatalf("redirect expected %s, got %s", tc.expectedRedirect, actualRedirect)
			}
		})
	}
}

func TestYAMLHandler_invalidYAML(t *testing.T) {
	_, err := handler.YAMLHandler([]byte("bad yaml"), fallbackHandler())

	if !strings.HasPrefix(err.Error(), "yaml: unmarshal errors") {
		t.Fatalf("error %v", err)
	}
}

func fallbackHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://example.com/redirect-through-fallback", http.StatusSeeOther)
	}
}
