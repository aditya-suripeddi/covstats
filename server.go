package main

// dependencies:
//
//    https://echo.labstack.com/guide/:  go get github.com/labstack/echo/v4
//
// 	  https://github.com/mongodb/mongo-go-driver: go get go.mongodb.org/mongo-driver/mongo
//
//    go get github.com/tidwall/gjson
//
//    https://gopkg.in/mgo.v2: go get gopkg.in/mgo.v2
//
//    go get github.com/spf13/viper
//

import (
	"fmt"
	"log"
	"os"

	"covstats/handlers"
	"covstats/repository"

	db "covstats/helpers/database"

	"github.com/labstack/echo/v4"
	//mdl "covstats/middleware"
	"github.com/spf13/viper"
)

//
//  references:
//
//        https://betterprogramming.pub/intro-77f65f73f6d3          - to making client api calls
//        https://www.mohfw.gov.in/data/datanew.json                - for statewise covid data
//        https://blog.logrocket.com/making-http-requests-in-go/    - more info on echo
//

func init() {
	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	var mongoCredential = map[string]string{
		"host":     viper.GetString(`database.mongodb.host`),
		"user":     viper.GetString(`database.mongodb.user`),
		"password": viper.GetString(`database.mongodb.password`),
		"db":       viper.GetString(`database.mongodb.db`),
	}
	appName := viper.GetString(`app.name`) // also the collection name 
	appHost := viper.GetString(`app.domain`)
	appPort := viper.GetString(`app.port`)

	mongodb, err := db.GetMongoDB(mongoCredential)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer mongodb.Logout()

	e := echo.New()
	//appMiddleware := mdl.InitAppMiddleware(appName)
	//e.Use(appMiddleware.CORS)

	urMongo := repository.NewRegionInfoRepositoryMongo(mongodb, appName)

	handlers.NewCovidStatsHandler(e, urMongo)
	handlers.NewReverseGeocodeHandler(e, urMongo)

	e.Start(fmt.Sprintf(`%s:%s`, appHost, appPort))

}
