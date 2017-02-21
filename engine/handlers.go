package engine

import (
	"net/http"
	"fmt"
)

func PostDeployHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return;
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type",  "application/json")

	fmt.Fprint(w, `{"status": "ok"}`)
}