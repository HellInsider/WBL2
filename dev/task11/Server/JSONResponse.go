package Server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Msg interface{} `json:"error"`
}

type Result struct {
	Msg interface{} `json:"result"`
}

func jsonResponse(err bool, w http.ResponseWriter, code int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if err {
		errResp := Error{
			Msg: msg.(string),
		}
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			http.Error(w, "Error: Response JSON", http.StatusInternalServerError)

		}
	} else {
		resResp := Result{
			Msg: msg,
		}
		if err := json.NewEncoder(w).Encode(resResp); err != nil {
			http.Error(w, "Error: Response JSON", http.StatusInternalServerError)
		}
	}
}

func isValid(w http.ResponseWriter, r *http.Request, args ...string) bool {
	if r.Method != args[0] {
		jsonResponse(true, w, http.StatusMethodNotAllowed, fmt.Sprintf("%v: didn't expected method", r.Method))
		return false
	}
	for _, p := range args[1:] {
		if !r.URL.Query().Has(p) {
			jsonResponse(true, w, http.StatusServiceUnavailable, "Error: Can't find param "+p)
			return false
		}
	}
	return true
}
