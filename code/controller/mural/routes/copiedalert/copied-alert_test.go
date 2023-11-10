package copiedalert

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	endpoint = "/mural/copied-alert"
)


func TestCopiedAlertSuccess(t *testing.T) {
    // Setup variables
	e := echo.New()
	q := make(url.Values)
	q.Set("alert", "success")

	// Create request
	full_url := endpoint + "?" + q.Encode()
    req := httptest.NewRequest(http.MethodPut, full_url, nil)
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()

	// Create request context
	c := e.NewContext(req, rec)
	c.SetPath(endpoint)

    if assert.NoError(t, OpenCopiedAlert(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        // assert.Equal(t, userJSON, rec.Body.String())
    }
}

func TestCopiedAlertFailure(t *testing.T) {
}


func TestCopiedAlertInvalid(t *testing.T) {
}