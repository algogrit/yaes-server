package service_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"algogrit.com/yaes-server/entities"
	"algogrit.com/yaes-server/internal/config"
	"algogrit.com/yaes-server/users/repository"
	"algogrit.com/yaes-server/users/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Json", func() {
	var jwtSigningKey string
	jwtSigningKey = "483175006c1088c849502ef22406ac4e"

	var userRepoStub repository.UserRepository
	var userService service.UserService

	var req *http.Request
	var rw *httptest.ResponseRecorder

	BeforeEach(func() {
		user := entities.User{Username: "algogrit"}
		user.ID = 1
		ctx := context.WithValue(context.Background(), config.LoggedInUser, user)

		req = httptest.NewRequest("GET", "/users", nil)
		req = req.WithContext(ctx)

		rw = httptest.NewRecorder()

		retrieveOthersFn := func(u entities.User) (users []*entities.User, err error) {
			return []*entities.User{
				&entities.User{Username: "ga"},
			}, nil
		}

		userRepoStub = &repository.UserRepoStub{RetrieveOthersFn: retrieveOthersFn}
		userService = service.New(userRepoStub, jwtSigningKey)
	})

	Describe("Index", func() {
		Context("when a user is logged in", func() {
			It("should serialize a list of users as json", func() {
				userService.Index(rw, req)

				response := rw.Result()
				Expect(response.StatusCode).To(Equal(http.StatusOK))

				// resByte, err := ioutil.ReadAll(response.Body)
				// resString := string(resByte)
				// Expect(err).To(BeNil())
				// Expect(resString).To(Equal("[{\"Username\": \"ga\"}]\n"))

				var users []entities.User
				json.NewDecoder(response.Body).Decode(&users)
				Expect(users[0].Username).To(Equal("ga"))
			})
		})
	})
})