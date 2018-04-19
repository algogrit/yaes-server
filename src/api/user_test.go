package api_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gauravagarwalr/yaes-server/src/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User API", func() {
	Describe("GET /users", func() {
		Context("when the user is not logged in", func() {
			It("should fail with unauthorized code", func() {
				req, _ := http.NewRequest("GET", "/users", nil)
				response := httptest.NewRecorder()
				api.Instance().ServeHTTP(response, req)

				Expect(response.Code).To(Equal(http.StatusUnauthorized))
			})
		})
	})
})
