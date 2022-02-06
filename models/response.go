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
func responseDefault(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		resWriter:   rw,
		contentType: "application/json",
	}
}

//method to Modify the propierties to send a response struct
func (resp *Response) modifyResponse() {
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
func SendResponse(rw http.ResponseWriter, data interface{}, message string) {
	response := responseDefault(rw)
	response.Data = data
	response.Message = message
	response.modifyResponse()
}

//method to modify the status if there's an error
func (resp *Response) notFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Request fail, not found"
}

func SendNotFound(rw http.ResponseWriter) {
	response := responseDefault(rw)
	response.notFound()
	response.modifyResponse()
}

//method to modify the status if there's an error
func (resp *Response) unproccesableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "Request fail, unprocessable entity"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := responseDefault(rw)
	response.unproccesableEntity()
	response.modifyResponse()
}
