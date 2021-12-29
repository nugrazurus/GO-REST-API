package app

import (
	u "go-rest-api/utils"
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler  {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "This resource was not found on our server"))
		next.ServeHTTP(w, r)	
	})
}