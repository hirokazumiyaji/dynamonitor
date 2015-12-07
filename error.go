package main

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, status int, errors []error) {
	errorStrings := make([]string, len(errors))
	for _, err := range errors {
		errorStrings = append(errorStrings, err.Error())
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"errors": errorStrings,
		},
	)
}
