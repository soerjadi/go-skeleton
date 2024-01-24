package user

import "net/http"

func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "hello world", nil
}
