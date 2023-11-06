package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
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

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	ErrCouldNotParseTempaltes = fmt.Errorf("could not parse templates")
)

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, name, data)
}

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}

	// validate env
	err = config.ValidateENV()
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}

	// setup database
	sqlDAL, err := sql.NewSQLiteDal(os.Getenv("DATABASE_FILE"))
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

	// setup movie controlle
	movie_controller := movie.NewTMDBController()
	api.MovieController = movie_controller
	api.RandomAnswerKey = rand.Intn(5)
	api.RandomPageKey = rand.Intn(300)


	// setup schedular
	scheduler :=  worker.NewMuralSchedular()

	// setup the project
	scheduler.InitProgram()
	// register all of the workers
	err = scheduler.RegisterWorkers()
	if err != nil {
		slog.Error(err.Error())
		panic(1)
	}

	// start scheduler
	scheduler.StartScheduler()

	// start setting up 
	e := echo.New()

	// define templates
	templates := map[string]*template.Template{}

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(mural_middleware.GenerateUserKey)

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


	error_template := template.Must(
		template.New("mural-error").ParseFiles("view/mural/mural-error.html"),
	)


	templates["404.html"] = error_template

	e.Renderer = &TemplateRenderer{
		templates: templates,
	}


	// setup routes
	e.Static("/static", "./static")
	e.Logger.Fatal(e.Start(":1323"))
}
