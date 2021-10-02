package main



import (
	"fmt"
	"log"
	"os"

	"github.com/aditya-suripeddi/covstats/handlers"
	"github.com/aditya-suripeddi/covstats/repository"

	 db "github.com/aditya-suripeddi/covstats/helpers/database"
	 //mdl "github.com/aditya-suripeddi/covstats/middleware"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)


// read configs 
func init() {
	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

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
	//appMiddleware := mdl.InitAppMiddleware(appName)
	//e.Use(appMiddleware.CORS)

	// create repository objects for CRUD
	urMongo := repository.NewRegionInfoRepositoryMongo(mongodb, appName)

	// setup handlers with repository object
	handlers.NewCovidStatsHandler(e, urMongo)
	handlers.NewReverseGeocodeHandler(e, urMongo)

	// start the server 
	e.Start(fmt.Sprintf(`%s:%s`, appHost, appPort))

}
