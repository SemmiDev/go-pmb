package controller

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-clean/entity"
	"go-clean/model"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestStudentController_Create(t *testing.T) {
	studentRepository.DeleteAll()
	createStudentRequest := model.CreateStudentRequest{
		Name:     "Sammi Aldhi Yanto",
		Identifier: "2003113948",
		Email: "sammidev@gmail.com",
	}
	requestBody, _ := json.Marshal(createStudentRequest)

	request := httptest.NewRequest("POST", "/api/students", bytes.NewBuffer(requestBody))
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
	assert.NotNil(t, createStudentResponse.Id)
	assert.Equal(t, createStudentRequest.Name, createStudentResponse.Name)
	assert.Equal(t, createStudentRequest.Identifier, createStudentResponse.Identifier)
	assert.Equal(t, createStudentRequest.Email, createStudentResponse.Email)
}

func TestStudentController_List(t *testing.T) {
	studentRepository.DeleteAll()
	product := entity.Student{
		Id:       uuid.New().String(),
		Name:     "sammidev",
		Identifier: "200311xxx",
		Email: "sammidev@mail.com",
	}
	studentRepository.Insert(product)

	request := httptest.NewRequest("GET", "/api/students", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.([]interface{})
	containsStudent := false

	for _, data := range list {
		jsonData, _ := json.Marshal(data)
		getStudentResponse := model.GetStudentResponse{}
		json.Unmarshal(jsonData, &getStudentResponse)
		if getStudentResponse.Id == product.Id {
			containsStudent = true
		}
	}

	assert.True(t, containsStudent)
}

