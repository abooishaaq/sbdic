package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqError struct {
	Err string `json:"err"`
}

func SendErrMsg(w http.ResponseWriter, err error) {
	fmt.Println(err)
	msg := ReqError{Err: err.Error()}
	marshalled, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(marshalled)
}
