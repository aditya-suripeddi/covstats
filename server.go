package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aditya-suripeddi/covstats/handlers"
	"github.com/aditya-suripeddi/covstats/repository"

	db "github.com/aditya-suripeddi/covstats/helpers/database"
	mdl "github.com/aditya-suripeddi/covstats/middleware"
	_ "github.com/aditya-suripeddi/covstats/docs/swagdocs"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	 echoSwagger "github.com/swaggo/echo-swagger"
)


// https://stackoverflow.com/questions/24790175/when-is-the-init-function-run
// read configs
func init() {
	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Covstats Swagger API
// @version 1.0
// @description Covid Stats for your region
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
// @schemes http

func main() {

	// read mongodb credentials from config.json
	var mongoCredential = map[string]string{
		"host":     viper.GetString(`database.mongodb.host`),
		"user":     viper.GetString(`database.mongodb.user`),
		"password": viper.GetString(`database.mongodb.password`),
		"db":       viper.GetString(`database.mongodb.db`),
	}

	// read appName, it is also the collection name in mongodb
	appName := viper.GetString(`app.name`)

	// read server host and port
	appHost := viper.GetString(`app.domain`)
	appPort := viper.GetString(`app.port`)

	// setup mongodb connection
	mongodb, err := db.GetMongoDB(mongoCredential)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer mongodb.Logout()

	e := echo.New()
	appMiddleware := mdl.InitAppMiddleware(appName)
	e.Use(appMiddleware.CORS)

	// create repository objects for CRUD
	urMongo := repository.NewRegionInfoRepositoryMongo(mongodb, appName)

	// setup handlers with repository object
	handlers.NewCovidStatsHandler(e, urMongo)
	handlers.NewReverseGeocodeHandler(e, urMongo)


	e.GET("/", handlers.ServerStatus)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start the server
	e.Start(fmt.Sprintf(`%s:%s`, appHost, appPort))

}
