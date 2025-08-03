package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	requestdto "go-jwt/controllers/requestDTO"
	domain "go-jwt/domain/user"
	errorHandler "go-jwt/errorHandler"
	mocks "go-jwt/tests/mocks"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserController_Signup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var userDomain = mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(userDomain)

	t.Run("when_body_invalid_then_return_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		request := requestdto.SignupDTO{
			Email:    "testemail.com",
			Password: "pas",
		}

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		controller.Signup(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when_body_valid_but_return_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		request := requestdto.SignupDTO{
			Email:    "teste@mail.com",
			Password: "password",
		}
		userCreate := domain.NewUserDomain(request.Email, request.Password)

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		userDomain.EXPECT().CreateUser(userCreate).Return(
			errorHandler.NewInternalError("error creating user"))
		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		controller.Signup(ctx)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("when_body_valid_then_return_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		request := requestdto.SignupDTO{
			Email:    "teste@mail.com",
			Password: "password",
		}
		userCreate := domain.NewUserDomain(request.Email, request.Password)

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		userDomain.EXPECT().CreateUser(userCreate).Return(nil)
		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		controller.Signup(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
