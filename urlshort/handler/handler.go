package handler

import (
	"net/http"

	"gopkg.in/yaml.v2"
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
	pathsToUrls, err := yamlToMap(yml)
	return MapHandler(pathsToUrls, fallback), err
}

func yamlToMap(yml []byte) (map[string]string, error) {
	type YamlData struct {
		Path string `yaml:path`
		Url  string `yaml:url`
	}

	yamlData := []YamlData{}
	err := yaml.Unmarshal(yml, &yamlData)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, data := range yamlData {
		pathsToUrls[data.Path] = data.Url
	}
	return pathsToUrls, nil
}
