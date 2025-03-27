package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWrite   http.ResponseWriter
}

// crea una estructura por defecto
func CreateDefaultResponse(rw http.ResponseWriter) Response {

	return Response{
		Status:      http.StatusOK,
		respWrite:   rw,
		contentType: "application/json",
	}
}

// metodo para manejar la respuesta al cliente
func (resp *Response) Send() {
	resp.respWrite.Header().Set("application/json", resp.contentType)
	resp.respWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)

	response.Data = data
	response.Send()
}

// Metodo para manejar error
func (resp *Response) NotFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource no Found"
}

// funcion para manejar error
func SendNoFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.Send()

}

// metodo
func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity no Found"
}
func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()

}
