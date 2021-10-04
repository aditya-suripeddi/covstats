package handlers

import (
	"log"
	"os"
	"net/http"
	"net/http/httptest"
    "testing"

	"github.com/aditya-suripeddi/covstats/repository"

	db "github.com/aditya-suripeddi/covstats/helpers/database"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)


// https://stackoverflow.com/questions/24790175/when-is-the-init-function-run
// read configs
func init() {
	viper.SetConfigFile(`../config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}


func TestCovidStatsHandler(t *testing.T) {

	// read mongodb credentials from config.json
	var mongoCredential = map[string]string{
		"host":     viper.GetString(`database.mongodb.host`),
		"user":     viper.GetString(`database.mongodb.user`),
		"password": viper.GetString(`database.mongodb.password`),
		"db":       viper.GetString(`database.mongodb.db`),
	}

	// read appName, it is also the collection name in mongodb
	appName := viper.GetString(`app.name`)

	// setup mongodb connection
	mongodb, err := db.GetMongoDB(mongoCredential)

	if err != nil {
		log.Fatal(err, ": mongodb setup failed")
		os.Exit(1)
	}

	defer mongodb.Logout()
	

	// setup echo
	e := echo.New()

	// create repository objects for CRUD
	urMongo := repository.NewRegionInfoRepositoryMongo(mongodb, appName)


    cshandler := &CovidStatsHandler{regionInfoRepo: urMongo}	
	req := httptest.NewRequest(http.MethodGet, "/states", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, cshandler.CovidStats(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		t.Log(rec.Body.String())
		assert.NotEmpty(t, rec.Body.String()) // not comparing contents yet, only checking for non emtpy body
	}

}
