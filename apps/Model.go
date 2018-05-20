package apps

import (
	"net/http"
	"bytes"
)

func GetProducts(writer http.ResponseWriter, request *http.Request){

	byte := &bytes.Buffer{}

	response := "Indra Octama"

	byte.WriteString(response)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(byte.Bytes())

}