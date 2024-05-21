package net

import "net/http"

// WithMethods is a helper function to check the HTTP method before calling the handler.
// If the method is not allowed, it will return a 405 Method Not Allowed.
// If the method is allowed, it will call the handler.
// The method parameter is variadic, so you can pass in multiple methods.
// If no methods are passed in, it will default to GET.
//
// Example:
//
//	This will only allow to GET requests to the /api/version route and call the getAppVersion handler:
//	http.Handle("/api/version", net.WithMethods(getAppVersion))
//
//	With multiple methods:
//	http.Handle("/api/version", net.WithMethods(getAppVersion, http.MethodGet, http.MethodPost))
func WithMethods(handler func(http.ResponseWriter, *http.Request), methods ...string) http.HandlerFunc {
	if len(methods) == 0 {
		// default to GET
		methods = append(methods, http.MethodGet)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, m := range methods {
			if r.Method == m {
				handler(w, r)
				return
			}
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
