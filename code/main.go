package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"mural/api"
	"mural/app"
	"mural/config"
	"mural/controller"
	"mural/controller/movie"
	"mural/db/sql"
	mural_middleware "mural/middleware"
	"mural/worker"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

func main() {
	// setup env
	mural_config, err := config.NewMuralConfig()
	config.Must(err)
	slog.Info(fmt.Sprintf("USING: %s", mural_config.Env))

	// setup database
	dal, err := sql.NewSQLiteDal(mural_config)
	config.Must(err)

	// setup analytics stuff
	var analytics_controller api.IAnalyticsController
	if mural_config.EnabledAnalytics {
		analytics_controller = api.NewPlausibleAnalytics(
			os.Getenv(mural_config.PlausibleURL),
			os.Getenv(mural_config.PlausibleAppDomain),
			os.Getenv(mural_config.AppURL),
		)
	} else {
		analytics_controller = api.STDAnalytics{}
	}

	service, err := app.NewMuralService(
		dal,
		mural_config,
		analytics_controller,
	)
	config.Must(err)

	// setup movie controller
	movie_controller := movie.NewTMDBController(mural_config.TMDBKey)

	// setup schedular
	scheduler := worker.NewMuralSchedular(
		movie_controller,
		service,
	)

	// setup the project
	scheduler.InitProgram()

	// register all of the workers
	if mural_config.Env == config.EnvTest {
		slog.Info("Registering Dev Workers")
		config.Must(scheduler.RegisterWorkersDev())
	} else {
		slog.Info("Registering Workers")
		config.Must(scheduler.RegisterWorkers())
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

	mural_middleware.InitSession(mural_config)
	e.Use(mural_middleware.GetUserKey)
	e.Use(mural_middleware.PassServiceData(
		service,
	))

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

	error_template := template.Must(template.New("mural-error").ParseFiles("view/mural/mural-error.html"))
	templates["404.html"] = error_template
	e.Renderer = &controller.TemplateRenderer{
		Templates: templates,
	}

	// setup routes
	e.Static("/static", "./static")
	if mural_config.Env == config.EnvTest && mural_config.EnableTLS {
		config.Must(
			e.StartTLS(mural_config.Host, "./ssl/certificate.pem", "./ssl/key.pem"),
		)
	} else {
		config.Must(e.Start(mural_config.Host))
	}
}
