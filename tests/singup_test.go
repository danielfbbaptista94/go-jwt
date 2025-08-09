package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	requestdto "go-jwt/controllers/requestDTO"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSingUp(t *testing.T) {
	t.Run("when_given_a_valid_body_return_status_created", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := getTestGinContext(recorder)

		request := requestdto.SignupDTO{
			Email:    "test@email.com",
			Password: "123456",
		}

		b, _ := json.Marshal(request)
		body := io.NopCloser(strings.NewReader(string(b)))

		makeRequest(ctx, []gin.Param{}, url.Values{}, "POST", body)
		UserController.Signup(ctx)

		assert.EqualValues(t, http.StatusCreated, recorder.Code)
	})
}
