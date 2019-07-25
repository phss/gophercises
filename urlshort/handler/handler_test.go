package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phss/gophercises/urlshort/handler"
)

var fallback http.HandlerFunc

func TestMapHandler(t *testing.T) {
	pathsToUrls := map[string]string{
		"/in-map": "http://example.com/in-map",
	}
	fallback = func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://example.com/redirect-through-fallback", http.StatusSeeOther)
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

			handler := handler.MapHandler(pathsToUrls, fallback)
			handler.ServeHTTP(response, request)

			actualRedirect := response.Header().Get("Location")
			if actualRedirect != tc.expectedRedirect {
				t.Fatalf("redirect expected %s, got %s", tc.expectedRedirect, actualRedirect)
			}
		})
	}
}
