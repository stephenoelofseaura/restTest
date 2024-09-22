package db

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ConvertAndWriteData(data any, w http.ResponseWriter) error {
	convertedData, err := json.Marshal(data)
	if err != nil {
		return errors.New("Error marshalling data, Error: " + err.Error())
	}
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(convertedData)
	if writeErr != nil {
		return errors.New("Error writing response")
	}
	return nil
}
