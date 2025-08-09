package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	requestdto "go-jwt/controllers/requestDTO"
	domain "go-jwt/domain/user"
	errorHandler "go-jwt/errorHandler"
	"go-jwt/tests/mocks"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserController_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var userDomain = mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(userDomain)

	t.Run("when_body_invalid_then_return_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		request := requestdto.LoginDTO{
			Email:    "testemail.com",
			Password: "pas",
		}

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		controller.Login(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when_body_valid_but_email_was_not_found_then_return_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		request := requestdto.LoginDTO{
			Email:    "tese@mail.com",
			Password: "password",
		}

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		userDomain.EXPECT().Login(request).Return(
			nil, "", errorHandler.NewNotFoundError("user not found"))
		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		controller.Login(ctx)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("when_body_valid_but_password_is_wrong_then_return_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		request := requestdto.LoginDTO{
			Email:    "teste@mail.com",
			Password: "paword",
		}

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		userDomain.EXPECT().Login(request).Return(nil, "", errorHandler.NewBadRequestError("invalid password"))
		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		controller.Login(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when_body_valid_but_failed_create_token", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		request := requestdto.LoginDTO{
			Email:    "teste@mail.com",
			Password: "12345",
		}

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		userDomain.EXPECT().Login(request).Return(nil, "", errorHandler.NewBadRequestError("failed to create token"))
		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		controller.Login(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when_body_valid_then_return_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
		request := requestdto.LoginDTO{
			Email:    "teste@mail.com",
			Password: "password",
		}

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		userDomain.EXPECT().Login(request).Return(
			domain.NewUserDomain(request.Email, "password"), token, nil)
		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		controller.Login(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
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
