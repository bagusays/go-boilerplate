package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func init() {
	// os.Chdir("../../../../")
}

func TestHealthCheck(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	res := rec.Result()
	defer res.Body.Close()

	h := NewHealthCheckHandler(e)

	// Assertions
	if assert.NoError(t, h.HealthCheckHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
