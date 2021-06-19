package controller

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-clean/internal/app/model"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestStudentController_Create(t *testing.T) {
	createStudentRequest := model.CreateStudentRequest{
		FullName:           "sammi aldhi yanto",
		Email:              "sammidev@gmail.com",
		PhoneNumber:        "0123821030123",
		Path:               2,
		Year:               2018,
		RegistrationNumber: "2342314121",
	}
	requestBody, _ := json.Marshal(createStudentRequest)

	request := httptest.NewRequest("POST", "/api/v1/students", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 201, 201)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "CREATED", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	createStudentResponse := model.CreateStudentResponse{}
	json.Unmarshal(jsonData, &createStudentResponse)

	assert.NotNil(t, createStudentResponse.Identifier)
	assert.NotNil(t, createStudentResponse.Password)
}

func TestStudentController_List(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/v1/students", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.([]interface{})
	assert.NotNil(t, list)
}

