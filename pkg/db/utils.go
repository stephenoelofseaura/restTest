package db

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"
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

func SplitCamelCaseString(input string) string {
	re := regexp.MustCompile("([A-Z][a-z0-9]+)")
	return strings.Join(re.FindAllString(input, -1), " ")
}
