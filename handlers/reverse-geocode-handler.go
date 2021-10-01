package handlers

import (
	"covstats/helpers/wrapper"
	"covstats/repository"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

type ReverseGeocodeHandler struct {
	regionInfoRepo repository.RegionInfoRepository
}

// NewUserHandler is constructor
func NewReverseGeocodeHandler(e *echo.Echo, repo repository.RegionInfoRepository) {
	rghandler := &ReverseGeocodeHandler{
		regionInfoRepo: repo,
	}

	e.GET("/lat/:lat/lon/:lon", rghandler.GetState)
}

// GetState - handler method for binding JSON body and scraping for statewise covid data
func (rghandler *ReverseGeocodeHandler) GetState(c echo.Context) error {

	//https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go
	// the code below is taken from link above

	lat := c.Param("lat")
	long := c.Param("lon")
	URL := fmt.Sprintf("https://us1.locationiq.com/v1/reverse.php?key=pk.8b79e5c7f4eb5381aab22c4c26d0e3d3&lat=%s&lon=%s&format=json", lat, long)

	var client http.Client

	respo, erro := client.Get(URL)

	if erro != nil {
		log.Fatal(erro)
		return wrapper.Error(http.StatusInternalServerError, erro.Error(), c)
	}

	defer respo.Body.Close()

	if respo.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("Received %s http status code from locationiq server", respo.Status)
		return wrapper.Error(http.StatusInternalServerError, errorMessage, c)
	}

	bodyBytes, err := ioutil.ReadAll(respo.Body)
	if err != nil {
		log.Fatal(err)
		return wrapper.Error(http.StatusInternalServerError, erro.Error(), c)
	}

	bodyString := string(bodyBytes)
	state := gjson.Get(bodyString, "address.state")
	message := "Reverse geocoding done with https://locationiq.com/"

	log.Println(state)
	regionInfo, err := rghandler.regionInfoRepo.FindByRegion(state.String())
	

	if err != nil {
		log.Fatal(err)
		return wrapper.Error(http.StatusInternalServerError, erro.Error(), c)
	}

	return wrapper.Data(http.StatusOK, regionInfo, message, c)
}
