package main

import (
	"html/template"
	"log/slog"
	"math/rand"
	"mural/api"
	"mural/config"
	"mural/controller"
	"mural/controller/movie"
	"mural/db"
	"mural/db/sql"
	mural_middleware "mural/middleware"
	"mural/worker"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

func Must(err error) {
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}
}

func main() {
	// load env
	err := godotenv.Load()

	environment := os.Getenv("ENV")

	// validate env
	Must(config.ValidateENV())

	// setup analytics stuff
	enable_analytics := os.Getenv(config.EnvEnableAnalytics)
	enable_analytics_bool, _ := strconv.ParseBool(enable_analytics)
	if (enable_analytics_bool) {
		api.AnalyticsController = api.NewPlausibleAnalytics(
			os.Getenv(config.EnvPlausibleURL),
			os.Getenv(config.EnvAppDomain),
			os.Getenv(config.EnvAppURL),
		)
	} else {
		api.AnalyticsController = api.STDAnalytics{}
	}

	// setup database
	sqlDAL, err := sql.NewSQLiteDal(os.Getenv(config.EnvDatabasFile))
	Must(err)

	db.DAL = sqlDAL
// setup movie controlle
	movie_controller := movie.NewTMDBController()
	api.MovieController = movie_controller
	api.RandomAnswerKey = rand.Intn(5)
	api.RandomPageKey = rand.Intn(300)


	// setup schedular
	scheduler := worker.NewMuralSchedular()

	// setup the project
	scheduler.InitProgram()
	// register all of the workers

	if strings.EqualFold(environment, "dev") {
		Must(scheduler.RegisterWorkersFreeplay())
	} else {
		Must(scheduler.RegisterWorkers())
	}

	// start scheduler
	scheduler.StartScheduler()

	// start setting up 
	e := echo.New()

	// define templates
	templates := map[string]*template.Template{}

	// middleware
	e.Use(slogecho.New(slog.Default()))
	e.Use(middleware.Recover())

	mural_middleware.InitSession()
	e.Use(mural_middleware.GetUserKey)

    // Define your routes and handlers here
	// setup routes and controllers
	route_conrollers := controller.GetRouteControllers()

	for _, route_controller := range route_conrollers {
		// add templates
		for _, template := range route_controller.Controller.GetTemplates() {
			templates[template.Name] = template.Template
		}

		// add routes
		route_controller.Router.ConfigureRouter(route_controller.Controller, e)
	}

	error_template := template.Must( template.New("mural-error").ParseFiles("view/mural/mural-error.html"),)
	templates["404.html"] = error_template
	e.Renderer = &controller.TemplateRenderer{
		Templates: templates,
	}

	// setup routes
	e.Static("/static", "./static")
	if strings.EqualFold(environment, "dev") {
		Must(e.Start("10.0.0.42:1323"))
	} else {
		Must(e.Start(":1323"))
	}
}
