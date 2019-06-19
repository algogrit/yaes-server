package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/algogrit/yaes-server/src/config/db"
	log "github.com/sirupsen/logrus"
	"syreclabs.com/go/faker"

	"github.com/algogrit/yaes-server/src/api"
	"github.com/algogrit/yaes-server/src/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var jwtHeader string

var _ = Describe("User API", func() {
	BeforeEach(func() {
		Cleaner.Acquire("users")

		user := model.User{
			Username:     faker.Internet().UserName(),
			FirstName:    faker.Name().FirstName(),
			LastName:     faker.Name().LastName(),
			MobileNumber: faker.PhoneNumber().PhoneNumber()}
		user.HashedPassword = model.HashAndSalt(faker.Internet().Password(8, 14))
		if err := db.Instance().Create(&user).Error; err != nil {
			log.Fatal("Unable to save record: ", err.Error())
		}

		jwtSigningKey := []byte("483175006c1088c849502ef22406ac4e")
		jwtToken := model.CreateJWTToken(user, jwtSigningKey)["token"]

		var jwtHeaderBuilder strings.Builder
		jwtHeaderBuilder.WriteString("Bearer ")
		jwtHeaderBuilder.WriteString(jwtToken)

		jwtHeader = jwtHeaderBuilder.String()
	})

	AfterEach(func() {
		Cleaner.Clean("users")
	})

	Describe("GET /users", func() {
		Context("when the user is logged in", func() {
			It("should return list of other users", func() {
				user := model.User{
					Username:     faker.Internet().UserName(),
					FirstName:    "Gaurav",
					LastName:     faker.Name().LastName(),
					MobileNumber: faker.PhoneNumber().PhoneNumber()}
				user.HashedPassword = model.HashAndSalt(faker.Internet().Password(8, 14))
				if err := db.Instance().Create(&user).Error; err != nil {
					log.Fatal("Unable to save record: ", err.Error())
				}

				req, _ := http.NewRequest("GET", "/users", nil)
				req.Header.Set("Authorization", jwtHeader)
				response := httptest.NewRecorder()
				api.Instance().ServeHTTP(response, req)

				Expect(response.Code).To(Equal(http.StatusOK))

				var users []model.User

				if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
					log.Fatal("Unable to decode users: ", err.Error())
				}
				Expect(users).ShouldNot(BeEmpty())
				Expect(len(users)).To(Equal(1))
				Expect(users[0].FirstName).To(Equal("Gaurav"))
				Expect(users[0].HashedPassword).Should(BeEmpty())
			})
		})

		Context("when the user is not logged in", func() {
			It("should fail with unauthorized code", func() {
				req, _ := http.NewRequest("GET", "/users", nil)
				response := httptest.NewRecorder()
				api.Instance().ServeHTTP(response, req)

				Expect(response.Code).To(Equal(http.StatusUnauthorized))
			})
		})
	})

	Describe("POST /users", func() {
		XIt("should create a user", func() {})
	})

	Describe("POST /login", func() {
		XIt("should create a session for valid credentials", func() {})
		XIt("shouldn't create a session for invalid credentials", func() {})
	})
})
