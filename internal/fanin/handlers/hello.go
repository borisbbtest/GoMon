package handlers

import (
	"fmt"
	"net/http"
)

// HelloHandler - хендлер теста сервера
//
// GET [/]
func (h *HTTP) HelloHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "Hello, world!")
}
