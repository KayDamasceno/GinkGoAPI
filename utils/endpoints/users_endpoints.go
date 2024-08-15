package utils

import (
	"github.com/gavv/httpexpect/v2"
	. "github.com/onsi/ginkgo/v2"
)


func CreateUser(t GinkgoTInterface, payload interface{} ) *httpexpect.Response {
	
	api := httpexpect.Default(t, desafioQaURL)

	postResponse := api.POST("/users").WithJSON(payload)

	return postResponse.Expect()
	
}


func GetUsers (t GinkgoTInterface) *httpexpect.Response {

	api := httpexpect.Default(t, desafioQaURL)

	getResponse := api.GET("/users")
	
	return getResponse.Expect()
}


func GetUserById (t GinkgoTInterface, id int) *httpexpect.Response {
	
	api := httpexpect.Default(t, desafioQaURL)

	getResponse := api.GET("/users/{id}").WithPath("id", id)

	return getResponse.Expect()
	
}

func UpdateUserById (t GinkgoTInterface, id int, updatedPayload interface{}) *httpexpect.Response {

	api := httpexpect.Default(t, desafioQaURL)

	putResponse := api.PUT("/users/{id}").WithPath("id", id).WithJSON(updatedPayload)

	return putResponse.Expect()
}

func DeleteUserById (t GinkgoTInterface, id int) *httpexpect.Response {
	
	api := httpexpect.Default(t, desafioQaURL)

	deleteResponse := api.DELETE("/users/{id}").WithPath("id", id)

	return deleteResponse.Expect()
	
}