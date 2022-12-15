package main

import (
	"net/http"
	"path"
)

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func newPathResolver() *pathResolver {
	return &pathResolver{
		handlers: make(map[string]http.HandlerFunc),
	}
}

func (pr *pathResolver) Add(path string, handler http.HandlerFunc) {
	pr.handlers[path] = handler
}

func (pr *pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// combine request method and path
	check := r.Method + " " + r.URL.Path

	// go through registered paths
	for pattern, handlerFunc := range pr.handlers {
		// check if current path matches one of the registered ones
		if ok, err := path.Match(pattern, check); ok && err == nil {
			// call the function to handle request
			handlerFunc(w, r)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	
	// if there is no match for the path, return the message that the page is not found
	http.NotFound(w, r)
}