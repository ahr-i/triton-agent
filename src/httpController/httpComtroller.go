package httpController

import "net/http"

func JSON(writer http.ResponseWriter, status int, response []byte) {
	writer.WriteHeader(status)
	writer.Header().Set("content-Type", "application/json")
	writer.Write(response)
}
