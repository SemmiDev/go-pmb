package controller

import (
	"bytes"
	"encoding/json"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestRegistrationController_Create(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.NotNil(t, createRegistrationResponse.Username)
	assert.NotNil(t, createRegistrationResponse.Password)
	assert.Equal(t, model.S2Bill, createRegistrationResponse.Bill)
	assert.NotEqual(t, model.S1D3D4Bill, createRegistrationResponse.Bill)
}

func TestRegistrationController_CreateFailedEmailIsExists(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:  "Sammi Aldhi Yanto",
		Email: "sammidev@gmail.com",
		Phone: "082387325971",
	}
	registrationService.Create(&createRegistrationRequest, model.S2)

	createRegistrationRequest2 := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest2)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t, "mailer has been recorded", webResponse.ErrorMessage)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationRequest2)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedPhoneIsExists(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:  "Sammi Aldhi Yanto",
		Email: "sammidev@gmail.com",
		Phone: "082387325971",
	}
	registrationService.Create(&createRegistrationRequest, model.S2)

	createRegistrationRequest2 := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev2@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest2)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t, "phone has been recorded", webResponse.ErrorMessage)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationRequest2)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedNameIsEmpty(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}(map[string]interface{}{"Required_Name": "Name Is Empty"}),
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedRequestsIsEmpty(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "",
		Email:   "",
		Phone:   "",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}{
			"Required_Email": "Email Is Empty",
			"Required_Name":  "Name Is Empty",
			"Required_Phone": "Phone Is Empty",
			"invalid_Phone":  "Phone Number Is Not Valid",
			"invalid_Email":  "Email Is Not Valid"},
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedInvalidPhone(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "sammi",
		Email:   "sammi@gmail.com",
		Phone:   "aoksoadal",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}{
			"invalid_Phone": "Phone Number Is Not Valid"},
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedInvalidPhoneAndEmail(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "sammi",
		Email:   "sammiasam",
		Phone:   "aoksoadal",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}{
			"invalid_Phone": "Phone Number Is Not Valid",
			"invalid_Email": "Email Is Not Valid",
		},
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedProgramNotRecognize(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "izzah",
		Email:   "izzah@gmail.com",
		Phone:   "08123912389123",
		Program: "xxxx",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}{
			"Program_Not_Available": "Please Chose Between S1D3D4 or S2",
		},
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_UpdateSuccess(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	temp, _ := registrationService.Create(&createRegistrationRequest, model.S2)

	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: temp.VirtualAccount,
	}

	requestBody, _ := json.Marshal(updateStatusRequest)
	request := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "Ok", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	assert.Equal(t, map[string]interface{}(map[string]interface{}{"status": "updated"}), webResponse.Data)
}

func TestRegistrationController_UpdateFailedEmptyVA(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	registrationService.Create(&createRegistrationRequest, model.S2)

	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: "",
	}

	requestBody, _ := json.Marshal(updateStatusRequest)
	request := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t, map[string]interface{}(map[string]interface{}{"Required_VA": "Virtual Account Is Empty"}), webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)
}

func TestRegistrationController_UpdateFailedVaNotFound(t *testing.T) {
	registrationRepository.DeleteAll()
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: "1241231321231",
	}

	requestBody, _ := json.Marshal(updateStatusRequest)
	request := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 500, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 500, webResponse.Code)
	assert.Equal(t, "Internal Server Error", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t, "va not found", webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)
}

func TestRegistrationController_LoginSuccess(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])
}

func TestRegistrationController_LoginFailedUsernameNotFound(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: "xxxxxxxxxxxxxxxxxxxx",
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 400, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 400, webResponse3.Code)
	assert.Equal(t, "Bad Request", webResponse3.Status)
	assert.Equal(t, true, webResponse3.Error)
	assert.Equal(t, "mongo: no documents in result", webResponse3.ErrorMessage)
	assert.Nil(t, webResponse3.Data)
}

func TestRegistrationController_LoginFailedStatusStillFalse(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------

	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 400, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 400, webResponse3.Code)
	assert.Equal(t, "Bad Request", webResponse3.Status)
	assert.Equal(t, true, webResponse3.Error)
	assert.Equal(t, "please pay the billing first", webResponse3.ErrorMessage)
	assert.Nil(t, webResponse3.Data)
}

func TestRegistrationController_LoginFailedPasswordNotMatching(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: "xxxxxxxxxxxxxxxxxxxxxx",
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 400, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 400, webResponse3.Code)
	assert.Equal(t, "Bad Request", webResponse3.Status)
	assert.Equal(t, true, webResponse3.Error)
	assert.Equal(t, "crypto/bcrypt: hashedPassword is not the hash of the given password", webResponse3.ErrorMessage)
	assert.Nil(t, webResponse3.Data)
}

func TestRegistrationController_LogoutSuccess(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])

	logoutRequest := httptest.NewRequest("POST", "/api/v1/auth/logout", nil)
	logoutRequest.Header.Set("Authorization", "Bearer "+loginResponse["access_token"])
	logoutRequest.Header.Set("Accept", "application/json")

	logoutResponse, _ := app.Test(logoutRequest)
	assert.Equal(t, 200, logoutResponse.StatusCode)
	logoutResponseBody, _ := ioutil.ReadAll(logoutResponse.Body)
	logoutWebResponse := model.WebResponse{}
	json.Unmarshal(logoutResponseBody, &logoutWebResponse)

	expected := map[string]string(map[string]string{"message": "Successfully logged out"})
	assert.Equal(t, 200, logoutWebResponse.Code)
	assert.Equal(t, "Ok", logoutWebResponse.Status)
	assert.Equal(t, false, logoutWebResponse.Error)
	assert.Equal(t, nil, logoutWebResponse.ErrorMessage)

	logoutResponseData := make(map[string]string)
	jsonDataLogoutResponse, _ := json.Marshal(logoutWebResponse.Data)
	json.Unmarshal(jsonDataLogoutResponse, &logoutResponseData)
	assert.Len(t, logoutResponseData, 1)
	assert.Equal(t, expected, logoutResponseData)
}

func TestRegistrationController_LogoutFailedSignatureIsInvalid(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])

	logoutRequest := httptest.NewRequest("POST", "/api/v1/auth/logout", nil)
	logoutRequest.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjVkOGJhZjFiLTU5ZWItNGViMC05ZjhhLWY2N2NkYzE1OTQzZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYyNzc4NDQ2OSwidXNlcl9pZCI6ImU1ZjEwZWYwLWE2ZDMtNDNhZS04YjBjLWMzNmNiOTcyODM5ZSJ9.jLcRKSjmPHa2axCCraImnw1-w9cThy4ZrM-PX0hBL47")
	logoutRequest.Header.Set("Accept", "application/json")

	logoutResponse, _ := app.Test(logoutRequest)
	assert.Equal(t, 401, logoutResponse.StatusCode)
	logoutResponseBody, _ := ioutil.ReadAll(logoutResponse.Body)
	logoutWebResponse := model.WebResponse{}
	json.Unmarshal(logoutResponseBody, &logoutWebResponse)

	assert.Equal(t, 401, logoutWebResponse.Code)
	assert.Equal(t, "Unauthorized", logoutWebResponse.Status)
	assert.Equal(t, true, logoutWebResponse.Error)
	assert.Equal(t, "signature is invalid", logoutWebResponse.ErrorMessage)
	assert.Equal(t, nil, logoutWebResponse.Data)
}

func TestRegistrationController_LogoutFailedSignatureIllegalBase64(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])

	logoutRequest := httptest.NewRequest("POST", "/api/v1/auth/logout", nil)
	logoutRequest.Header.Set("Authorization", "Bearer aeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjVkOGJhZjFiLTU5ZWItNGViMC05ZjhhLWY2N2NkYzE1OTQzZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYyNzc4NDQ2OSwidXNlcl9pZCI6ImU1ZjEwZWYwLWE2ZDMtNDNhZS04YjBjLWMzNmNiOTcyODM5ZSJ9.jLcRKSjmPHa2axCCraImnw1-w9cThy4ZrM-PX0hBL48")
	logoutRequest.Header.Set("Accept", "application/json")

	logoutResponse, _ := app.Test(logoutRequest)
	assert.Equal(t, 401, logoutResponse.StatusCode)
	logoutResponseBody, _ := ioutil.ReadAll(logoutResponse.Body)
	logoutWebResponse := model.WebResponse{}
	json.Unmarshal(logoutResponseBody, &logoutWebResponse)

	assert.Equal(t, 401, logoutWebResponse.Code)
	assert.Equal(t, "Unauthorized", logoutWebResponse.Status)
	assert.Equal(t, true, logoutWebResponse.Error)
	assert.Equal(t, "illegal base64 data at input byte 37", logoutWebResponse.ErrorMessage)
	assert.Equal(t, nil, logoutWebResponse.Data)
}

func TestRegistrationController_LogoutFailedSignatureInvalidCharacter(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])

	logoutRequest := httptest.NewRequest("POST", "/api/v1/auth/logout", nil)
	logoutRequest.Header.Set("Authorization", "Bearer xyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjVkOGJhZjFiLTU5ZWItNGViMC05ZjhhLWY2N2NkYzE1OTQzZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYyNzc4NDQ2OSwidXNlcl9pZCI6ImU1ZjEwZWYwLWE2ZDMtNDNhZS04YjBjLWMzNmNiOTcyODM5ZSJ9.jLcRKSjmPHa2axCCraImnw1-w9cThy4ZrM-PX0hBL48")
	logoutRequest.Header.Set("Accept", "application/json")

	logoutResponse, _ := app.Test(logoutRequest)
	assert.Equal(t, 401, logoutResponse.StatusCode)
	logoutResponseBody, _ := ioutil.ReadAll(logoutResponse.Body)
	logoutWebResponse := model.WebResponse{}
	json.Unmarshal(logoutResponseBody, &logoutWebResponse)

	assert.Equal(t, 401, logoutWebResponse.Code)
	assert.Equal(t, "Unauthorized", logoutWebResponse.Status)
	assert.Equal(t, true, logoutWebResponse.Error)
	assert.Equal(t, "invalid character 'Ç' looking for beginning of value", logoutWebResponse.ErrorMessage)
	assert.Equal(t, nil, logoutWebResponse.Data)
}

func TestRegistrationController_RefreshTokenSuccess(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])

	// REFRESH

	refreshRequest := map[string]string{
		"refresh_token": loginResponse["refresh_token"],
	}

	requestBodyRefreshToken, _ := json.Marshal(refreshRequest)
	requestRefreshToken := httptest.NewRequest("POST", "/api/v1/auth/refresh", bytes.NewBuffer(requestBodyRefreshToken))
	requestRefreshToken.Header.Set("Content-Type", "application/json")
	requestRefreshToken.Header.Set("Accept", "application/json")
	responseRefreshToken, _ := app.Test(requestRefreshToken)
	assert.Equal(t, 201, responseRefreshToken.StatusCode)
	responseBodyRefreshToken, _ := ioutil.ReadAll(responseRefreshToken.Body)
	webResponseRefreshToken := model.WebResponse{}
	json.Unmarshal(responseBodyRefreshToken, &webResponseRefreshToken)
	assert.Equal(t, 201, webResponseRefreshToken.Code)
	assert.Equal(t, "Created", webResponseRefreshToken.Status)
	assert.Equal(t, false, webResponseRefreshToken.Error)
	assert.Equal(t, nil, webResponseRefreshToken.ErrorMessage)

	datamapRefreshToken := make(map[string]string)
	jsonDataRefreshToken, _ := json.Marshal(webResponseRefreshToken.Data)
	datamapRefresh := datamapRefreshToken
	json.Unmarshal(jsonDataRefreshToken, &datamapRefresh)
	assert.NotNil(t, datamapRefresh["access_token"])
	assert.NotNil(t, datamapRefresh["refresh_token"])
}

func TestRegistrationController_RefreshTokenFailedBecauseNotInputRefreshToken(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])

	// REFRESH

	refreshRequest := map[string]string{}

	requestBodyRefreshToken, _ := json.Marshal(refreshRequest)
	requestRefreshToken := httptest.NewRequest("POST", "/api/v1/auth/refresh", bytes.NewBuffer(requestBodyRefreshToken))
	requestRefreshToken.Header.Set("Content-Type", "application/json")
	requestRefreshToken.Header.Set("Accept", "application/json")
	responseRefreshToken, _ := app.Test(requestRefreshToken)
	assert.Equal(t, 400, responseRefreshToken.StatusCode)
	responseBodyRefreshToken, _ := ioutil.ReadAll(responseRefreshToken.Body)
	webResponseRefreshToken := model.WebResponse{}
	json.Unmarshal(responseBodyRefreshToken, &webResponseRefreshToken)
	assert.Equal(t, 400, webResponseRefreshToken.Code)
	assert.Equal(t, "Bad Request", webResponseRefreshToken.Status)
	assert.True(t, webResponseRefreshToken.Error)
	assert.Equal(t, "please input the refresh token", webResponseRefreshToken.ErrorMessage)
}

func TestRegistrationController_RefreshTokenFailedBecauseInvalidCharacter(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])

	// REFRESH

	refreshRequest := map[string]string{
		"refresh_token": "xyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjgzODgyODYsInJlZnJlc2hfdXVpZCI6IjZjNjQ3Y2RjLWNkZDktNGNiYS1hMjM1LTgyZGZkYzAyMWM4MSsrZTVmMTBlZjAtYTZkMy00M2FlLThiMGMtYzM2Y2I5NzI4MzllIiwidXNlcl9pZCI6ImU1ZjEwZWYwLWE2ZDMtNDNhZS04YjBjLWMzNmNiOTcyODM5ZSJ9.DVWv7NbAOING_e5IBNBehXtoDO-tizIZdaxw-mVEu3c",
	}

	requestBodyRefreshToken, _ := json.Marshal(refreshRequest)
	requestRefreshToken := httptest.NewRequest("POST", "/api/v1/auth/refresh", bytes.NewBuffer(requestBodyRefreshToken))
	requestRefreshToken.Header.Set("Content-Type", "application/json")
	requestRefreshToken.Header.Set("Accept", "application/json")
	responseRefreshToken, _ := app.Test(requestRefreshToken)
	assert.Equal(t, 401, responseRefreshToken.StatusCode)
	responseBodyRefreshToken, _ := ioutil.ReadAll(responseRefreshToken.Body)
	webResponseRefreshToken := model.WebResponse{}
	json.Unmarshal(responseBodyRefreshToken, &webResponseRefreshToken)
	assert.Equal(t, 401, webResponseRefreshToken.Code)
	assert.Equal(t, "Unauthorized", webResponseRefreshToken.Status)
	assert.True(t, webResponseRefreshToken.Error)
	assert.Equal(t, "invalid character 'Ç' looking for beginning of value", webResponseRefreshToken.ErrorMessage)
}

func TestRegistrationController_RefreshTokenFailedBecauseInvalidSegments(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])

	// REFRESH

	refreshRequest := map[string]string{
		"refresh_token": "mVEu3c",
	}

	requestBodyRefreshToken, _ := json.Marshal(refreshRequest)
	requestRefreshToken := httptest.NewRequest("POST", "/api/v1/auth/refresh", bytes.NewBuffer(requestBodyRefreshToken))
	requestRefreshToken.Header.Set("Content-Type", "application/json")
	requestRefreshToken.Header.Set("Accept", "application/json")
	responseRefreshToken, _ := app.Test(requestRefreshToken)
	assert.Equal(t, 401, responseRefreshToken.StatusCode)
	responseBodyRefreshToken, _ := ioutil.ReadAll(responseRefreshToken.Body)
	webResponseRefreshToken := model.WebResponse{}
	json.Unmarshal(responseBodyRefreshToken, &webResponseRefreshToken)
	assert.Equal(t, 401, webResponseRefreshToken.Code)
	assert.Equal(t, "Unauthorized", webResponseRefreshToken.Status)
	assert.True(t, webResponseRefreshToken.Error)
	assert.Equal(t, "token contains an invalid number of segments", webResponseRefreshToken.ErrorMessage)
}

func TestRegistrationController_RefreshTokenFailedBecauseCannotUnmarshal(t *testing.T) {
	// DELETE ----------------------------------
	registrationRepository.DeleteAll()
	// END DELETE ----------------------------------

	// CREATE ----------------------------------
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	// END CREATE ----------------------------------

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)

	// UPDATE STATUS ----------------------------------
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: createRegistrationResponse.VirtualAccount,
	}
	requestBodyUpdate, _ := json.Marshal(updateStatusRequest)
	request2 := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBodyUpdate))
	request2.Header.Set("Content-Type", "application/json")
	request2.Header.Set("Accept", "application/json")

	response2, _ := app.Test(request2)
	assert.Equal(t, 200, response2.StatusCode)
	responseBody2, _ := ioutil.ReadAll(response.Body)

	webResponse2 := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse2)
	assert.Equal(t, 201, webResponse2.Code)
	assert.Equal(t, "Created", webResponse2.Status)
	assert.Equal(t, false, webResponse2.Error)
	assert.Equal(t, nil, webResponse2.ErrorMessage)
	json.Unmarshal(responseBody2, &webResponse2)
	// END UPDATE STATUS ----------------------------------

	loginRequest := model.LoginRequest{
		Username: createRegistrationResponse.Username,
		Password: createRegistrationResponse.Password,
	}
	requestBodyLogin, _ := json.Marshal(loginRequest)
	request3 := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(requestBodyLogin))
	request3.Header.Set("Content-Type", "application/json")
	request3.Header.Set("Accept", "application/json")

	response3, _ := app.Test(request3)
	assert.Equal(t, 200, response3.StatusCode)
	responseBody3, _ := ioutil.ReadAll(response3.Body)
	webResponse3 := model.WebResponse{}
	json.Unmarshal(responseBody3, &webResponse3)
	assert.Equal(t, 200, webResponse3.Code)
	assert.Equal(t, "Ok", webResponse3.Status)
	assert.Equal(t, false, webResponse3.Error)
	assert.Equal(t, nil, webResponse3.ErrorMessage)

	datamap := make(map[string]string)
	jsonData3, _ := json.Marshal(webResponse3.Data)
	loginResponse := datamap
	json.Unmarshal(jsonData3, &loginResponse)
	assert.NotNil(t, loginResponse["access_token"])
	assert.NotNil(t, loginResponse["refresh_token"])

	// REFRESH

	refreshRequest := map[string]int{
		"refresh_token": 123,
	}

	requestBodyRefreshToken, _ := json.Marshal(refreshRequest)
	requestRefreshToken := httptest.NewRequest("POST", "/api/v1/auth/refresh", bytes.NewBuffer(requestBodyRefreshToken))
	requestRefreshToken.Header.Set("Content-Type", "application/json")
	requestRefreshToken.Header.Set("Accept", "application/json")
	responseRefreshToken, _ := app.Test(requestRefreshToken)
	assert.Equal(t, 422, responseRefreshToken.StatusCode)
	responseBodyRefreshToken, _ := ioutil.ReadAll(responseRefreshToken.Body)
	webResponseRefreshToken := model.WebResponse{}
	json.Unmarshal(responseBodyRefreshToken, &webResponseRefreshToken)
	assert.Equal(t, 422, webResponseRefreshToken.Code)
	assert.Equal(t, "Unprocessable Entity", webResponseRefreshToken.Status)
	assert.True(t, webResponseRefreshToken.Error)
	assert.Equal(t, "json: cannot unmarshal \"123}\" into Go struct field map[string]string.refresh_token of type string", webResponseRefreshToken.ErrorMessage)
}
