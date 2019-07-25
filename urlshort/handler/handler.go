package handler

import (
	"net/http"
)

// MapHandler redirects requests for registered paths, in map format, otherwise fallbacks.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, found := pathsToUrls[r.URL.Path]
		if found {
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler redirects requests for registered paths, in YAML format, otherwise fallbacks.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		fallback.ServeHTTP(w, r)
	}, nil
}
