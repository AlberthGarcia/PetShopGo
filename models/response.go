package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Create a struct to response
type Response struct {
	Status      int64       `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	resWriter   http.ResponseWriter
	contentType string
}

//Built a response default if all it's okay
func ResponseDefault(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		resWriter:   rw,
		contentType: "application/json",
	}
}

//Modify the propierties to send a response struct
func (resp *Response) ModifyResponse() {
	//Changed the header for an type application/JSON
	resp.resWriter.Header().Set("Content-type", resp.contentType)
	//Write the status of the header
	resp.resWriter.WriteHeader(int(resp.Status))

	//Become the response in a slice of byte whit marshal
	output, err := json.Marshal(&resp)
	if err != nil {
		panic(err)
	}

	//Send the response become in string to read
	fmt.Fprintln(resp.resWriter, string(output))
}

//Send the response with all struct
func SendResponse(rw http.ResponseWriter, data interface{}) {
	response := ResponseDefault(rw)
	response.Data = data
	response.ModifyResponse()
}
