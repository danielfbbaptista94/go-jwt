package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-jwt/configuration/initializers"
	"go-jwt/controllers"
	requestdto "go-jwt/controllers/requestDTO"
	"go-jwt/domain/user/services"
	"go-jwt/models/user/repositories"
	"go-jwt/tests/connection"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

var (
	UserController controllers.UserControllerInterface
)

func TestMain(m *testing.M) {
	db, closeConnection := connection.OpenConnection()
	defer func() {
		closeConnection()
	}()

	initializers.DB = db

	_ = initializers.DB.AutoMigrate(&User{})

	repo := repositories.NewUserRepository()
	service := services.NewUserDomainService(repo)
	UserController = controllers.NewUserControllerInterface(service)

	os.Exit(m.Run())
}

func TestLogin(t *testing.T) {
	t.Run("when_not_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		request := requestdto.LoginDTO{
			Email:    "test@email.com",
			Password: "123456",
		}

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		UserController.Login(ctx)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("when_found_login_success", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)

		userTest := User{
			Email:    "test@email.com",
			Password: string(hashedPassword),
		}

		_ = initializers.DB.Create(&userTest)

		request := requestdto.LoginDTO{
			Email:    "test@email.com",
			Password: "123456",
		}

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		UserController.Login(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func (User) TableName() string {
	return "user_models"
}

func getTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return ctx
}

func makeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
