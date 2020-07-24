package http_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/internal/config"
	userHTTP "algogrit.com/yaes-server/users/http"
	"algogrit.com/yaes-server/users/service"

	"github.com/golang/mock/gomock"
	"github.com/justinas/alice"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Json", func() {
	// var jwtSigningKey string
	// jwtSigningKey = "483175006c1088c849502ef22406ac4e"

	var mockUserService *service.MockUserService
	var userHandler http.Handler

	var req *http.Request
	var rw *httptest.ResponseRecorder

	currentUser := entities.User{Username: "algogrit"}
	currentUser.ID = 1

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockUserService = service.NewMockUserService(mockCtrl)
		// TODO: Mocking wouldn't help here at all, we need to stub alice.Chain
		// mockAuthChain := auth.NewMockAuth(mockCtrl)

		ctx := context.WithValue(context.Background(), config.LoggedInUser, currentUser)

		req = httptest.NewRequest("GET", "/users", nil)
		req = req.WithContext(ctx)

		rw = httptest.NewRecorder()

		userHandler = userHTTP.NewHTTPHandler(mockUserService, alice.Chain{})
	})

	Describe("Index", func() {
		Context("when a user is logged in", func() {
			It("should serialize a list of users as json", func() {
				mockUserService.
					EXPECT().
					Index(gomock.Any(), gomock.Eq(currentUser)).
					Return([]*entities.User{{Username: "ga"}}, nil)

				userHandler.ServeHTTP(rw, req)

				response := rw.Result()
				Expect(response.StatusCode).To(Equal(http.StatusOK))

				// resByte, err := ioutil.ReadAll(response.Body)
				// resString := string(resByte)
				// Expect(err).To(BeNil())
				// Expect(resString).To(Equal("[{\"Username\": \"ga\"}]\n"))

				var users []entities.User
				json.NewDecoder(response.Body).Decode(&users)

				Expect(len(users)).To(Equal(1))
				Expect(users[0].Username).To(Equal("ga"))
			})
		})
	})
})
