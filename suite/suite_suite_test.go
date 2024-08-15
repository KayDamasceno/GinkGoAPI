package suite_test

import (
	"fmt"
	"ginkgoapi/schemas"
	templates "ginkgoapi/templates"
	endpoints "ginkgoapi/utils/endpoints"
	helpers "ginkgoapi/utils/helpers"
	"net/http"
	"testing"

	"github.com/bxcodec/faker/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User API", func ()  {

	BeforeEach(func ()  {
		fmt.Println("Token acquired")
	})

	AfterEach(func ()  {
		fmt.Println("CleanUp done")
	})

	Context("Create User", func ()  {
		It("should successfully create a user", func ()  {
			var user map[string]interface{}
			
			helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

			user["name"] = faker.Name()
			user["last_name"] = faker.LastName()
			user["email"] = faker.Email()

			response := endpoints.CreateUser(GinkgoT(), user)

			Expect(response.Raw().StatusCode).To(Equal(http.StatusCreated))
			Expect(response.JSON().Object().Value("id")).ShouldNot(BeNil())

		
		})
	})

	Context("Get Users", Label("Smoke"), func ()  {
		It("should successfully retrieve users", func ()  {
			response := endpoints.GetUsers(GinkgoT())

			Expect(response.Raw().StatusCode).To(Equal(http.StatusOK))
			Expect(response.JSON().Array().Schema(schemas.UsersGetSchema)).To(Succeed())
		})
	})

	Context("Get User By Id", Label("Regression"), func ()  {

		It("should successfully retrieve a user by Id", func ()  {
			var user map[string]interface{}
			
			helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)

			user["name"] = faker.Name()
			user["last_name"] = faker.LastName()
			user["email"] = faker.Email()

			createResponse := endpoints.CreateUser(GinkgoT(), user)
			Expect(createResponse.Raw().StatusCode).To(Equal(http.StatusCreated))

			id := createResponse.JSON().Object().Value("id").Number().Raw()

			getResponse := endpoints.GetUserById(GinkgoT(), int(id))
			Expect(getResponse.Raw().StatusCode).To(Equal(http.StatusOK))

			userObject := getResponse.JSON().Object()
			Expect(userObject.Value("name").Raw()).To(Equal(user["name"]))
			Expect(userObject.Value("last_name").Raw()).To(Equal(user["last_name"]))
			Expect(userObject.Value("email").Raw()).To(Equal(user["email"]))
		})
	})

	Context("Update user By Id", Label("Regression", "Smoke"), func ()  {

		It("should successfully update a user by Id", func ()  {
			var user map[string]interface{}
			var userUpdated map[string]interface{}
			
			helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)
			helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &userUpdated)

			user["name"] = faker.Name()
			user["last_name"] = faker.LastName()
			user["email"] = faker.Email()

			createResponse := endpoints.CreateUser(GinkgoT(), user)
			Expect(createResponse.Raw().StatusCode).To(Equal(http.StatusCreated))

			id := createResponse.JSON().Object().Value("id").Number().Raw()


			userUpdated["name"] = faker.Name()
			userUpdated["last_name"] = faker.LastName()
			userUpdated["email"] = faker.Email()

			updateResponse := endpoints.UpdateUserById(GinkgoT(), int(id), userUpdated)
			Expect(updateResponse.Raw().StatusCode).To(Equal(http.StatusOK))

			getResponse := endpoints.GetUserById(GinkgoT(), int(id))
			Expect(getResponse.Raw().StatusCode).To(Equal(http.StatusOK))

			userObject := getResponse.JSON().Object()
			Expect(userObject.Value("name").Raw()).To(Equal(userUpdated["name"]))
			Expect(userObject.Value("last_name").Raw()).To(Equal(userUpdated["last_name"]))
			Expect(userObject.Value("email").Raw()).To(Equal(userUpdated["email"]))
		})
	})

	Context("Delete user By Id", Label("Regression"), func ()  {

		It("should successfully delete a user by Id", func ()  {
			var user map[string]interface{}
		
			
			helpers.ConvertJsonTemplateToMap(templates.UserCreateTemplate, &user)
			

			user["name"] = faker.Name()
			user["last_name"] = faker.LastName()
			user["email"] = faker.Email()

			createResponse := endpoints.CreateUser(GinkgoT(), user)
			Expect(createResponse.Raw().StatusCode).To(Equal(http.StatusCreated))

			id := createResponse.JSON().Object().Value("id").Number().Raw()


			
			delteResponse := endpoints.DeleteUserById(GinkgoT(), int(id))
			Expect(delteResponse.Raw().StatusCode).To(Equal(http.StatusOK))

			getResponse := endpoints.GetUserById(GinkgoT(), int(id))
			Expect(getResponse.Raw().StatusCode).To(Equal(http.StatusNotFound))

			Expect(getResponse.JSON().Object().Value("message").Raw()).To(Equal("User not found"))
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Suite Suite")
}
