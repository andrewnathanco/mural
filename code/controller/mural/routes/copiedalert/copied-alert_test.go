package copiedalert

import (
	"log/slog"
	"mural/config"
	"mural/db"
	"mural/db/sql"
	"mural/middleware"
	"mural/worker"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/andybalholm/cascadia"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

var (
	endpoint = "/mural/copied-alert"
)

func init() {
	// setup sessions
	os.Setenv("SESSION_KEY", "test")
	os.Setenv("DATABASE_FILE", "./test/mural.db")
	os.Setenv("TMDB_KEY", "test")

	err := config.ValidateENV()
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}

	middleware.InitSession()

	// setup database
	sqlDAL, err := sql.NewSQLiteDal(os.Getenv(config.EnvDatabasFile))
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}

	// setup the metadata for the app
	err = sqlDAL.SetupMetadata()
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}

	db.DAL = sqlDAL

	// setup schedular
	scheduler := worker.NewMuralSchedular()

	// setup the project
	scheduler.InitProgram()
}

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
		dat, err := os.ReadFile("/view/mural/alerts/copied-success.html")
		assert.NoError(t, err)

        assert.Equal(t, http.StatusOK, rec.Code)
		doc, err := html.Parse(strings.NewReader(string(dat)))
		assert.NoError(t, err)

		sel, err := cascadia.Parse("#copied-message")
		assert.NoError(t, err)

		node := cascadia.Query(doc, sel)
		data := node.Data
		assert.Equal(t, data, "Copied to clipboard!")
    }
}

func TestCopiedAlertFailure(t *testing.T) {
}


func TestCopiedAlertInvalid(t *testing.T) {
}