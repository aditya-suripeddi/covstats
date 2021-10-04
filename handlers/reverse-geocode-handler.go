package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/aditya-suripeddi/covstats/helpers/wrapper"
	"github.com/aditya-suripeddi/covstats/model"
	"github.com/aditya-suripeddi/covstats/repository"

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

func isValid(lat string, lon string) bool {
	//https://stackoverflow.com/questions/3518504/regular-expression-for-matching-latitude-longitude-coordinates
	// https://stackoverflow.com/questions/66624011/how-to-validate-an-email-address-in-go
	regex := regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`)
	return regex.MatchString(lat) && regex.MatchString(lon)
}

// @Summary Get state from lat, lon and send covstats in that state and India
// @Tags root
// @Accept json
// @Produce json
// @Param lat path string true "latitude"
// @Param lon path string true "longitude"
// @Success 200 {object}  wrapper.HttpSuccess{Data=model.Region}
// @Failure 500 {object}  wrapper.HttpFail
// @Router /lat/{lat}/lon/{lon} [get]
func (rghandler *ReverseGeocodeHandler) GetState(c echo.Context) error {

	//https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go
	// the code below is taken from link above

	lat := c.Param("lat")
	lon := c.Param("lon")

	//  need to review:  
	//if !isValid(lat, lon) {
	// 	errorMessage := fmt.Sprintf("lat %s, lon %s validation failed", lat, lon)
	// 	return wrapper.Error(http.StatusBadRequest, errorMessage, c)
	// }

	URL := fmt.Sprintf("https://us1.locationiq.com/v1/reverse.php?key=pk.8b79e5c7f4eb5381aab22c4c26d0e3d3&lat=%s&lon=%s&format=json", lat, lon)

	var client http.Client

	respo, erro := client.Get(URL)

	if erro != nil {
		log.Fatal(erro)
		return wrapper.Error(http.StatusInternalServerError, erro.Error(), c)
	}

	defer respo.Body.Close()

	if respo.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("Received %s http status code from locationiq server, check your lat & lon values", respo.Status)
		return wrapper.Error(respo.StatusCode, errorMessage, c)
	}

	bodyBytes, err := ioutil.ReadAll(respo.Body)
	if err != nil {
		log.Fatal(err)
		return wrapper.Error(http.StatusInternalServerError, erro.Error(), c)
	}

	bodyString := string(bodyBytes)
	state := gjson.Get(bodyString, "address.state")
	message := "Reverse geocoding done with https://locationiq.com/"

	if state.Type == gjson.Null {
		errorMessage := fmt.Sprintf("State not found in locationiq server response, check your lat, lon values: %s", bodyString)
		return wrapper.Error(http.StatusInternalServerError, errorMessage, c)
	}

	log.Println(state)

	var covstats model.Region

	regionInfo, err := rghandler.regionInfoRepo.FindByRegion(state.String())

	if err != nil {
		log.Fatal(err)
		return wrapper.Error(http.StatusInternalServerError, erro.Error(), c)
	}

	nationInfo, erro := rghandler.regionInfoRepo.FindByRegion("India")

	if erro != nil {
		log.Fatal(err)
		return wrapper.Error(http.StatusInternalServerError, erro.Error(), c)
	}

	covstats = append(covstats, *regionInfo, *nationInfo)

	return wrapper.Data(http.StatusOK, covstats, message, c)
}
