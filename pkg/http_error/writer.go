package httperror

import (
	"net/http"
)

// Write writes a httpError or error to ResponseWriter
func Write(w http.ResponseWriter, err error) {
	httpErr, ok := err.(Error)

	if ok {
		http.Error(w, httpErr.Error(), httpErr.StatusCode())
		return
	}

	http.Error(w, httpErr.Error(), http.StatusUnprocessableEntity)
}
