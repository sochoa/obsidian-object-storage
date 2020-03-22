package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteResponse(responseWriter http.ResponseWriter,
	bucket string,
	pathTrail string,
	status int,
	msg *string,
	err error) {

	var resp_data map[string]string = make(map[string]string, 0)
	resp_data["bucket"] = bucket
	resp_data["path"] = pathTrail
	bytesBuffer := new(bytes.Buffer)
	jsonEncoder := json.NewEncoder(bytesBuffer)

	// Add the message to the response, if its set
	if msg != nil {
		resp_data["details"] = *msg
	}

	// Add the error to the response, if its set
	err_msg := "n/a"
	if err != nil {
		err_msg = fmt.Sprintf("%v", err)
		resp_data["error"] = err_msg
	}

	// Set response status code normally, if its not an error
	if status >= 400 {
		http.Error(responseWriter, err_msg, status)
	} else {
		responseWriter.WriteHeader(status)
	}

	// Common encode JSON to a string and write the response
	jsonEncoder.Encode(resp_data)
	responseWriter.Write(bytesBuffer.Bytes())
}

func WriteStatusWithMessage(responseWriter http.ResponseWriter,
	bucket string,
	pathTrail string,
	status int,
	msg string) {
	WriteResponse(responseWriter, bucket, pathTrail, status, &msg, nil)
}

func WriteStatus(responseWriter http.ResponseWriter,
	bucket string,
	pathTrail string,
	status int) {
	WriteResponse(responseWriter, bucket, pathTrail, status, nil, nil)
}

func WriteErrorStatusWithMessage(responseWriter http.ResponseWriter,
	bucket string,
	pathTrail string,
	status int,
	msg string,
	err error) {
	WriteResponse(responseWriter, bucket, pathTrail, status, &msg, err)
}

func WriteErrorStatus(responseWriter http.ResponseWriter,
	bucket string,
	pathTrail string,
	status int,
	err error) {
	WriteResponse(responseWriter, bucket, pathTrail, status, nil, err)
}
